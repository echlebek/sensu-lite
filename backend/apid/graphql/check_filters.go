package graphql

import (
	v2 "github.com/echlebek/sensu-lite/api/core/v2"
	"github.com/echlebek/sensu-lite/backend/apid/graphql/filter"
	"github.com/echlebek/sensu-lite/util/strings"
)

// CheckFilters returns collection of filters used for matching resources.
func CheckFilters() map[string]filter.Filter {
	filters := map[string]filter.Filter{
		// published:true | published:false
		"published": filter.Boolean(func(res v2.Resource, v bool) bool {
			return res.(*v2.CheckConfig).Publish == v
		}),
		// subscription:unix | subscription:db
		"subscription": filter.String(func(res v2.Resource, v string) bool {
			return strings.InArray(v, res.(*v2.CheckConfig).Subscriptions)
		}),
	}

	// merge global filters
	for k, f := range GlobalFilters {
		filters[k] = f
	}

	return filters
}
