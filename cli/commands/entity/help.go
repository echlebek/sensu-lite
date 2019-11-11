package entity

import (
	"github.com/echlebek/sensu-lite/cli"
	"github.com/spf13/cobra"
)

// HelpCommand defines new parent
func HelpCommand(cli *cli.SensuCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "entity",
		Short: "Manage entities",
	}

	// Add sub-commands
	cmd.AddCommand(
		CreateCommand(cli),
		DeleteCommand(cli),
		ListCommand(cli),
		InfoCommand(cli),
		UpdateCommand(cli),
	)

	return cmd
}
