package graphql

import (
	corev2 "github.com/echlebek/sensu-lite/api/core/v2"
	"github.com/echlebek/sensu-lite/backend/apid/graphql/globalid"
	"github.com/echlebek/sensu-lite/backend/apid/graphql/schema"
	"github.com/echlebek/sensu-lite/graphql"
)

var _ schema.RuleFieldResolvers = (*ruleImpl)(nil)
var _ schema.RoleFieldResolvers = (*roleImpl)(nil)

//
// Implement RuleFieldResolvers
//

type ruleImpl struct {
	schema.RuleAliases
}

// IsTypeOf is used to determine if a given value is associated with the type
func (*ruleImpl) IsTypeOf(s interface{}, p graphql.IsTypeOfParams) bool {
	_, ok := s.(corev2.Rule)
	return ok
}

//
// Implement RoleFieldResolvers
//

type roleImpl struct {
	schema.RoleAliases
}

// ID implements response to request for 'id' field.
func (*roleImpl) ID(p graphql.ResolveParams) (string, error) {
	return globalid.RoleTranslator.EncodeToString(p.Source), nil
}

// IsTypeOf is used to determine if a given value is associated with the type
func (*roleImpl) IsTypeOf(s interface{}, p graphql.IsTypeOfParams) bool {
	_, ok := s.(*corev2.Role)
	return ok
}
