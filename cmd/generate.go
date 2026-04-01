package cmd

import (
	"github.com/foomo/grafana-dashboards/internal/config"
	"github.com/foomo/grafana-dashboards/pkg/api"
	manifest2 "github.com/foomo/grafana-dashboards/pkg/common/manifest"
	foomokeellib "github.com/foomo/grafana-dashboards/pkg/library/foomo/keel"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewGenerate() *cobra.Command {
	c := viper.New()

	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate resources",
		RunE: func(cmd *cobra.Command, args []string) error {

			cfg, err := config.Load(c, cmd)
			if err != nil {
				return err
			}

			client := api.NewClient(cfg.Grafana)

			if err := api.CreateFolders(cmd.Context(), client, []api.Folder{
				{
					UID:  "sesamy",
					Name: "sesamy",
					Folders: []api.Folder{
						{
							UID:  "sesamy_site",
							Name: "Site",
							Folders: []api.Folder{
								{
									UID:     "sesamy_site_backend",
									Name:    "Backend",
									Folders: nil,
								},
							},
						},
					},
				},
			}, ""); err != nil {
				return err
			}

			tags := []string{"sesamy", "sesamy-site", "sesamy-site-backend"}
			builder := foomokeellib.NewServerDashboard("squadron-sesamy-site", "site-sesamy-gtm-tagging").Tags(tags)
			// builder := gotsrpc.NewServerDashboard("squadron-sesamy-site", "site-backend").Tags(tags)
			// builder := squadron.NewReleasesDashboard().Tags(tags)
			// builder := otelhttp.NewClientDashboard("squadron-sesamy-site", "site-backend").Tags(tags)

			d, err := builder.Build()
			if err != nil {
				return err
			}

			// fmt.Println(string(out))
			return manifest2.Generate(manifest2.Dashboard(d, "sesamy_site_backend"), c.GetBool("raw"))
			// return quick.Highlight(os.Stdout, string(out), "yaml", "terminal", "monokai")
		},
	}

	flags := cmd.Flags()

	flags.StringSliceP("config", "c", []string{"grafana.yaml"}, "config files (default is grafana.yaml)")
	_ = c.BindPFlag("config", flags.Lookup("config"))

	flags.BoolP("raw", "r", false, "Print unformatted output")
	_ = c.BindPFlag("raw", flags.Lookup("raw"))

	return cmd
}
