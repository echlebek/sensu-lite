package tessen

import (
	"github.com/echlebek/sensu-lite/cli"
	"github.com/spf13/cobra"
)

// HelpCommand defines new parent
func HelpCommand(cli *cli.SensuCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tessen",
		Short: "Manage tessen configuration",
	}

	// Add sub-commands
	cmd.AddCommand(
		OptInCommand(cli),
		OptOutCommand(cli),
		InfoCommand(cli),
	)

	return cmd
}
