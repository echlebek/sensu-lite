package cmd

import (
	"github.com/echlebek/sensu-lite/version"
	"github.com/spf13/cobra"
)

func VersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show the sensu-backend version information",
		Run: func(cmd *cobra.Command, args []string) {
			version.Println("sensu-backend")
		},
	}

	return cmd
}
