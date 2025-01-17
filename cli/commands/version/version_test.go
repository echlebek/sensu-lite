package version

import (
	"os"
	"testing"

	test "github.com/echlebek/sensu-lite/cli/commands/testing"
	"github.com/stretchr/testify/assert"
)

func TestCommand(t *testing.T) {
	assert := assert.New(t)
	stdout := test.NewFileCapture(&os.Stdout)
	cmd := Command()
	assert.NotNil(cmd, "Returns a Command instance")
	assert.Equal("version", cmd.Use, "Configures the name")

	// Run command w/o any flags
	stdout.Start()
	cmd.Run(cmd, []string{})
	stdout.Stop()
	assert.Regexp("sensuctl version", stdout.Output())
}
