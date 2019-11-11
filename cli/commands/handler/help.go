package handler

import (
	"github.com/echlebek/sensu-lite/cli"
	"github.com/spf13/cobra"
)

// HelpCommand defines new parent
func HelpCommand(cli *cli.SensuCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "handler",
		Short: "Manage handlers",
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
