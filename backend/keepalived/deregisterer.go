package keepalived

import (
	"context"
	"fmt"

	"github.com/echlebek/sensu-lite/api/core/v2"
	"github.com/echlebek/sensu-lite/backend/messaging"
	"github.com/echlebek/sensu-lite/backend/store"
	"github.com/echlebek/sensu-lite/types"
)

// A Deregisterer provides a mechanism for deregistering entities and
// notifying the rest of the backend when a deregistration occurs.
type Deregisterer interface {
	// Deregister an entity and return an error if there was any problem during the
	// deregistration process.
	Deregister(e *types.Entity) error
}

// Deregistration is an adapter for deregistering an entity from the store and
// publishing a deregistration event to WizardBus.
type Deregistration struct {
	EntityStore store.EntityStore
	EventStore  store.EventStore
	MessageBus  messaging.MessageBus
}

// Deregister an entity and all of its associated events.
func (d *Deregistration) Deregister(entity *types.Entity) error {
	ctx := context.WithValue(context.Background(), types.NamespaceKey, entity.Namespace)

	if err := d.EntityStore.DeleteEntity(ctx, entity); err != nil {
		return fmt.Errorf("error deleting entity in store: %s", err)
	}

	events, err := d.EventStore.GetEventsByEntity(ctx, entity.Name, &store.SelectionPredicate{})
	if err != nil {
		return fmt.Errorf("error fetching events for entity: %s", err)
	}

	for _, event := range events {
		if !event.HasCheck() {
			return fmt.Errorf("error deleting event without check")
		}

		if err := d.EventStore.DeleteEventByEntityCheck(
			ctx, entity.Name, event.Check.Name,
		); err != nil {
			return fmt.Errorf("error deleting event for entity: %s", err)
		}

		event.Check.Output = "Resolving due to entity deregistering"
		event.Check.Status = 0
		event.Check.History = []types.CheckHistory{}

		if err := d.MessageBus.Publish(messaging.TopicEvent, event); err != nil {
			return fmt.Errorf("error publishing deregistration event: %s", err)
		}
	}

	if entity.Deregistration.Handler != "" {
		deregistrationCheck := &types.Check{
			ObjectMeta:    v2.NewObjectMeta("deregistration", entity.Namespace),
			Interval:      1,
			Subscriptions: []string{""},
			Command:       "",
			Handlers:      []string{entity.Deregistration.Handler},
			Status:        1,
		}

		deregistrationEvent := &types.Event{
			Entity: entity,
			Check:  deregistrationCheck,
		}

		return d.MessageBus.Publish(messaging.TopicEvent, deregistrationEvent)
	}

	logger.WithField("entity", entity.GetName()).Info("entity deregistered")
	return nil
}
