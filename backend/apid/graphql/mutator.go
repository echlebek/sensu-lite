package graphql

import (
	corev2 "github.com/echlebek/sensu-lite/api/core/v2"
	"github.com/echlebek/sensu-lite/backend/apid/graphql/globalid"
	"github.com/echlebek/sensu-lite/backend/apid/graphql/schema"
	"github.com/echlebek/sensu-lite/graphql"
	"github.com/echlebek/sensu-lite/types"
)

var _ schema.MutatorFieldResolvers = (*mutatorImpl)(nil)

//
// Implement MutatorFieldResolvers
//

type mutatorImpl struct {
	schema.MutatorAliases
}

// ID implements response to request for 'id' field.
func (*mutatorImpl) ID(p graphql.ResolveParams) (string, error) {
	return globalid.MutatorTranslator.EncodeToString(p.Source), nil
}

// IsTypeOf is used to determine if a given value is associated with the type
func (*mutatorImpl) IsTypeOf(s interface{}, p graphql.IsTypeOfParams) bool {
	_, ok := s.(*types.Mutator)
	return ok
}

// ToJSON implements response to request for 'toJSON' field.
func (*mutatorImpl) ToJSON(p graphql.ResolveParams) (interface{}, error) {
	return types.WrapResource(p.Source.(corev2.Resource)), nil
}
