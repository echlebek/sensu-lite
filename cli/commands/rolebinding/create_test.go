package rolebinding

import (
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
	assert.Regexp("role binding", cmd.Short)
}

func TestCreateCommandRunEClosureMissingArgs(t *testing.T) {
	assert := assert.New(t)
	cli := test.NewMockCLI()

	cmd := CreateCommand(cli)
	out, err := test.RunCmd(cmd, []string{})

	// Print help usage
	assert.NotEmpty(out)
	assert.Error(err)
}

func TestCreateCommandRoleRef(t *testing.T) {
	assert := assert.New(t)
	cli := test.NewMockCLI()
	client := cli.Client.(*client.MockClient)
	client.On("CreateRoleBinding", mock.Anything).Return(nil)

	// No role or ClusterRole provided
	cmd := CreateCommand(cli)
	require.NoError(t, cmd.Flags().Set("user", "foo"))
	_, err := test.RunCmd(cmd, []string{"admin"})
	assert.Error(err)

	// Role provided
	cmd = CreateCommand(cli)
	require.NoError(t, cmd.Flags().Set("role", "admin"))
	require.NoError(t, cmd.Flags().Set("user", "foo"))
	out, err := test.RunCmd(cmd, []string{"admin"})
	assert.NoError(err)
	assert.Regexp("Created", out)

	// Role and ClusterRole both provided
	cmd = CreateCommand(cli)
	require.NoError(t, cmd.Flags().Set("role", "admin"))
	require.NoError(t, cmd.Flags().Set("cluster-role", "admin"))
	_, err = test.RunCmd(cmd, []string{"admin"})
	assert.Error(err)
}

func TestCreateCommandSubjects(t *testing.T) {
	assert := assert.New(t)
	cli := test.NewMockCLI()
	client := cli.Client.(*client.MockClient)
	client.On("CreateRoleBinding", mock.Anything).Return(nil)

	// No user or group provided
	cmd := CreateCommand(cli)
	require.NoError(t, cmd.Flags().Set("role", "admin"))
	_, err := test.RunCmd(cmd, []string{"admin"})
	assert.Error(err)

	// A user was provided
	cmd = CreateCommand(cli)
	require.NoError(t, cmd.Flags().Set("role", "admin"))
	require.NoError(t, cmd.Flags().Set("user", "foo"))
	out, err := test.RunCmd(cmd, []string{"admin"})
	assert.NoError(err)
	assert.Regexp("Created", out)

	// A group was provided
	cmd = CreateCommand(cli)
	require.NoError(t, cmd.Flags().Set("role", "admin"))
	require.NoError(t, cmd.Flags().Set("group", "bar"))
	_, err = test.RunCmd(cmd, []string{"admin"})
	assert.NoError(err)
}
