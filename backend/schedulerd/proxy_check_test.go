package schedulerd

import (
	"reflect"
	"testing"

	time "github.com/echlebek/timeproxy"
	corev2 "github.com/echlebek/sensu-lite/api/core/v2"
	"github.com/echlebek/sensu-lite/backend/store/cache"
	"github.com/stretchr/testify/assert"
)

func TestMatchEntities(t *testing.T) {
	entity1 := &corev2.Entity{
		ObjectMeta: corev2.ObjectMeta{
			Name:      "entity1",
			Namespace: "default",
			Labels:    map[string]string{"proxy_type": "switch"},
		},
		EntityClass: "proxy",
		System:      corev2.System{Hostname: "foo.local"},
	}
	entity2 := &corev2.Entity{
		ObjectMeta: corev2.ObjectMeta{
			Name:      "entity2",
			Namespace: "default",
			Labels:    map[string]string{"proxy_type": "sensor"},
		},
		Deregister:  true,
		EntityClass: "proxy",
	}
	entity3 := &corev2.Entity{
		ObjectMeta: corev2.ObjectMeta{
			Name:      "entity3",
			Namespace: "default",
		},
		EntityClass: "agent",
	}

	tests := []struct {
		name             string
		entityAttributes []string
		entities         []corev2.Resource
		want             []*corev2.Entity
	}{
		{
			name:             "standard string attribute",
			entityAttributes: []string{`entity.name == "entity1"`},
			entities:         []corev2.Resource{entity1, entity2, entity3},
			want:             []*corev2.Entity{entity1},
		},
		{
			name:             "standard bool attribute",
			entityAttributes: []string{`entity.deregister == true`},
			entities:         []corev2.Resource{entity1, entity2, entity3},
			want:             []*corev2.Entity{entity2},
		},
		{
			name:             "nested standard attribute",
			entityAttributes: []string{`entity.system.hostname == "foo.local"`},
			entities:         []corev2.Resource{entity1, entity2, entity3},
			want:             []*corev2.Entity{entity1},
		},
		{
			name:             "multiple matches",
			entityAttributes: []string{`entity.entity_class == "proxy"`},
			entities:         []corev2.Resource{entity1, entity2, entity3},
			want:             []*corev2.Entity{entity1, entity2},
		},
		{
			name:             "invalid expression",
			entityAttributes: []string{`foo &&`},
			entities:         []corev2.Resource{entity1, entity2, entity3},
		},
		{
			name: "multiple entity attributes",
			entityAttributes: []string{
				`entity.entity_class == "proxy"`,
				`entity.labels.proxy_type == "sensor"`,
			},
			entities: []corev2.Resource{entity1, entity2, entity3},
			want:     []*corev2.Entity{entity2},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := &corev2.ProxyRequests{
				EntityAttributes: tc.entityAttributes,
			}
			cacher := cache.NewFromResources(tc.entities, true)
			got := matchEntities(cacher.Get("default"), p)

			if len(got) != len(tc.want) {
				t.Errorf("Expected %d entities, got %d", len(tc.want), len(got))
				return
			}

			for i := range tc.want {
				if !reflect.DeepEqual(got[i], tc.want[i]) {
					t.Errorf("MatchEntities() = %v, want %v", got, tc.want)
					return
				}
			}

		})
	}
}

func TestSplayCalculation(t *testing.T) {
	assert := assert.New(t)

	check := corev2.FixtureCheckConfig("check1")
	check.ProxyRequests = corev2.FixtureProxyRequests(true)

	// 10s * 90% / 3 = 3
	check.Interval = 10
	splay, err := calculateSplayInterval(check, 3)
	assert.Equal(3*time.Second, splay)
	assert.Nil(err)

	// 20s * 50% / 5 = 2
	check.Interval = 20
	check.ProxyRequests.SplayCoverage = 50
	splay, err = calculateSplayInterval(check, 5)
	assert.Equal(2*time.Second, splay)
	assert.Nil(err)

	// invalid cron string
	check.Cron = "invalid"
	splay, err = calculateSplayInterval(check, 5)
	assert.Equal(time.Duration(0), splay)
	assert.NotNil(err)

	// at most, 60s from current time * 50% / 2 = 15
	// this test will depend on when it is run, but the
	// largest splay calculation will be 15
	check.Cron = "* * * * *"
	splay, err = calculateSplayInterval(check, 2)
	assert.True(splay >= 0 && splay <= 15*time.Second)
	assert.Nil(err)
}

func TestSubstituteProxyEntityTokens(t *testing.T) {
	assert := assert.New(t)

	entity := corev2.FixtureEntity("entity1")
	check := corev2.FixtureCheckConfig("check1")
	check.Subscriptions = []string{"subscription1"}
	check.ProxyRequests = corev2.FixtureProxyRequests(true)

	substitutedProxyEntityTokens, err := substituteProxyEntityTokens(entity, check)
	if err != nil {
		assert.FailNow(err.Error())
	}
	assert.Equal(entity.Name, substitutedProxyEntityTokens.ProxyEntityName)
}

func BenchmarkMatchEntities1000(b *testing.B) {
	entity := corev2.FixtureEntity("foo")
	// non-matching expression to avoid short-circuiting behaviour
	expression := "entity.system.arch == 'amd65'"

	entities := make([]corev2.Resource, 100)
	expressions := make([]string, 10)

	for i := range entities {
		entities[i] = entity
	}
	for i := range expressions {
		expressions[i] = expression
	}

	req := &corev2.ProxyRequests{EntityAttributes: expressions}
	// slice := cache.MakeSliceCache(entities, true)
	cacher := cache.NewFromResources(entities, true)
	resources := cacher.Get("default")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = matchEntities(resources, req)
	}
}
