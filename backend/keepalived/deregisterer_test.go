package keepalived

import (
	"testing"

	"github.com/echlebek/sensu-lite/backend/store"

	"github.com/echlebek/sensu-lite/backend/messaging"
	"github.com/echlebek/sensu-lite/testing/mockbus"
	"github.com/echlebek/sensu-lite/testing/mockstore"
	"github.com/echlebek/sensu-lite/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeregister(t *testing.T) {
	assert := assert.New(t)

	mockStore := &mockstore.MockStore{}
	mockBus := &mockbus.MockBus{}

	adapter := &Deregistration{
		EntityStore: mockStore,
		EventStore:  mockStore,
		MessageBus:  mockBus,
	}

	entity := types.FixtureEntity("entity")
	entity.Deregister = true
	check := types.FixtureCheck("check")
	event := types.FixtureEvent(entity.Name, check.Name)

	mockStore.On("GetEventsByEntity", mock.Anything, entity.Name, &store.SelectionPredicate{}).Return([]*types.Event{event}, nil)
	mockStore.On("DeleteEventByEntityCheck", mock.Anything, entity.Name, check.Name).Return(nil)
	mockStore.On("DeleteEntity", mock.Anything, entity).Return(nil)

	mockBus.On("Publish", mock.AnythingOfType("string"), mock.Anything).Return(nil)

	assert.NoError(adapter.Deregister(entity))
}

func TestDeregistrationHandler(t *testing.T) {
	assert := assert.New(t)

	mockStore := &mockstore.MockStore{}
	mockBus := &mockbus.MockBus{}

	adapter := &Deregistration{
		EventStore:  mockStore,
		EntityStore: mockStore,
		MessageBus:  mockBus,
	}

	entity := types.FixtureEntity("entity")
	entity.Deregister = true
	entity.Deregistration = types.Deregistration{
		Handler: "deregistration",
	}
	check := types.FixtureCheck("check")

	mockStore.On("GetEventsByEntity", mock.Anything, entity.Name, &store.SelectionPredicate{}).Return([]*types.Event{}, nil)
	mockStore.On("DeleteEventByEntityCheck", mock.Anything, entity.Name, check.Name).Return(nil)
	mockStore.On("DeleteEntity", mock.Anything, entity).Return(nil)

	mockBus.On("Publish", messaging.TopicEvent, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		event := args[1].(*types.Event)
		assert.Equal("deregistration", event.Entity.Deregistration.Handler)
	})

	assert.NoError(adapter.Deregister(entity))
}
