package cmd

import (
	"os"

	"github.com/alecthomas/chroma/quick"
	"github.com/foomo/grafana-dashboards/internal/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.yaml.in/yaml/v3"
)

func NewConfig() *cobra.Command {
	c := viper.New()

	cmd := &cobra.Command{
		Use:   "config",
		Short: "Print config",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.Load(c, cmd)
			if err != nil {
				return err
			}

			out, err := yaml.Marshal(cfg)
			if err != nil {
				return errors.Wrap(err, "failed to marshal config")
			}

			return quick.Highlight(os.Stdout, string(out), "yaml", "terminal", "monokai")
		},
	}

	flags := cmd.Flags()

	flags.StringSliceP("config", "c", []string{"grafana.yaml"}, "config files (default is grafana.yaml)")
	_ = c.BindPFlag("config", flags.Lookup("config"))

	return cmd
}
