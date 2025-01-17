package agent

import (
	"context"
	"testing"

	"github.com/echlebek/sensu-lite/command"
	"github.com/echlebek/sensu-lite/testing/mockexecutor"

	"github.com/echlebek/sensu-lite/transport"
	"github.com/echlebek/sensu-lite/types"
	"github.com/stretchr/testify/assert"
)

func TestExecuteHook(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	hookConfig := types.FixtureHookConfig("hook")
	hookConfig.Stdin = true

	config, cleanup := FixtureConfig()
	defer cleanup()
	agent, err := NewAgent(config)
	if err != nil {
		t.Fatal(err)
	}
	ch := make(chan *transport.Message, 1)
	agent.sendq = ch
	ex := &mockexecutor.MockExecutor{}
	agent.executor = ex
	execution := command.FixtureExecutionResponse(0, "")
	ex.Return(execution, nil)

	evt := &types.Event{
		Check: &types.Check{
			ObjectMeta: types.ObjectMeta{
				Name: "check",
			},
		},
	}

	hook := agent.executeHook(ctx, hookConfig, evt, nil)

	assert.NotZero(hook.Executed)
	assert.Equal(int32(0), hook.Status)
	assert.Equal("", hook.Output)

	execution.Output = "hello"
	hook = agent.executeHook(ctx, hookConfig, evt, nil)

	assert.NotZero(hook.Executed)
	assert.Equal(int32(0), hook.Status)
	assert.Equal("hello", hook.Output)
}

func TestPrepareHook(t *testing.T) {
	assert := assert.New(t)

	config, cleanup := FixtureConfig()
	defer cleanup()
	agent, err := NewAgent(config)
	if err != nil {
		t.Fatal(err)
	}

	// nil hook
	assert.False(agent.prepareHook(nil))

	// Invalid hook
	hook := types.FixtureHookConfig("hook")
	hook.Command = ""
	assert.False(agent.prepareHook(hook))

	// Valid check
	hook.Command = "{{ .name }}"
	assert.True(agent.prepareHook(hook))
}

func TestHookInList(t *testing.T) {
	assert := assert.New(t)
	hook1 := types.FixtureHook("hook1")
	hook2 := types.FixtureHook("hook2")

	testCases := []struct {
		name     string
		hookName string
		hookList []*types.Hook
		expected bool
	}{
		{
			name:     "Empty list",
			hookName: "hook1",
			hookList: []*types.Hook{},
			expected: false,
		},
		{
			name:     "Hook in populated list",
			hookName: "hook1",
			hookList: []*types.Hook{hook2, hook1},
			expected: true,
		},
		{
			name:     "Hook not in populated list",
			hookName: "hook1",
			hookList: []*types.Hook{hook2, hook2},
			expected: false,
		},
		{
			name:     "No hook name provided",
			hookName: "",
			hookList: []*types.Hook{hook1, hook2},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			in := hookInList(tc.hookName, tc.hookList)
			assert.Equal(tc.expected, in)
		})
	}
}
