package dump

import (
	"testing"

	test "github.com/echlebek/sensu-lite/cli/commands/testing"
	"github.com/stretchr/testify/assert"
)

func TestCommand(t *testing.T) {
	assert := assert.New(t)

	cli := test.NewCLI()
	cmd := Command(cli)

	assert.NotNil(cmd, "cmd should be returned")
	assert.NotNil(cmd.RunE, "cmd should be able to be executed")
	assert.Regexp("dump", cmd.Use)
}

func TestCommandArgs(t *testing.T) {
	assert := assert.New(t)

	cli := test.NewCLI()
	cmd := Command(cli)

	out, err := test.RunCmd(cmd, []string{})
	assert.NotEmpty(out)
	assert.Error(err)

	// invalid resources
	out, err = test.RunCmd(cmd, []string{"check,foo"})
	assert.Empty(out)
	assert.Error(err)
}

func TestListFlags(t *testing.T) {
	assert := assert.New(t)

	cli := test.NewCLI()
	cmd := Command(cli)

	flag := cmd.Flag("all-namespaces")
	assert.NotNil(flag)

	flag = cmd.Flag("format")
	assert.NotNil(flag)

	flag = cmd.Flag("file")
	assert.NotNil(flag)
}
