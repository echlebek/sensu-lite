// Code generated by scripts/gengraphql.go. DO NOT EDIT.

package schema

import (
	graphql1 "github.com/graphql-go/graphql"
	graphql "github.com/echlebek/sensu-lite/graphql"
)

// ResourceType self descriptive
var ResourceType = graphql.NewType("Resource", graphql.InterfaceKind)

// RegisterResource registers Resource object type with given service.
func RegisterResource(svc *graphql.Service, impl graphql.InterfaceTypeResolver) {
	svc.RegisterInterface(_InterfaceTypeResourceDesc, impl)
}
func _InterfaceTypeResourceConfigFn() graphql1.InterfaceConfig {
	return graphql1.InterfaceConfig{
		Description: "self descriptive",
		Fields: graphql1.Fields{
			"metadata": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Metadata contains the name, namespace, labels and annotations of the resource,",
				Name:              "metadata",
				Type:              graphql1.NewNonNull(graphql.OutputType("ObjectMeta")),
			},
			"toJSON": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "toJSON returns a REST API compatible representation of the resource. Handy for\nsharing snippets that can then be imported with `sensuctl create`.",
				Name:              "toJSON",
				Type:              graphql1.NewNonNull(graphql.OutputType("JSON")),
			},
		},
		Name: "Resource",
		ResolveType: func(_ graphql1.ResolveTypeParams) *graphql1.Object {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see InterfaceTypeResolver.")
		},
	}
}

// describe Resource's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _InterfaceTypeResourceDesc = graphql.InterfaceDesc{Config: _InterfaceTypeResourceConfigFn}
