package graphql

import (
	v2 "github.com/echlebek/sensu-lite/api/core/v2"
	"github.com/echlebek/sensu-lite/backend/apid/graphql/filter"
	"github.com/echlebek/sensu-lite/util/strings"
)

// EntityFilters returns collection of filters used for matching resources.
func EntityFilters() map[string]filter.Filter {
	filters := map[string]filter.Filter{
		// class:proxy | class:agent
		"class": filter.String(func(res v2.Resource, v string) bool {
			return res.(*v2.Entity).EntityClass == v
		}),
		// subscription:unix | subscription:db
		"subscription": filter.String(func(res v2.Resource, v string) bool {
			return strings.InArray(v, res.(*v2.Entity).Subscriptions)
		}),
	}

	// merge global filters
	for k, f := range GlobalFilters {
		filters[k] = f
	}

	return filters
}
