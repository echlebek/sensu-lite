package silenced

import (
	"errors"
	"testing"

	client "github.com/echlebek/sensu-lite/cli/client/testing"
	test "github.com/echlebek/sensu-lite/cli/commands/testing"
	"github.com/echlebek/sensu-lite/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestInfoCommand(t *testing.T) {
	assert := assert.New(t)

	cli := test.NewCLI()
	cmd := InfoCommand(cli)

	assert.NotNil(cmd, "cmd should be returned")
	assert.NotNil(cmd.RunE, "cmd should be able to be executed")
	assert.Regexp("info", cmd.Use)
	assert.Regexp("silenced", cmd.Short)
}

func TestInfoCommandRunEClosure(t *testing.T) {
	assert := assert.New(t)

	cli := test.NewCLI()
	client := cli.Client.(*client.MockClient)
	client.On("FetchSilenced", mock.Anything).Return(types.FixtureSilenced("foo:bar"), nil)

	cmd := InfoCommand(cli)
	out, err := test.RunCmd(cmd, []string{"foo:bar"})
	require.NoError(t, err)

	assert.NotEmpty(out)
	assert.Contains(out, "foo:bar")
}

func TestInfoCommandRunMissingArgs(t *testing.T) {
	assert := assert.New(t)

	cli := test.NewCLI()
	cmd := InfoCommand(cli)
	out, err := test.RunCmd(cmd, []string{"wrong", "stuff"})

	require.Error(t, err)
	assert.NotEmpty(out)
	assert.Contains(out, "Usage")
}

func TestInfoCommandRunEClosureWithTable(t *testing.T) {
	assert := assert.New(t)

	cli := test.NewCLI()
	client := cli.Client.(*client.MockClient)
	client.On("FetchSilenced", mock.Anything).Return(types.FixtureSilenced("foo:bar"), nil)

	cmd := InfoCommand(cli)
	require.NoError(t, cmd.Flags().Set("format", "tabular"))

	out, err := test.RunCmd(cmd, []string{"foo:bar"})
	require.NoError(t, err)

	assert.NotEmpty(out)
	assert.Contains(out, "Reason")
	assert.Contains(out, "Subscription")
	assert.Contains(out, "Namespace")
}

func TestInfoCommandRunEClosureWithErr(t *testing.T) {
	assert := assert.New(t)

	cli := test.NewCLI()
	client := cli.Client.(*client.MockClient)
	client.On("FetchSilenced", mock.Anything).Return(&types.Silenced{}, errors.New("my-err"))

	cmd := InfoCommand(cli)
	out, err := test.RunCmd(cmd, []string{"foo:bar"})

	assert.Equal("my-err", err.Error())
	assert.Empty(out)
}
