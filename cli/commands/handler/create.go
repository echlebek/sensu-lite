package handler

import (
	"errors"
	"fmt"

	"github.com/echlebek/sensu-lite/cli"
	"github.com/echlebek/sensu-lite/cli/commands/flags"
	"github.com/echlebek/sensu-lite/cli/commands/helpers"
	"github.com/echlebek/sensu-lite/types"
	"github.com/spf13/cobra"
)

// CreateCommand adds command that allows the user to create new handlers
func CreateCommand(cli *cli.SensuCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "create [NAME]",
		Short:        "create new handlers",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 1 {
				_ = cmd.Help()
				return errors.New("invalid argument(s) received")
			}

			isInteractive, _ := cmd.Flags().GetBool(flags.Interactive)
			opts := newHandlerOpts()

			if len(args) > 0 {
				opts.Name = args[0]
			}

			opts.Namespace = cli.Config.Namespace()

			if isInteractive {
				if err := opts.administerQuestionnaire(false); err != nil {
					return err
				}
			} else {
				opts.withFlags(cmd.Flags())
			}

			handler := types.Handler{}
			opts.Copy(&handler)

			if err := handler.Validate(); err != nil {
				if !isInteractive {
					_ = cmd.Help()
					return errors.New("invalid argument(s) received")
				}
				return err
			}

			err := cli.Client.CreateHandler(&handler)
			if err != nil {
				return err
			}

			fmt.Fprintln(cmd.OutOrStdout(), "Created")
			return nil
		},
	}

	cmd.Flags().String("command", "", "command to be executed. The event data is passed to the process via STDIN")
	cmd.Flags().String("env-vars", "", "comma separated list of key=value environment variables for the mutator command")
	cmd.Flags().String("filters", "", "comma separated list of filters to use when filtering events for the handler")
	cmd.Flags().String("handlers", "", "comma separated list of handlers to call using the handler set")
	cmd.Flags().StringP("mutator", "m", "", "Sensu event mutator (name) to use to mutate event data for the handler")
	cmd.Flags().String("socket-host", "", "host of handler socket")
	cmd.Flags().String("socket-port", "", "port of handler socket")
	cmd.Flags().StringP("timeout", "i", "", "execution duration timeout in seconds (hard stop)")
	cmd.Flags().StringP("type", "t", typeDefault, "type of handler (pipe, tcp, udp, or set)")
	cmd.Flags().StringP("runtime-assets", "r", "", "comma separated list of assets this handler depends on")

	helpers.AddInteractiveFlag(cmd.Flags())
	return cmd
}
