package subcommands

import (
	"errors"
	"fmt"

	"github.com/echlebek/sensu-lite/cli"
	"github.com/spf13/cobra"
)

// RemoveProxyRequestsCommand adds a command that allows a user to remove the
// proxy requests for a check
func RemoveProxyRequestsCommand(cli *cli.SensuCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "remove-proxy-requests [NAME]",
		Short:        "removes proxy requests from a check",
		SilenceUsage: false,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Print usage if we do not receive one argument
			if len(args) != 1 {
				_ = cmd.Help()
				return errors.New("invalid argument(s) received")
			}

			check, err := cli.Client.FetchCheck(args[0])
			if err != nil {
				return err
			}
			check.ProxyRequests = nil

			if err := check.Validate(); err != nil {
				return err
			}
			if err := cli.Client.UpdateCheck(check); err != nil {
				return err
			}

			fmt.Fprintln(cmd.OutOrStdout(), "Removed")
			return nil
		},
	}

	return cmd
}
