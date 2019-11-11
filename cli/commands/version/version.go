package version

import (
	"github.com/echlebek/sensu-lite/cli/commands/hooks"
	"github.com/echlebek/sensu-lite/version"
	"github.com/spf13/cobra"
)

// Command defines the version command
func Command() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show the sensuctl version information",
		Run: func(cmd *cobra.Command, args []string) {
			version.Println("sensuctl")
		},
		Annotations: map[string]string{
			hooks.ConfigurationRequirement: hooks.ConfigurationNotRequired,
		},
	}
}
