package filter

import (
	"errors"
	"net/http"
	"testing"

	corev2 "github.com/echlebek/sensu-lite/api/core/v2"
	client "github.com/echlebek/sensu-lite/cli/client/testing"
	"github.com/echlebek/sensu-lite/cli/commands/flags"
	"github.com/echlebek/sensu-lite/cli/commands/helpers"
	test "github.com/echlebek/sensu-lite/cli/commands/testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestListCommand(t *testing.T) {
	assert := assert.New(t)

	cli := test.NewCLI()
	cmd := ListCommand(cli)

	assert.NotNil(cmd, "cmd should be returned")
	assert.NotNil(cmd.RunE, "cmd should be able to be executed")
	assert.Regexp("list", cmd.Use)
	assert.Regexp("filters", cmd.Short)
}

func TestListCommandRunEClosure(t *testing.T) {
	assert := assert.New(t)

	cli := test.NewCLI()
	client := cli.Client.(*client.MockClient)
	resources := []corev2.EventFilter{}
	client.On("List", mock.Anything, &resources, mock.Anything, mock.Anything).Return(nil).Run(
		func(args mock.Arguments) {
			resources := args[1].(*[]corev2.EventFilter)
			*resources = []corev2.EventFilter{
				*corev2.FixtureEventFilter("name-one"),
				*corev2.FixtureEventFilter("name-two"),
			}
		},
	)

	cmd := ListCommand(cli)
	require.NoError(t, cmd.Flags().Set("format", "json"))
	out, err := test.RunCmd(cmd, []string{})

	assert.NotEmpty(out)
	assert.Contains(out, "name-one")
	assert.Contains(out, "name-two")
	assert.Nil(err)
}

func TestListCommandRunEClosureWithAll(t *testing.T) {
	assert := assert.New(t)

	cli := test.NewCLI()
	client := cli.Client.(*client.MockClient)
	resources := []corev2.EventFilter{}
	client.On("List", mock.Anything, &resources, mock.Anything, mock.Anything).Return(nil).Run(
		func(args mock.Arguments) {
			resources := args[1].(*[]corev2.EventFilter)
			*resources = []corev2.EventFilter{
				*corev2.FixtureEventFilter("name-one"),
			}
		},
	)

	cmd := ListCommand(cli)
	require.NoError(t, cmd.Flags().Set(flags.Format, "json"))
	require.NoError(t, cmd.Flags().Set(flags.AllNamespaces, "t"))
	out, err := test.RunCmd(cmd, []string{})
	assert.NotEmpty(out)
	assert.Nil(err)
}

func TestListCommandRunEClosureWithTable(t *testing.T) {
	assert := assert.New(t)
	cli := test.NewCLI()

	filter := corev2.FixtureEventFilter("name-one")

	client := cli.Client.(*client.MockClient)
	resources := []corev2.EventFilter{}
	client.On("List", mock.Anything, &resources, mock.Anything, mock.Anything).Return(nil).Run(
		func(args mock.Arguments) {
			resources := args[1].(*[]corev2.EventFilter)
			*resources = []corev2.EventFilter{
				*filter,
			}
		},
	)

	cmd := ListCommand(cli)
	require.NoError(t, cmd.Flags().Set("format", "none"))
	out, err := test.RunCmd(cmd, []string{})

	assert.NotEmpty(out)
	assert.Contains(out, "Name")        // heading
	assert.Contains(out, "Action")      // heading
	assert.Contains(out, "Expressions") // heading
	assert.Nil(err)
}

// Test to ensure check command list output does not escape alphanumeric chars
func TestListCommandRunEClosureWithErr(t *testing.T) {
	assert := assert.New(t)

	cli := test.NewCLI()
	client := cli.Client.(*client.MockClient)
	resources := []corev2.EventFilter{}
	client.On("List", mock.Anything, &resources, mock.Anything, mock.Anything).Return(errors.New("my-err"))

	cmd := ListCommand(cli)
	out, err := test.RunCmd(cmd, []string{})

	assert.NotNil(err)
	assert.Equal("my-err", err.Error())
	assert.Empty(out)
}

func TestListCommandRunEClosureWithAlphaNumericChars(t *testing.T) {
	assert := assert.New(t)

	cli := test.NewCLI()
	client := cli.Client.(*client.MockClient)
	filter := corev2.FixtureEventFilter("name-one")
	filter.Expressions = append(filter.Expressions, "10 > 0")
	resources := []corev2.EventFilter{}
	client.On("List", mock.Anything, &resources, mock.Anything, mock.Anything).Return(nil).Run(
		func(args mock.Arguments) {
			resources := args[1].(*[]corev2.EventFilter)
			*resources = []corev2.EventFilter{
				*filter,
			}
		},
	)

	cmd := ListCommand(cli)
	require.NoError(t, cmd.Flags().Set(flags.Format, "json"))
	require.NoError(t, cmd.Flags().Set(flags.AllNamespaces, "t"))
	out, err := test.RunCmd(cmd, []string{})
	assert.NotEmpty(out)
	assert.Contains(out, "10 > 0")
	assert.Nil(err)
}

func TestListFlags(t *testing.T) {
	assert := assert.New(t)

	cli := test.NewCLI()
	cmd := ListCommand(cli)

	flag := cmd.Flag("all-namespaces")
	assert.NotNil(flag)

	flag = cmd.Flag("format")
	assert.NotNil(flag)
}

func TestListCommandRunEClosureWithTableAllow(t *testing.T) {
	assert := assert.New(t)
	cli := test.NewCLI()

	filter := corev2.FixtureEventFilter("name-one")
	filter.Expressions = append(filter.Expressions, "event.check.name == 'dev'")

	client := cli.Client.(*client.MockClient)
	resources := []corev2.EventFilter{}
	client.On("List", mock.Anything, &resources, mock.Anything, mock.Anything).Return(nil).Run(
		func(args mock.Arguments) {
			resources := args[1].(*[]corev2.EventFilter)
			*resources = []corev2.EventFilter{
				*filter,
			}
		},
	)

	cmd := ListCommand(cli)
	require.NoError(t, cmd.Flags().Set("format", "none"))
	out, err := test.RunCmd(cmd, []string{})

	assert.NotEmpty(out)
	assert.Contains(out, "Name")                                                       // heading
	assert.Contains(out, "Action")                                                     // heading
	assert.Contains(out, "Expressions")                                                // heading
	assert.Contains(out, "(event.check.team == 'ops') && (event.check.name == 'dev')") // allow &&
	assert.Nil(err)
}

func TestListCommandRunEClosureWithTableDeny(t *testing.T) {
	assert := assert.New(t)
	cli := test.NewCLI()

	filter := corev2.FixtureEventFilter("name-one")
	filter.Expressions = append(filter.Expressions, "event.check.name == 'dev'")
	filter.Action = corev2.EventFilterActionDeny

	client := cli.Client.(*client.MockClient)
	resources := []corev2.EventFilter{}
	client.On("List", mock.Anything, &resources, mock.Anything, mock.Anything).Return(nil).Run(
		func(args mock.Arguments) {
			resources := args[1].(*[]corev2.EventFilter)
			*resources = []corev2.EventFilter{
				*filter,
			}
		},
	)

	cmd := ListCommand(cli)
	require.NoError(t, cmd.Flags().Set("format", "none"))
	out, err := test.RunCmd(cmd, []string{})

	assert.NotEmpty(out)
	assert.Contains(out, "Name")                                                       // heading
	assert.Contains(out, "Action")                                                     // heading
	assert.Contains(out, "Expressions")                                                // heading
	assert.Contains(out, "(event.check.team == 'ops') || (event.check.name == 'dev')") // deny ||
	assert.Nil(err)
}

func TestListCommandRunEClosureWithHeader(t *testing.T) {
	assert := assert.New(t)

	cli := test.NewMockCLI()
	config := cli.Config.(*client.MockConfig)
	config.On("Format").Return("none")

	client := cli.Client.(*client.MockClient)
	var header http.Header
	resources := []corev2.EventFilter{}
	client.On("List", mock.Anything, &resources, mock.Anything, &header).Return(nil).Run(
		func(args mock.Arguments) {
			resources := args[1].(*[]corev2.EventFilter)
			*resources = []corev2.EventFilter{}
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
