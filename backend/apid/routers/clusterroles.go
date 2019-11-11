package routers

import (
	"github.com/gorilla/mux"
	corev2 "github.com/echlebek/sensu-lite/api/core/v2"
	"github.com/echlebek/sensu-lite/backend/apid/handlers"
	"github.com/echlebek/sensu-lite/backend/store"
)

// ClusterRolesRouter handles requests for ClusterRoles.
type ClusterRolesRouter struct {
	handlers handlers.Handlers
}

// NewClusterRolesRouter instantiates a new router for ClusterRoles.
func NewClusterRolesRouter(store store.ResourceStore) *ClusterRolesRouter {
	return &ClusterRolesRouter{
		handlers: handlers.Handlers{
			Resource: &corev2.ClusterRole{},
			Store:    store,
		},
	}
}

// Mount the ClusterRolesRouter on the given parent Router
func (r *ClusterRolesRouter) Mount(parent *mux.Router) {
	routes := ResourceRoute{
		Router:     parent,
		PathPrefix: "/{resource:clusterroles}",
	}

	routes.Del(r.handlers.DeleteResource)
	routes.Get(r.handlers.GetResource)
	routes.List(r.handlers.ListResources, corev2.ClusterRoleFields)
	routes.Post(r.handlers.CreateResource)
	routes.Put(r.handlers.CreateOrUpdateResource)
}
