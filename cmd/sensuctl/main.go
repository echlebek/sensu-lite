//+build !windows

// Use a windows-only main file in order to get an .exe when cross compiling.
package main

import (
	"os"

	"github.com/echlebek/sensu-lite/cli"
	"github.com/echlebek/sensu-lite/cli/commands"
	hooks "github.com/echlebek/sensu-lite/cli/commands/hooks"
	"github.com/echlebek/sensu-lite/cli/commands/root"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := root.Command()
	sensuCli := cli.New(rootCmd.PersistentFlags())

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, _ []string) error {
		return hooks.ConfigurationPresent(cmd, sensuCli)
	}

	commands.AddCommands(rootCmd, sensuCli)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
