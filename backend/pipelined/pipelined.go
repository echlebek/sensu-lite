// Package pipelined provides the traditional Sensu event pipeline.
package pipelined

import (
	"sync"
	"sync/atomic"

	"github.com/echlebek/sensu-lite/asset"
	"github.com/echlebek/sensu-lite/backend/messaging"
	"github.com/echlebek/sensu-lite/backend/store"
	"github.com/echlebek/sensu-lite/command"
	"github.com/echlebek/sensu-lite/rpc"
	"github.com/echlebek/sensu-lite/types"
)

// ExtensionExecutorGetterFunc gets an ExtensionExecutor. Used to decouple
// Pipelined from gRPC.
type ExtensionExecutorGetterFunc func(*types.Extension) (rpc.ExtensionExecutor, error)

// Pipelined handles incoming Sensu events and puts them through a
// Sensu event pipeline, i.e. filter -> mutator -> handler. The Sensu
// handler configuration determines which Sensu filters and mutator
// are used.
type Pipelined struct {
	assetGetter       asset.Getter
	stopping          chan struct{}
	running           *atomic.Value
	wg                *sync.WaitGroup
	errChan           chan error
	eventChan         chan interface{}
	subscription      messaging.Subscription
	store             store.Store
	bus               messaging.MessageBus
	extensionExecutor ExtensionExecutorGetterFunc
	executor          command.Executor
	workerCount       int
}

// Config configures a Pipelined.
type Config struct {
	Store                   store.Store
	Bus                     messaging.MessageBus
	ExtensionExecutorGetter ExtensionExecutorGetterFunc
	AssetGetter             asset.Getter
	BufferSize              int
	WorkerCount             int
}

// Option is a functional option used to configure Pipelined.
type Option func(*Pipelined) error

// New creates a new Pipelined with supplied Options applied.
func New(c Config, options ...Option) (*Pipelined, error) {
	if c.BufferSize == 0 {
		c.BufferSize = 1
	}
	if c.WorkerCount == 0 {
		c.WorkerCount = 1
	}

	p := &Pipelined{
		store:             c.Store,
		bus:               c.Bus,
		extensionExecutor: c.ExtensionExecutorGetter,
		stopping:          make(chan struct{}, 1),
		running:           &atomic.Value{},
		wg:                &sync.WaitGroup{},
		errChan:           make(chan error, 1),
		eventChan:         make(chan interface{}, c.BufferSize),
		workerCount:       c.WorkerCount,
		executor:          command.NewExecutor(),
		assetGetter:       c.AssetGetter,
	}
	for _, o := range options {
		if err := o(p); err != nil {
			return nil, err
		}
	}
	return p, nil
}

// Receiver returns the event channel for pipelined.
func (p *Pipelined) Receiver() chan<- interface{} {
	return p.eventChan
}

// Start pipelined, subscribing to the "event" message bus topic to
// pass Sensu events to the pipelines for handling (goroutines).
func (p *Pipelined) Start() error {
	sub, err := p.bus.Subscribe(messaging.TopicEvent, "pipelined", p)
	if err != nil {
		return err
	}
	p.subscription = sub

	p.createPipelines(p.workerCount, p.eventChan)

	return nil
}

// Stop pipelined.
func (p *Pipelined) Stop() error {
	p.running.Store(false)
	close(p.stopping)
	p.wg.Wait()
	close(p.errChan)
	err := p.subscription.Cancel()
	close(p.eventChan)

	return err
}

// Err returns a channel to listen for terminal errors on.
func (p *Pipelined) Err() <-chan error {
	return p.errChan
}

// Name returns the daemon name
func (p *Pipelined) Name() string {
	return "pipelined"
}

// createPipelines creates several goroutines, responsible for pulling
// Sensu events from a channel (bound to message bus "event" topic)
// and for handling them.
func (p *Pipelined) createPipelines(count int, channel chan interface{}) {
	for i := 1; i <= count; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			for {
				select {
				case <-p.stopping:
					return
				case msg := <-channel:
					event, ok := msg.(*types.Event)
					if !ok {
						continue
					}

					if err := p.handleEvent(event); err != nil {
						logger.Error(err)
					}
				}
			}
		}()
	}
}
