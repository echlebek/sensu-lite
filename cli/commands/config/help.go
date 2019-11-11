package config

import (
	"github.com/echlebek/sensu-lite/cli"
	"github.com/spf13/cobra"
)

// HelpCommand defines new parent
func HelpCommand(cli *cli.SensuCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Modify sensuctl configuration",
	}

	// Add sub-commands
	cmd.AddCommand(
		SetFormatCommand(cli),
		SetNamespaceCommand(cli),
		ViewCommand(cli),
	)

	return cmd
}
