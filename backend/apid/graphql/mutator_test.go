package graphql

import (
	"testing"

	v2 "github.com/echlebek/sensu-lite/api/core/v2"
	"github.com/echlebek/sensu-lite/graphql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMutatorTypeToJSONField(t *testing.T) {
	src := v2.FixtureMutator("name")
	imp := &mutatorImpl{}

	res, err := imp.ToJSON(graphql.ResolveParams{Source: src})
	require.NoError(t, err)
	assert.NotEmpty(t, res)
}
