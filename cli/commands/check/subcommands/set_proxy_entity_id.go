package subcommands

import (
	"errors"
	"fmt"

	"github.com/echlebek/sensu-lite/cli"
	"github.com/spf13/cobra"
)

// SetProxyEntityNameCommand updates the proxy entity name of a check
func SetProxyEntityNameCommand(cli *cli.SensuCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "set-proxy-entity-name [NAME] [VALUE]",
		Short:        "set proxy entity name of a check",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				_ = cmd.Help()
				return errors.New("invalid argument(s) received")
			}

			checkName := args[0]
			value := args[1]

			check, err := cli.Client.FetchCheck(checkName)
			if err != nil {
				return err
			}
			check.ProxyEntityName = value

			if err := check.Validate(); err != nil {
				return err
			}
			if err := cli.Client.UpdateCheck(check); err != nil {
				return err
			}

			fmt.Fprintln(cmd.OutOrStdout(), "Updated")
			return nil
		},
	}

	return cmd
}
