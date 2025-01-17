package globalid

import "github.com/echlebek/sensu-lite/types"

//
// Checks
//

var checkName = "checks"

// CheckTranslator global ID resource
var CheckTranslator = commonTranslator{
	name:       checkName,
	encodeFunc: standardEncoder(checkName, "Name"),
	decodeFunc: standardDecoder,
	isResponsibleFunc: func(record interface{}) bool {
		_, ok := record.(*types.CheckConfig)
		return ok
	},
}

// Register entity encoder/decoder
func init() { registerTranslator(CheckTranslator) }
