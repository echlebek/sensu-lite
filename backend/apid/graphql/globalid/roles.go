package globalid

import "github.com/echlebek/sensu-lite/types"

//
// Roles
//
var roleName = "roles"

// RoleTranslator global ID resource
var RoleTranslator = commonTranslator{
	name:       roleName,
	encodeFunc: standardEncoder(roleName, "Name"),
	decodeFunc: standardDecoder,
	isResponsibleFunc: func(record interface{}) bool {
		_, ok := record.(*types.Role)
		return ok
	},
}

// Register entity encoder/decoder
func init() { registerTranslator(RoleTranslator) }
