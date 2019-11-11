package commands

import (
	"github.com/echlebek/sensu-lite/cli"
	"github.com/echlebek/sensu-lite/cli/commands/apikey"
	"github.com/echlebek/sensu-lite/cli/commands/asset"
	"github.com/echlebek/sensu-lite/cli/commands/check"
	"github.com/echlebek/sensu-lite/cli/commands/cluster"
	"github.com/echlebek/sensu-lite/cli/commands/clusterrole"
	"github.com/echlebek/sensu-lite/cli/commands/clusterrolebinding"
	"github.com/echlebek/sensu-lite/cli/commands/completion"
	"github.com/echlebek/sensu-lite/cli/commands/config"
	"github.com/echlebek/sensu-lite/cli/commands/configure"
	"github.com/echlebek/sensu-lite/cli/commands/create"
	"github.com/echlebek/sensu-lite/cli/commands/delete"
	"github.com/echlebek/sensu-lite/cli/commands/dump"
	"github.com/echlebek/sensu-lite/cli/commands/edit"
	"github.com/echlebek/sensu-lite/cli/commands/entity"
	"github.com/echlebek/sensu-lite/cli/commands/env"
	"github.com/echlebek/sensu-lite/cli/commands/event"
	"github.com/echlebek/sensu-lite/cli/commands/filter"
	"github.com/echlebek/sensu-lite/cli/commands/handler"
	"github.com/echlebek/sensu-lite/cli/commands/hook"
	"github.com/echlebek/sensu-lite/cli/commands/logout"
	"github.com/echlebek/sensu-lite/cli/commands/mutator"
	"github.com/echlebek/sensu-lite/cli/commands/namespace"
	"github.com/echlebek/sensu-lite/cli/commands/role"
	"github.com/echlebek/sensu-lite/cli/commands/rolebinding"
	"github.com/echlebek/sensu-lite/cli/commands/silenced"
	"github.com/echlebek/sensu-lite/cli/commands/tessen"
	"github.com/echlebek/sensu-lite/cli/commands/user"
	"github.com/spf13/cobra"
)

// AddCommands adds management commands to given command
func AddCommands(rootCmd *cobra.Command, cli *cli.SensuCli) {
	rootCmd.AddCommand(
		configure.Command(cli),
		completion.Command(rootCmd),
		env.Command(cli),
		logout.Command(cli),

		// Management Commands
		asset.HelpCommand(cli),
		apikey.HelpCommand(cli),
		check.HelpCommand(cli),
		config.HelpCommand(cli),
		clusterrole.HelpCommand(cli),
		clusterrolebinding.HelpCommand(cli),
		entity.HelpCommand(cli),
		event.HelpCommand(cli),
		filter.HelpCommand(cli),
		handler.HelpCommand(cli),
		hook.HelpCommand(cli),
		mutator.HelpCommand(cli),
		namespace.HelpCommand(cli),
		role.HelpCommand(cli),
		rolebinding.HelpCommand(cli),
		user.HelpCommand(cli),
		silenced.HelpCommand(cli),
		create.CreateCommand(cli),
		delete.DeleteCommand(cli),
		//extension.HelpCommand(cli),
		cluster.HelpCommand(cli),
		edit.Command(cli),
		tessen.HelpCommand(cli),
		dump.Command(cli),
	)

	for _, cmd := range rootCmd.Commands() {
		rootCmd.ValidArgs = append(rootCmd.ValidArgs, cmd.Use)
	}
}
