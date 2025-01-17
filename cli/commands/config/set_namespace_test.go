package config

import (
	"errors"
	"testing"

	"github.com/echlebek/sensu-lite/cli"
	clienttest "github.com/echlebek/sensu-lite/cli/client/testing"
	test "github.com/echlebek/sensu-lite/cli/commands/testing"
	"github.com/stretchr/testify/assert"
)

func TestSetNamespaceCommand(t *testing.T) {
	assert := assert.New(t)

	cli := &cli.SensuCli{}
	cmd := SetNamespaceCommand(cli)

	assert.NotNil(cmd, "cmd should be returned")
	assert.NotNil(cmd.RunE, "cmd should be able to be executed")
	assert.Regexp("set-namespace", cmd.Use)
	assert.Regexp("Set namespace", cmd.Short)
}

func TestSetNamespaceBadsArgs(t *testing.T) {
	assert := assert.New(t)

	cli := &cli.SensuCli{}
	cmd := SetNamespaceCommand(cli)

	// No args...
	out, err := test.RunCmd(cmd, []string{})
	assert.NotEmpty(out, "output should display help usage")
	assert.Error(err, "error should be returned")

	// Too many args...
	out, err = test.RunCmd(cmd, []string{"one", "two"})
	assert.NotEmpty(out, "output should display help usage")
	assert.Error(err, "error should be returned")
}

func TestSetNamespaceExec(t *testing.T) {
	assert := assert.New(t)

	cli := test.NewMockCLI()
	cmd := SetNamespaceCommand(cli)

	config := cli.Config.(*clienttest.MockConfig)
	config.On("SaveNamespace", "default").Return(nil)

	out, err := test.RunCmd(cmd, []string{"default"})
	assert.Equal(out, "Updated\n")
	assert.Nil(err, "Should not produce any errors")
}

func TestSetNamespaceWithWriteErr(t *testing.T) {
	assert := assert.New(t)

	cli := test.NewMockCLI()
	cmd := SetNamespaceCommand(cli)

	config := cli.Config.(*clienttest.MockConfig)
	config.On("SaveNamespace", "default").Return(errors.New("blah"))

	out, err := test.RunCmd(cmd, []string{"default"})
	assert.Contains(out, "Unable to write")
	assert.Nil(err, "Should not return an error")
}
