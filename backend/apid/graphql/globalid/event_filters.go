package globalid

import "github.com/echlebek/sensu-lite/types"

//
// Event Filters
//

var eventFilterName = "filters"

// EventFilterTranslator global ID resource
var EventFilterTranslator = commonTranslator{
	name:       eventFilterName,
	encodeFunc: standardEncoder(eventFilterName, "Name"),
	decodeFunc: standardDecoder,
	isResponsibleFunc: func(record interface{}) bool {
		_, ok := record.(*types.EventFilter)
		return ok
	},
}

// Register event filter encoder/decoder
func init() { registerTranslator(EventFilterTranslator) }
