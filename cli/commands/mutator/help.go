package mutator

import (
	"github.com/echlebek/sensu-lite/cli"
	"github.com/spf13/cobra"
)

// HelpCommand defines new parent
func HelpCommand(cli *cli.SensuCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mutator",
		Short: "Manage mutators",
	}

	// Add sub-commands
	cmd.AddCommand(
		CreateCommand(cli),
		DeleteCommand(cli),
		InfoCommand(cli),
		ListCommand(cli),
		UpdateCommand(cli),
	)

	return cmd
}
