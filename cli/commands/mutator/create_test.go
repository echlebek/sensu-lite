package mutator

import (
	"errors"
	"testing"

	client "github.com/echlebek/sensu-lite/cli/client/testing"
	test "github.com/echlebek/sensu-lite/cli/commands/testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateCommand(t *testing.T) {
	assert := assert.New(t)

	cli := test.NewMockCLI()
	cmd := CreateCommand(cli)

	assert.NotNil(cmd, "cmd should be returned")
	assert.NotNil(cmd.RunE, "cmd should be able to be executed")
	assert.Regexp("create", cmd.Use)
	assert.Regexp("mutators", cmd.Short)
}

func TestCreateCommandRunEClosureWithoutFlags(t *testing.T) {
	assert := assert.New(t)

	cli := test.NewMockCLI()
	cmd := CreateCommand(cli)
	out, err := test.RunCmd(cmd, []string{"hello"})

	assert.Empty(out)
	assert.NotNil(err)
}

func TestCreateCommandRunEClosureWithAllFlags(t *testing.T) {
	assert := assert.New(t)

	cli := test.NewMockCLI()
	client := cli.Client.(*client.MockClient)
	client.On("CreateMutator", mock.Anything).Return(nil)

	cmd := CreateCommand(cli)
	require.NoError(t, cmd.Flags().Set("command", "echo 'I like turtles'"))
	require.NoError(t, cmd.Flags().Set("timeout", "60"))
	require.NoError(t, cmd.Flags().Set("env-vars", "key1=val1,key2=val2"))
	out, err := test.RunCmd(cmd, []string{"can-holla"})

	assert.Regexp("Created", out)
	assert.Nil(err)
}

func TestCreateCommandRunEClosureWithServerErr(t *testing.T) {
	assert := assert.New(t)

	cli := test.NewMockCLI()
	client := cli.Client.(*client.MockClient)
	client.On("CreateMutator", mock.Anything).Return(errors.New("whoops"))

	cmd := CreateCommand(cli)
	require.NoError(t, cmd.Flags().Set("command", "echo 'I like turtles'"))
	out, err := test.RunCmd(cmd, []string{"can-holla"})

	assert.Empty(out)
	assert.NotNil(err)
	assert.Equal("whoops", err.Error())
}
