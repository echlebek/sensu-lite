package rolebinding

import (
	"github.com/echlebek/sensu-lite/cli"
	"github.com/spf13/cobra"
)

// HelpCommand defines new parent
func HelpCommand(cli *cli.SensuCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "role-binding",
		Short: "Manage role bindings",
	}

	// Add sub-commands
	cmd.AddCommand(
		CreateCommand(cli),
		DeleteCommand(cli),
		ListCommand(cli),
		InfoCommand(cli),
	)

	return cmd
}
