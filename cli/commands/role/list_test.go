package role

import (
	"errors"
	"net/http"
	"testing"

	corev2 "github.com/echlebek/sensu-lite/api/core/v2"
	client "github.com/echlebek/sensu-lite/cli/client/testing"
	"github.com/echlebek/sensu-lite/cli/commands/helpers"
	test "github.com/echlebek/sensu-lite/cli/commands/testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestListCommand(t *testing.T) {
	assert := assert.New(t)
	cli := test.NewCLI()
	cmd := ListCommand(cli)
	assert.NotNil(cmd, "cmd should be returned")
	assert.NotNil(cmd.RunE, "cmd should be able to be executed")
	assert.Regexp("list", cmd.Use)
	assert.Regexp("roles", cmd.Short)
}
func TestListCommandRunEClosureJSONFormat(t *testing.T) {
	assert := assert.New(t)
	cli := test.NewCLI()
	client := cli.Client.(*client.MockClient)
	resources := []corev2.Role{}
	client.On("List", mock.Anything, &resources, mock.Anything, mock.Anything).Return(nil).Run(
		func(args mock.Arguments) {
			resources := args[1].(*[]corev2.Role)
			*resources = []corev2.Role{
				*corev2.FixtureRole("one", "default"),
				*corev2.FixtureRole("two", "default"),
			}
		},
	)
	cmd := ListCommand(cli)
	out, err := test.RunCmd(cmd, []string{})
	assert.NotEmpty(out)
	assert.Nil(err)
	assert.NotContains(out, "==")
}
func TestListCommandRunEClosureTabularFormat(t *testing.T) {
	assert := assert.New(t)
	cli := test.NewMockCLI()
	config := cli.Config.(*client.MockConfig)
	config.On("Format").Return("")
	client := cli.Client.(*client.MockClient)
	resources := []corev2.Role{}
	client.On("List", mock.Anything, &resources, mock.Anything, mock.Anything).Return(nil).Run(
		func(args mock.Arguments) {
			resources := args[1].(*[]corev2.Role)
			*resources = []corev2.Role{
				*corev2.FixtureRole("one", "default"),
				*corev2.FixtureRole("two", "default"),
			}
		},
	)
	cmd := ListCommand(cli)
	out, err := test.RunCmd(cmd, []string{})
	assert.NotEmpty(out)
	assert.Contains(out, "Name")
	assert.Contains(out, "one")
	assert.Contains(out, "two")
	assert.Nil(err)
}
func TestListCommandRunEClosureWithErr(t *testing.T) {
	assert := assert.New(t)
	cli := test.NewCLI()
	client := cli.Client.(*client.MockClient)
	resources := []corev2.Role{}
	client.On("List", mock.Anything, &resources, mock.Anything, mock.Anything).Return(errors.New("fire"))
	cmd := ListCommand(cli)
	out, err := test.RunCmd(cmd, []string{})
	assert.Empty(out)
	assert.NotNil(err)
	assert.Equal("fire", err.Error())
}

func TestListCommandRunEClosureWithHeader(t *testing.T) {
	assert := assert.New(t)

	cli := test.NewMockCLI()
	config := cli.Config.(*client.MockConfig)
	config.On("Format").Return("none")

	client := cli.Client.(*client.MockClient)
	var header http.Header
	resources := []corev2.Role{}
	client.On("List", mock.Anything, &resources, mock.Anything, &header).Return(nil).Run(
		func(args mock.Arguments) {
			resources := args[1].(*[]corev2.Role)
			*resources = []corev2.Role{}
			header := args[3].(*http.Header)
			*header = make(http.Header)
			header.Add(helpers.HeaderWarning, "E_TOO_MANY_ENTITIES")
		},
	)

	cmd := ListCommand(cli)
	out, err := test.RunCmd(cmd, []string{})

	assert.NotEmpty(out)
	assert.Nil(err)
	assert.Contains(out, "E_TOO_MANY_ENTITIES")
	assert.Contains(out, "==")
}
