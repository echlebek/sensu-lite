package rolebinding

import (
	"testing"

	client "github.com/echlebek/sensu-lite/cli/client/testing"
	test "github.com/echlebek/sensu-lite/cli/commands/testing"
	"github.com/stretchr/testify/assert"
)

func TestInfoCommand(t *testing.T) {
	cli := test.NewMockCLI()
	cli.Config.(*client.MockConfig).On("Format").Return("json")
	cmd := InfoCommand(cli)

	assert.NotNil(t, cmd, "cmd should be returned")
	assert.NotNil(t, cmd.RunE, "cmd should be able to be executed")
	assert.Regexp(t, "info", cmd.Use)
	assert.Regexp(t, "role binding", cmd.Short)
}
