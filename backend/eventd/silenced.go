package eventd

import (
	"context"

	corev2 "github.com/echlebek/sensu-lite/api/core/v2"
	"github.com/echlebek/sensu-lite/backend/store/cache"
	stringsutil "github.com/echlebek/sensu-lite/util/strings"
)

// addToSilencedBy takes a silenced entry ID and adds it to a silence of IDs if
// it's not already present in order to avoid duplicated elements
func addToSilencedBy(id string, ids []string) []string {
	if !stringsutil.InArray(id, ids) {
		ids = append(ids, id)
	}
	return ids
}

// getSilenced retrieves all silenced entries for a given event, using the
// entity subscription, the check subscription and the check name while
// supporting wildcard silenced entries (e.g. subscription:*)
func getSilenced(ctx context.Context, event *corev2.Event, cache *cache.Resource) {
	if !event.HasCheck() {
		return
	}

	resources := cache.Get(event.Check.Namespace)
	entries := make([]*corev2.Silenced, len(resources))
	for i, resource := range resources {
		entries[i] = resource.Resource.(*corev2.Silenced)
	}

	// Determine which entries silence this event
	silencedIDs := silencedBy(event, entries)

	// Add to the event all silenced entries ID that actually silence it
	event.Check.Silenced = silencedIDs
}

// silencedBy determines which of the given silenced entries silenced a given
// event and return a list of silenced entry IDs
func silencedBy(event *corev2.Event, silencedEntries []*corev2.Silenced) []string {
	silencedBy := event.SilencedBy(silencedEntries)
	names := make([]string, 0, len(silencedBy))
	for _, entry := range silencedBy {
		names = addToSilencedBy(entry.Name, names)
	}
	return names
}
