package routers

import (
	"github.com/gorilla/mux"
	corev2 "github.com/echlebek/sensu-lite/api/core/v2"
	"github.com/echlebek/sensu-lite/backend/apid/handlers"
	"github.com/echlebek/sensu-lite/backend/store"
)

// EventFiltersRouter handles /filters requests.
type EventFiltersRouter struct {
	handlers handlers.Handlers
}

// NewEventFiltersRouter creates a new EventFiltersRouter.
func NewEventFiltersRouter(store store.ResourceStore) *EventFiltersRouter {
	return &EventFiltersRouter{
		handlers: handlers.Handlers{
			Resource: &corev2.EventFilter{},
			Store:    store,
		},
	}
}

// Mount the EventFiltersRouter to a parent Router
func (r *EventFiltersRouter) Mount(parent *mux.Router) {
	routes := ResourceRoute{
		Router:     parent,
		PathPrefix: "/namespaces/{namespace}/{resource:filters}",
	}

	routes.Del(r.handlers.DeleteResource)
	routes.Get(r.handlers.GetResource)
	routes.List(r.handlers.ListResources, corev2.EventFilterFields)
	routes.ListAllNamespaces(r.handlers.ListResources, "/{resource:filters}", corev2.EventFilterFields)
	routes.Post(r.handlers.CreateResource)
	routes.Put(r.handlers.CreateOrUpdateResource)
}
