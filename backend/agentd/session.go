package agentd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"github.com/gogo/protobuf/proto"
	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
	corev2 "github.com/echlebek/sensu-lite/api/core/v2"
	"github.com/echlebek/sensu-lite/backend/messaging"
	"github.com/echlebek/sensu-lite/backend/ringv2"
	"github.com/echlebek/sensu-lite/backend/store"
	"github.com/echlebek/sensu-lite/handler"
	"github.com/echlebek/sensu-lite/transport"
	"github.com/sirupsen/logrus"
)

var (
	sessionCounter = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "sensu_go_agent_sessions",
			Help: "Number of active agent sessions on this backend",
		},
		[]string{"namespace"},
	)
)

// ProtobufSerializationHeader is the Content-Type header which indicates protobuf serialization.
const ProtobufSerializationHeader = "application/octet-stream"

// JSONSerializationHeader is the Content-Type header which indicates JSON serialization.
const JSONSerializationHeader = "application/json"

// MarshalFunc is the function signature for protobuf/JSON marshaling.
type MarshalFunc = func(pb proto.Message) ([]byte, error)

// UnmarshalFunc is the function signature for protobuf/JSON unmarshaling.
type UnmarshalFunc = func(buf []byte, pb proto.Message) error

// UnmarshalJSON is a wrapper to deserialize proto messages with JSON.
func UnmarshalJSON(b []byte, msg proto.Message) error { return json.Unmarshal(b, &msg) }

// MarshalJSON is a wrapper to serialize proto messages with JSON.
func MarshalJSON(msg proto.Message) ([]byte, error) { return json.Marshal(msg) }

// SessionStore specifies the storage requirements of the Session.
type SessionStore interface {
	store.EntityStore
	store.NamespaceStore
}

// A Session is a server-side connection between a Sensu backend server and
// the Sensu agent process via the Sensu transport. It is responsible for
// relaying messages to the message bus on behalf of the agent and from the
// bus to the agent from other daemons. It handles transport handshaking and
// transport channel multiplexing/demultiplexing.
type Session struct {
	cfg          SessionConfig
	conn         transport.Transport
	store        SessionStore
	handler      *handler.MessageHandler
	wg           *sync.WaitGroup
	sendq        chan *transport.Message
	checkChannel chan interface{}
	bus          messaging.MessageBus
	ringPool     *ringv2.Pool
	ctx          context.Context
	cancel       context.CancelFunc
	marshal      MarshalFunc
	unmarshal    UnmarshalFunc

	subscriptions chan messaging.Subscription
}

func newSessionHandler(s *Session) *handler.MessageHandler {
	handler := handler.NewMessageHandler()
	handler.AddHandler(transport.MessageTypeKeepalive, s.handleKeepalive)
	handler.AddHandler(transport.MessageTypeEvent, s.handleEvent)

	return handler
}

// A SessionConfig contains all of the ncessary information to initialize
// an agent session.
type SessionConfig struct {
	ContentType   string
	Namespace     string
	AgentAddr     string
	AgentName     string
	User          string
	Subscriptions []string
	RingPool      *ringv2.Pool
	WriteTimeout  int
}

// NewSession creates a new Session object given the triple of a transport
// connection, message bus, and store.
// The Session is responsible for stopping itself, and does so when it
// encounters a receive error.
func NewSession(ctx context.Context, cfg SessionConfig, conn transport.Transport, bus messaging.MessageBus, store store.Store, unmarshal UnmarshalFunc, marshal MarshalFunc) (*Session, error) {
	// Validate the agent namespace
	if _, err := store.GetNamespace(ctx, cfg.Namespace); err != nil {
		return nil, fmt.Errorf(
			"could not retrieve the namespace '%s': %s", cfg.Namespace, err.Error(),
		)
	}

	logger.WithFields(logrus.Fields{
		"addr":          cfg.AgentAddr,
		"namespace":     cfg.Namespace,
		"agent":         cfg.AgentName,
		"subscriptions": cfg.Subscriptions,
	}).Info("agent connected")

	ctx, cancel := context.WithCancel(ctx)

	s := &Session{
		conn:          conn,
		cfg:           cfg,
		wg:            &sync.WaitGroup{},
		sendq:         make(chan *transport.Message, 10),
		checkChannel:  make(chan interface{}, 100),
		store:         store,
		bus:           bus,
		subscriptions: make(chan messaging.Subscription, len(cfg.Subscriptions)),
		ctx:           ctx,
		cancel:        cancel,
		ringPool:      cfg.RingPool,
		unmarshal:     unmarshal,
		marshal:       marshal,
	}
	s.handler = newSessionHandler(s)
	return s, nil
}

// Receiver returns the check channel for the session.
func (s *Session) Receiver() chan<- interface{} {
	return s.checkChannel
}

func (s *Session) receiver() {
	defer func() {
		s.cancel()
		s.wg.Done()
		logger.Info("shutting down agent session: stopping receiver")
	}()

	for {
		if err := s.ctx.Err(); err != nil {
			return
		}
		msg, err := s.conn.Receive()
		if err != nil {
			switch err := err.(type) {
			case transport.ConnectionError, transport.ClosedError:
				logger.WithFields(logrus.Fields{
					"addr":  s.cfg.AgentAddr,
					"agent": s.cfg.AgentName,
				}).WithError(err).Warn("stopping session")
			default:
				logger.WithError(err).Error("recv error")
			}
			return
		}
		if err := s.handler.Handle(s.ctx, msg.Type, msg.Payload); err != nil {
			logger.WithError(err).WithFields(logrus.Fields{
				"type":    msg.Type,
				"payload": string(msg.Payload)}).Error("error handling message")
		}
	}
}

func (s *Session) sender() {
	defer func() {
		s.cancel()
		s.wg.Done()
		logger.Info("shutting down agent session: stopping sender")
	}()

	for {
		var msg *transport.Message
		select {
		case msg = <-s.sendq:
		case c := <-s.checkChannel:
			request, ok := c.(*corev2.CheckRequest)
			if !ok {
				logger.Error("session received non-config over check channel")
				continue
			}

			configBytes, err := s.marshal(request)
			if err != nil {
				logger.WithError(err).Error("session failed to serialize check request")
				continue
			}

			msg = transport.NewMessage(corev2.CheckRequestType, configBytes)
		case <-s.ctx.Done():
			return
		}
		logger.WithField("payload_size", len(msg.Payload)).Debug("session - sending message")
		if err := s.conn.Send(msg); err != nil {
			switch err := err.(type) {
			case transport.ConnectionError, transport.ClosedError:
			default:
				logger.WithError(err).Error("send error")
			}
			return
		}
	}
}

// Start a Session.
// 1. Start sender
// 2. Start receiver
// 3. Start goroutine that waits for context cancellation, and shuts down service.
func (s *Session) Start() (err error) {
	sessionCounter.WithLabelValues(s.cfg.Namespace).Inc()
	s.wg = &sync.WaitGroup{}
	s.wg.Add(2)
	go s.sender()
	go s.receiver()
	go func() {
		<-s.ctx.Done()
		s.stop()
	}()

	namespace := s.cfg.Namespace
	agentName := fmt.Sprintf("%s:%s-%s", namespace, s.cfg.AgentName, uuid.New().String())

	defer func() {
		if err != nil {
			s.cancel()
		}
	}()

	for _, sub := range s.cfg.Subscriptions {
		// Ignore empty subscriptions
		if sub == "" {
			continue
		}

		topic := messaging.SubscriptionTopic(namespace, sub)
		logger.WithField("topic", topic).Debug("subscribing to topic")
		subscription, err := s.bus.Subscribe(topic, agentName, s)
		if err != nil {
			logger.WithError(err).Error("error starting subscription")
			return err
		}
		s.subscriptions <- subscription
	}
	close(s.subscriptions)

	return nil
}

// Stop a running session. This will cause the send and receive loops to
// shutdown. Blocks until the session has shutdown.
func (s *Session) Stop() {
	s.cancel()
	s.wg.Wait()
}

func (s *Session) stop() {
	defer func() {
		if err := s.conn.Close(); err != nil {
			logger.WithError(err).Error("error closing session")
		}
	}()

	sessionCounter.WithLabelValues(s.cfg.Namespace).Dec()
	s.wg.Wait()

	for sub := range s.subscriptions {
		if err := sub.Cancel(); err != nil {
			logger.WithError(err).Error("unable to unsubscribe from message bus")
		}
	}
	close(s.checkChannel)
	if s.ringPool == nil {
		// This is a bit of a hack - allow ringPool to be nil for the benefit
		// of the tests.
		return
	}
	for _, sub := range s.cfg.Subscriptions {
		ring := s.ringPool.Get(ringv2.Path(s.cfg.Namespace, sub))
		logger.WithFields(logrus.Fields{
			"namespace": s.cfg.Namespace,
			"agent":     s.cfg.AgentName,
		}).Info("removing agent from ring")
		if err := ring.Remove(context.Background(), s.cfg.AgentName); err != nil {
			logger.WithError(err).Error("unable to remove agent from ring")
		}
	}
}

// handleKeepalive is the keepalive message handler.
func (s *Session) handleKeepalive(ctx context.Context, payload []byte) error {
	keepalive := &corev2.Event{}
	err := s.unmarshal(payload, keepalive)
	if err != nil {
		return err
	}

	if err := keepalive.Validate(); err != nil {
		return err
	}

	// Not done by event.Validate()
	if keepalive.Timestamp == 0 {
		return errors.New("keepalive contains invalid timestamp")
	}

	keepalive.Entity.Subscriptions = addEntitySubscription(keepalive.Entity.Name, keepalive.Entity.Subscriptions)

	return s.bus.Publish(messaging.TopicKeepalive, keepalive)
}

// handleEvent is the event message handler.
func (s *Session) handleEvent(ctx context.Context, payload []byte) error {
	// Decode the payload to an event
	event := &corev2.Event{}
	if err := s.unmarshal(payload, event); err != nil {
		return err
	}

	// Validate the received event
	if err := event.Validate(); err != nil {
		return err
	}

	// Verify if we have a source in the event and if so, use it as the entity by
	// creating or retrieving it from the store
	if event.HasCheck() {
		if err := getProxyEntity(event, s.store); err != nil {
			return err
		}
	}

	// Add the entity subscription to the subscriptions of this entity
	event.Entity.Subscriptions = addEntitySubscription(event.Entity.Name, event.Entity.Subscriptions)

	return s.bus.Publish(messaging.TopicEventRaw, event)
}
