package panel

import (
	pkgtempo "github.com/foomo/grafana-dashboards/pkg/datasource/tempo"
	"github.com/foomo/grafana-dashboards/pkg/util"
	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
	"github.com/grafana/grafana-foundation-sdk/go/units"
)

func NewErrorRateTimeseries(datasource dashboard.DataSourceRef) *timeseries.PanelBuilder {
	return NewTimeSeries("Errors", datasource).
		Unit(units.RequestsPerSecond).
		Links([]cog.Builder[dashboard.DashboardLink]{
			dashboard.NewDashboardLinkBuilder("Traces").
				Url(pkgtempo.MustGetExploreURL(
					pkgtempo.NewNamespaceTraceQLFilter("$namespace"),
					pkgtempo.NewServiceTraceQLFilter("$service"),
					pkgtempo.NewStatusTraceQLFilter("error"),
				)).
				TargetBlank(true),
		}).
		Thresholds(
			dashboard.NewThresholdsConfigBuilder().Steps([]dashboard.Threshold{
				{
					Value: util.ToPtr(0.0),
					Color: "green",
				},
				{
					Value: util.ToPtr(0.1),
					Color: "yellow",
				},
				{
					Value: util.ToPtr(0.3),
					Color: "red",
				},
			}),
		).
		ThresholdsStyle(
			common.NewGraphThresholdsStyleConfigBuilder().Mode(
				common.GraphThresholdsStyleModeArea,
			),
		)
}
