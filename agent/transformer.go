package agent

import "github.com/echlebek/sensu-lite/types"

// A Transformer handles transforming Sensu metrics to other output metric formats
type Transformer interface {
	// Transform transforms a metric in a different output metric format to Sensu Metric
	// Format
	Transform() []*types.MetricPoint
}
