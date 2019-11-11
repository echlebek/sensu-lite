package user

import (
	"github.com/echlebek/sensu-lite/cli"
	"github.com/spf13/cobra"
)

// HelpCommand defines new parent
func HelpCommand(cli *cli.SensuCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "Manage users",
	}

	// Add sub-commands
	cmd.AddCommand(
		AddGroupCommand(cli),
		CreateCommand(cli),
		DeleteCommand(cli),
		ListCommand(cli),
		ReinstateCommand(cli),
		RemoveGroupCommand(cli),
		RemoveAllGroupsCommand(cli),
		SetGroupsCommand(cli),
		SetPasswordCommand(cli),
		TestCredsCommand(cli),
	)

	return cmd
}
