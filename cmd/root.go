package cmd

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewRoot represents the base command when called without any subcommands
func NewRoot() *cobra.Command {
	c := viper.New()

	cmd := &cobra.Command{
		Use:           "grafana",
		Short:         "Grafana dashboard generator",
		SilenceErrors: true,
		SilenceUsage:  true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			pterm.PrintDebugMessages = viper.GetBool("debug")
		},
	}

	flags := cmd.PersistentFlags()

	flags.BoolP("debug", "v", false, "output debug information")
	_ = c.BindPFlag("debug", flags.Lookup("debug"))

	return cmd
}
