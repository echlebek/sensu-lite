// Code generated by scripts/gengraphql.go. DO NOT EDIT.

package schema

import (
	graphql1 "github.com/graphql-go/graphql"
	mapstructure "github.com/mitchellh/mapstructure"
	graphql "github.com/echlebek/sensu-lite/graphql"
)

// QueryViewerFieldResolver implement to resolve requests for the Query's viewer field.
type QueryViewerFieldResolver interface {
	// Viewer implements response to request for viewer field.
	Viewer(p graphql.ResolveParams) (interface{}, error)
}

// QueryNamespaceFieldResolverArgs contains arguments provided to namespace when selected
type QueryNamespaceFieldResolverArgs struct {
	Name string // Name - self descriptive
}

// QueryNamespaceFieldResolverParams contains contextual info to resolve namespace field
type QueryNamespaceFieldResolverParams struct {
	graphql.ResolveParams
	Args QueryNamespaceFieldResolverArgs
}

// QueryNamespaceFieldResolver implement to resolve requests for the Query's namespace field.
type QueryNamespaceFieldResolver interface {
	// Namespace implements response to request for namespace field.
	Namespace(p QueryNamespaceFieldResolverParams) (interface{}, error)
}

// QueryEventFieldResolverArgs contains arguments provided to event when selected
type QueryEventFieldResolverArgs struct {
	Namespace string // Namespace - self descriptive
	Entity    string // Entity - self descriptive
	Check     string // Check - self descriptive
}

// QueryEventFieldResolverParams contains contextual info to resolve event field
type QueryEventFieldResolverParams struct {
	graphql.ResolveParams
	Args QueryEventFieldResolverArgs
}

// QueryEventFieldResolver implement to resolve requests for the Query's event field.
type QueryEventFieldResolver interface {
	// Event implements response to request for event field.
	Event(p QueryEventFieldResolverParams) (interface{}, error)
}

// QueryEntityFieldResolverArgs contains arguments provided to entity when selected
type QueryEntityFieldResolverArgs struct {
	Namespace string // Namespace - self descriptive
	Name      string // Name - self descriptive
}

// QueryEntityFieldResolverParams contains contextual info to resolve entity field
type QueryEntityFieldResolverParams struct {
	graphql.ResolveParams
	Args QueryEntityFieldResolverArgs
}

// QueryEntityFieldResolver implement to resolve requests for the Query's entity field.
type QueryEntityFieldResolver interface {
	// Entity implements response to request for entity field.
	Entity(p QueryEntityFieldResolverParams) (interface{}, error)
}

// QueryMutatorFieldResolverArgs contains arguments provided to mutator when selected
type QueryMutatorFieldResolverArgs struct {
	Namespace string // Namespace - self descriptive
	Name      string // Name - self descriptive
}

// QueryMutatorFieldResolverParams contains contextual info to resolve mutator field
type QueryMutatorFieldResolverParams struct {
	graphql.ResolveParams
	Args QueryMutatorFieldResolverArgs
}

// QueryMutatorFieldResolver implement to resolve requests for the Query's mutator field.
type QueryMutatorFieldResolver interface {
	// Mutator implements response to request for mutator field.
	Mutator(p QueryMutatorFieldResolverParams) (interface{}, error)
}

// QueryCheckFieldResolverArgs contains arguments provided to check when selected
type QueryCheckFieldResolverArgs struct {
	Namespace string // Namespace - self descriptive
	Name      string // Name - self descriptive
}

// QueryCheckFieldResolverParams contains contextual info to resolve check field
type QueryCheckFieldResolverParams struct {
	graphql.ResolveParams
	Args QueryCheckFieldResolverArgs
}

// QueryCheckFieldResolver implement to resolve requests for the Query's check field.
type QueryCheckFieldResolver interface {
	// Check implements response to request for check field.
	Check(p QueryCheckFieldResolverParams) (interface{}, error)
}

// QueryEventFilterFieldResolverArgs contains arguments provided to eventFilter when selected
type QueryEventFilterFieldResolverArgs struct {
	Namespace string // Namespace - self descriptive
	Name      string // Name - self descriptive
}

// QueryEventFilterFieldResolverParams contains contextual info to resolve eventFilter field
type QueryEventFilterFieldResolverParams struct {
	graphql.ResolveParams
	Args QueryEventFilterFieldResolverArgs
}

// QueryEventFilterFieldResolver implement to resolve requests for the Query's eventFilter field.
type QueryEventFilterFieldResolver interface {
	// EventFilter implements response to request for eventFilter field.
	EventFilter(p QueryEventFilterFieldResolverParams) (interface{}, error)
}

// QueryHandlerFieldResolverArgs contains arguments provided to handler when selected
type QueryHandlerFieldResolverArgs struct {
	Namespace string // Namespace - self descriptive
	Name      string // Name - self descriptive
}

// QueryHandlerFieldResolverParams contains contextual info to resolve handler field
type QueryHandlerFieldResolverParams struct {
	graphql.ResolveParams
	Args QueryHandlerFieldResolverArgs
}

// QueryHandlerFieldResolver implement to resolve requests for the Query's handler field.
type QueryHandlerFieldResolver interface {
	// Handler implements response to request for handler field.
	Handler(p QueryHandlerFieldResolverParams) (interface{}, error)
}

// QuerySuggestFieldResolverArgs contains arguments provided to suggest when selected
type QuerySuggestFieldResolverArgs struct {
	Q string /*
	Q - If the value of a field does not contain the value of this argument it will
	be omitted from the response. Operation is case-insensitive.
	*/
	Ref string /*
	Ref is used to uniquely identify a resource in the system as well as a field
	on said resource. Refs take the form: :group/:version/:type/:field. The
	field segment may be a path in and of it's own, eg. metadata/name would
	refer to the name field nested inside a resource's metadata.

	The following are valid example values for this argument:

	    `core/v2/asset/metadata/name`
	    `core/v2/asset/metadata/labels`
	    `core/v2/asset/metadata/labels/region`
	    `core/v2/check_config/subscriptions`
	    `core/v2/check_config/command`
	    `core/v2/check_config/timeout`
	    `core/v2/entity/system/os`
	    `core/v2/entity/system/platform`
	    `core/v2/filter/metadata/name`
	    `core/v2/handler/command`
	    `core/v2/hook_config/command`
	    `core/v2/mutator/command`
	    `core/v2/mutator/timeout`
	    `core/v2/silenced/creator`
	*/
	Namespace string          // Namespace - self descriptive
	Limit     int             // Limit - self descriptive
	Order     SuggestionOrder // Order - self descriptive
}

// QuerySuggestFieldResolverParams contains contextual info to resolve suggest field
type QuerySuggestFieldResolverParams struct {
	graphql.ResolveParams
	Args QuerySuggestFieldResolverArgs
}

// QuerySuggestFieldResolver implement to resolve requests for the Query's suggest field.
type QuerySuggestFieldResolver interface {
	// Suggest implements response to request for suggest field.
	Suggest(p QuerySuggestFieldResolverParams) (interface{}, error)
}

// QueryNodeFieldResolverArgs contains arguments provided to node when selected
type QueryNodeFieldResolverArgs struct {
	ID string // ID - The ID of an object.
}

// QueryNodeFieldResolverParams contains contextual info to resolve node field
type QueryNodeFieldResolverParams struct {
	graphql.ResolveParams
	Args QueryNodeFieldResolverArgs
}

// QueryNodeFieldResolver implement to resolve requests for the Query's node field.
type QueryNodeFieldResolver interface {
	// Node implements response to request for node field.
	Node(p QueryNodeFieldResolverParams) (interface{}, error)
}

// QueryWrappedNodeFieldResolverArgs contains arguments provided to wrappedNode when selected
type QueryWrappedNodeFieldResolverArgs struct {
	ID string // ID - The ID of an object.
}

// QueryWrappedNodeFieldResolverParams contains contextual info to resolve wrappedNode field
type QueryWrappedNodeFieldResolverParams struct {
	graphql.ResolveParams
	Args QueryWrappedNodeFieldResolverArgs
}

// QueryWrappedNodeFieldResolver implement to resolve requests for the Query's wrappedNode field.
type QueryWrappedNodeFieldResolver interface {
	// WrappedNode implements response to request for wrappedNode field.
	WrappedNode(p QueryWrappedNodeFieldResolverParams) (interface{}, error)
}

//
// QueryFieldResolvers represents a collection of methods whose products represent the
// response values of the 'Query' type.
//
// == Example SDL
//
//   """
//   Dog's are not hooman.
//   """
//   type Dog implements Pet {
//     "name of this fine beast."
//     name:  String!
//
//     "breed of this silly animal; probably shibe."
//     breed: [Breed]
//   }
//
// == Example generated interface
//
//   // DogResolver ...
//   type DogFieldResolvers interface {
//     DogNameFieldResolver
//     DogBreedFieldResolver
//
//     // IsTypeOf is used to determine if a given value is associated with the Dog type
//     IsTypeOf(interface{}, graphql.IsTypeOfParams) bool
//   }
//
// == Example implementation ...
//
//   // DogResolver implements DogFieldResolvers interface
//   type DogResolver struct {
//     logger logrus.LogEntry
//     store interface{
//       store.BreedStore
//       store.DogStore
//     }
//   }
//
//   // Name implements response to request for name field.
//   func (r *DogResolver) Name(p graphql.ResolveParams) (interface{}, error) {
//     // ... implementation details ...
//     dog := p.Source.(DogGetter)
//     return dog.GetName()
//   }
//
//   // Breed implements response to request for breed field.
//   func (r *DogResolver) Breed(p graphql.ResolveParams) (interface{}, error) {
//     // ... implementation details ...
//     dog := p.Source.(DogGetter)
//     breed := r.store.GetBreed(dog.GetBreedName())
//     return breed
//   }
//
//   // IsTypeOf is used to determine if a given value is associated with the Dog type
//   func (r *DogResolver) IsTypeOf(p graphql.IsTypeOfParams) bool {
//     // ... implementation details ...
//     _, ok := p.Value.(DogGetter)
//     return ok
//   }
//
type QueryFieldResolvers interface {
	QueryViewerFieldResolver
	QueryNamespaceFieldResolver
	QueryEventFieldResolver
	QueryEntityFieldResolver
	QueryMutatorFieldResolver
	QueryCheckFieldResolver
	QueryEventFilterFieldResolver
	QueryHandlerFieldResolver
	QuerySuggestFieldResolver
	QueryNodeFieldResolver
	QueryWrappedNodeFieldResolver
}

// QueryAliases implements all methods on QueryFieldResolvers interface by using reflection to
// match name of field to a field on the given value. Intent is reduce friction
// of writing new resolvers by removing all the instances where you would simply
// have the resolvers method return a field.
//
// == Example SDL
//
//    type Dog {
//      name:   String!
//      weight: Float!
//      dob:    DateTime
//      breed:  [Breed]
//    }
//
// == Example generated aliases
//
//   type DogAliases struct {}
//   func (_ DogAliases) Name(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Weight(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Dob(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Breed(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//
// == Example Implementation
//
//   type DogResolver struct { // Implements DogResolver
//     DogAliases
//     store store.BreedStore
//   }
//
//   // NOTE:
//   // All other fields are satisified by DogAliases but since this one
//   // requires hitting the store we implement it in our resolver.
//   func (r *DogResolver) Breed(p graphql.ResolveParams) interface{} {
//     dog := v.(*Dog)
//     return r.BreedsById(dog.BreedIDs)
//   }
//
type QueryAliases struct{}

// Viewer implements response to request for 'viewer' field.
func (_ QueryAliases) Viewer(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// Namespace implements response to request for 'namespace' field.
func (_ QueryAliases) Namespace(p QueryNamespaceFieldResolverParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// Event implements response to request for 'event' field.
func (_ QueryAliases) Event(p QueryEventFieldResolverParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// Entity implements response to request for 'entity' field.
func (_ QueryAliases) Entity(p QueryEntityFieldResolverParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// Mutator implements response to request for 'mutator' field.
func (_ QueryAliases) Mutator(p QueryMutatorFieldResolverParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// Check implements response to request for 'check' field.
func (_ QueryAliases) Check(p QueryCheckFieldResolverParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// EventFilter implements response to request for 'eventFilter' field.
func (_ QueryAliases) EventFilter(p QueryEventFilterFieldResolverParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// Handler implements response to request for 'handler' field.
func (_ QueryAliases) Handler(p QueryHandlerFieldResolverParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// Suggest implements response to request for 'suggest' field.
func (_ QueryAliases) Suggest(p QuerySuggestFieldResolverParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// Node implements response to request for 'node' field.
func (_ QueryAliases) Node(p QueryNodeFieldResolverParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// WrappedNode implements response to request for 'wrappedNode' field.
func (_ QueryAliases) WrappedNode(p QueryWrappedNodeFieldResolverParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// QueryType The query root of Sensu's GraphQL interface.
var QueryType = graphql.NewType("Query", graphql.ObjectKind)

// RegisterQuery registers Query object type with given service.
func RegisterQuery(svc *graphql.Service, impl QueryFieldResolvers) {
	svc.RegisterObject(_ObjectTypeQueryDesc, impl)
}
func _ObjTypeQueryViewerHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(QueryViewerFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Viewer(frp)
	}
}

func _ObjTypeQueryNamespaceHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(QueryNamespaceFieldResolver)
	return func(p graphql1.ResolveParams) (interface{}, error) {
		frp := QueryNamespaceFieldResolverParams{ResolveParams: p}
		err := mapstructure.Decode(p.Args, &frp.Args)
		if err != nil {
			return nil, err
		}

		return resolver.Namespace(frp)
	}
}

func _ObjTypeQueryEventHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(QueryEventFieldResolver)
	return func(p graphql1.ResolveParams) (interface{}, error) {
		frp := QueryEventFieldResolverParams{ResolveParams: p}
		err := mapstructure.Decode(p.Args, &frp.Args)
		if err != nil {
			return nil, err
		}

		return resolver.Event(frp)
	}
}

func _ObjTypeQueryEntityHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(QueryEntityFieldResolver)
	return func(p graphql1.ResolveParams) (interface{}, error) {
		frp := QueryEntityFieldResolverParams{ResolveParams: p}
		err := mapstructure.Decode(p.Args, &frp.Args)
		if err != nil {
			return nil, err
		}

		return resolver.Entity(frp)
	}
}

func _ObjTypeQueryMutatorHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(QueryMutatorFieldResolver)
	return func(p graphql1.ResolveParams) (interface{}, error) {
		frp := QueryMutatorFieldResolverParams{ResolveParams: p}
		err := mapstructure.Decode(p.Args, &frp.Args)
		if err != nil {
			return nil, err
		}

		return resolver.Mutator(frp)
	}
}

func _ObjTypeQueryCheckHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(QueryCheckFieldResolver)
	return func(p graphql1.ResolveParams) (interface{}, error) {
		frp := QueryCheckFieldResolverParams{ResolveParams: p}
		err := mapstructure.Decode(p.Args, &frp.Args)
		if err != nil {
			return nil, err
		}

		return resolver.Check(frp)
	}
}

func _ObjTypeQueryEventFilterHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(QueryEventFilterFieldResolver)
	return func(p graphql1.ResolveParams) (interface{}, error) {
		frp := QueryEventFilterFieldResolverParams{ResolveParams: p}
		err := mapstructure.Decode(p.Args, &frp.Args)
		if err != nil {
			return nil, err
		}

		return resolver.EventFilter(frp)
	}
}

func _ObjTypeQueryHandlerHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(QueryHandlerFieldResolver)
	return func(p graphql1.ResolveParams) (interface{}, error) {
		frp := QueryHandlerFieldResolverParams{ResolveParams: p}
		err := mapstructure.Decode(p.Args, &frp.Args)
		if err != nil {
			return nil, err
		}

		return resolver.Handler(frp)
	}
}

func _ObjTypeQuerySuggestHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(QuerySuggestFieldResolver)
	return func(p graphql1.ResolveParams) (interface{}, error) {
		frp := QuerySuggestFieldResolverParams{ResolveParams: p}
		err := mapstructure.Decode(p.Args, &frp.Args)
		if err != nil {
			return nil, err
		}

		return resolver.Suggest(frp)
	}
}

func _ObjTypeQueryNodeHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(QueryNodeFieldResolver)
	return func(p graphql1.ResolveParams) (interface{}, error) {
		frp := QueryNodeFieldResolverParams{ResolveParams: p}
		err := mapstructure.Decode(p.Args, &frp.Args)
		if err != nil {
			return nil, err
		}

		return resolver.Node(frp)
	}
}

func _ObjTypeQueryWrappedNodeHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(QueryWrappedNodeFieldResolver)
	return func(p graphql1.ResolveParams) (interface{}, error) {
		frp := QueryWrappedNodeFieldResolverParams{ResolveParams: p}
		err := mapstructure.Decode(p.Args, &frp.Args)
		if err != nil {
			return nil, err
		}

		return resolver.WrappedNode(frp)
	}
}

func _ObjectTypeQueryConfigFn() graphql1.ObjectConfig {
	return graphql1.ObjectConfig{
		Description: "The query root of Sensu's GraphQL interface.",
		Fields: graphql1.Fields{
			"check": &graphql1.Field{
				Args: graphql1.FieldConfigArgument{
					"name": &graphql1.ArgumentConfig{
						Description: "self descriptive",
						Type:        graphql1.NewNonNull(graphql1.String),
					},
					"namespace": &graphql1.ArgumentConfig{
						Description: "self descriptive",
						Type:        graphql1.NewNonNull(graphql1.String),
					},
				},
				DeprecationReason: "",
				Description:       "check fetches the check config associated with the given set of arguments.",
				Name:              "check",
				Type:              graphql.OutputType("CheckConfig"),
			},
			"entity": &graphql1.Field{
				Args: graphql1.FieldConfigArgument{
					"name": &graphql1.ArgumentConfig{
						Description: "self descriptive",
						Type:        graphql1.NewNonNull(graphql1.String),
					},
					"namespace": &graphql1.ArgumentConfig{
						Description: "self descriptive",
						Type:        graphql1.NewNonNull(graphql1.String),
					},
				},
				DeprecationReason: "",
				Description:       "Entity fetches the entity associated with the given set of arguments.",
				Name:              "entity",
				Type:              graphql.OutputType("Entity"),
			},
			"event": &graphql1.Field{
				Args: graphql1.FieldConfigArgument{
					"check": &graphql1.ArgumentConfig{
						Description: "self descriptive",
						Type:        graphql1.String,
					},
					"entity": &graphql1.ArgumentConfig{
						Description: "self descriptive",
						Type:        graphql1.NewNonNull(graphql1.String),
					},
					"namespace": &graphql1.ArgumentConfig{
						Description: "self descriptive",
						Type:        graphql1.NewNonNull(graphql1.String),
					},
				},
				DeprecationReason: "",
				Description:       "Event fetches the event associated with the given set of arguments.",
				Name:              "event",
				Type:              graphql.OutputType("Event"),
			},
			"eventFilter": &graphql1.Field{
				Args: graphql1.FieldConfigArgument{
					"name": &graphql1.ArgumentConfig{
						Description: "self descriptive",
						Type:        graphql1.NewNonNull(graphql1.String),
					},
					"namespace": &graphql1.ArgumentConfig{
						Description: "self descriptive",
						Type:        graphql1.NewNonNull(graphql1.String),
					},
				},
				DeprecationReason: "",
				Description:       "eventFilter fetches the event filter associated with the given set of arguments.",
				Name:              "eventFilter",
				Type:              graphql.OutputType("EventFilter"),
			},
			"handler": &graphql1.Field{
				Args: graphql1.FieldConfigArgument{
					"name": &graphql1.ArgumentConfig{
						Description: "self descriptive",
						Type:        graphql1.NewNonNull(graphql1.String),
					},
					"namespace": &graphql1.ArgumentConfig{
						Description: "self descriptive",
						Type:        graphql1.NewNonNull(graphql1.String),
					},
				},
				DeprecationReason: "",
				Description:       "handler fetch the handler associated with the given set of arguments.",
				Name:              "handler",
				Type:              graphql.OutputType("Handler"),
			},
			"mutator": &graphql1.Field{
				Args: graphql1.FieldConfigArgument{
					"name": &graphql1.ArgumentConfig{
						Description: "self descriptive",
						Type:        graphql1.NewNonNull(graphql1.String),
					},
					"namespace": &graphql1.ArgumentConfig{
						Description: "self descriptive",
						Type:        graphql1.NewNonNull(graphql1.String),
					},
				},
				DeprecationReason: "",
				Description:       "Mutator fetch the mutator associated with the given set of arguments.",
				Name:              "mutator",
				Type:              graphql.OutputType("Mutator"),
			},
			"namespace": &graphql1.Field{
				Args: graphql1.FieldConfigArgument{"name": &graphql1.ArgumentConfig{
					Description: "self descriptive",
					Type:        graphql1.NewNonNull(graphql1.String),
				}},
				DeprecationReason: "",
				Description:       "Namespace fetches the namespace object associated with the given name.",
				Name:              "namespace",
				Type:              graphql.OutputType("Namespace"),
			},
			"node": &graphql1.Field{
				Args: graphql1.FieldConfigArgument{"id": &graphql1.ArgumentConfig{
					Description: "The ID of an object.",
					Type:        graphql1.NewNonNull(graphql1.ID),
				}},
				DeprecationReason: "",
				Description:       "Node fetches an object given its ID.",
				Name:              "node",
				Type:              graphql.OutputType("Node"),
			},
			"suggest": &graphql1.Field{
				Args: graphql1.FieldConfigArgument{
					"limit": &graphql1.ArgumentConfig{
						DefaultValue: 10,
						Description:  "self descriptive",
						Type:         graphql1.Int,
					},
					"namespace": &graphql1.ArgumentConfig{
						Description: "self descriptive",
						Type:        graphql1.NewNonNull(graphql1.String),
					},
					"order": &graphql1.ArgumentConfig{
						DefaultValue: "FREQUENCY",
						Description:  "self descriptive",
						Type:         graphql.InputType("SuggestionOrder"),
					},
					"q": &graphql1.ArgumentConfig{
						DefaultValue: "",
						Description:  "If the value of a field does not contain the value of this argument it will\nbe omitted from the response. Operation is case-insensitive.",
						Type:         graphql1.String,
					},
					"ref": &graphql1.ArgumentConfig{
						Description: "Ref is used to uniquely identify a resource in the system as well as a field\non said resource. Refs take the form: :group/:version/:type/:field. The\nfield segment may be a path in and of it's own, eg. metadata/name would\nrefer to the name field nested inside a resource's metadata.\n\nThe following are valid example values for this argument:\n\n    `core/v2/asset/metadata/name`\n    `core/v2/asset/metadata/labels`\n    `core/v2/asset/metadata/labels/region`\n    `core/v2/check_config/subscriptions`\n    `core/v2/check_config/command`\n    `core/v2/check_config/timeout`\n    `core/v2/entity/system/os`\n    `core/v2/entity/system/platform`\n    `core/v2/filter/metadata/name`\n    `core/v2/handler/command`\n    `core/v2/hook_config/command`\n    `core/v2/mutator/command`\n    `core/v2/mutator/timeout`\n    `core/v2/silenced/creator`",
						Type:        graphql1.NewNonNull(graphql1.String),
					},
				},
				DeprecationReason: "",
				Description:       "Given a ref, field and a namespace returns a set of suggested values.\n\nAs an example if you would like a list of check names you might use:\n`suggest(ref: \"core/v2/check_config/metadata/name\", namespace: \"default\")`\n\nOr, if you would like a list of subscriptions...\n`suggest(ref: \"core/v2/entity/subscriptions\", namespace: \"default\")`\n\nYou may filter the results with the `q` argument, for example:\n`suggest(ref: \"core/v2/check_config/metadata/name\", namespace: \"default\", q: \"disk\")`\n\nBy default the results are ordered by the frequency in which the result occurs in the set. The `order` argument allow you to tweak this behaviour, for example:\n`suggest(ref: \"core/v2/check_config/metadata/name\", namespace: \"default\", order: ALPHA_DESC)`",
				Name:              "suggest",
				Type:              graphql.OutputType("SuggestionResultSet"),
			},
			"viewer": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Current viewer.",
				Name:              "viewer",
				Type:              graphql.OutputType("Viewer"),
			},
			"wrappedNode": &graphql1.Field{
				Args: graphql1.FieldConfigArgument{"id": &graphql1.ArgumentConfig{
					Description: "The ID of an object.",
					Type:        graphql1.NewNonNull(graphql1.ID),
				}},
				DeprecationReason: "",
				Description:       "Node fetches an object given its ID and returns it as wrapped resource.",
				Name:              "wrappedNode",
				Type:              graphql.OutputType("JSON"),
			},
		},
		Interfaces: []*graphql1.Interface{},
		IsTypeOf: func(_ graphql1.IsTypeOfParams) bool {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see QueryFieldResolvers.")
		},
		Name: "Query",
	}
}

// describe Query's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _ObjectTypeQueryDesc = graphql.ObjectDesc{
	Config: _ObjectTypeQueryConfigFn,
	FieldHandlers: map[string]graphql.FieldHandler{
		"check":       _ObjTypeQueryCheckHandler,
		"entity":      _ObjTypeQueryEntityHandler,
		"event":       _ObjTypeQueryEventHandler,
		"eventFilter": _ObjTypeQueryEventFilterHandler,
		"handler":     _ObjTypeQueryHandlerHandler,
		"mutator":     _ObjTypeQueryMutatorHandler,
		"namespace":   _ObjTypeQueryNamespaceHandler,
		"node":        _ObjTypeQueryNodeHandler,
		"suggest":     _ObjTypeQuerySuggestHandler,
		"viewer":      _ObjTypeQueryViewerHandler,
		"wrappedNode": _ObjTypeQueryWrappedNodeHandler,
	},
}
