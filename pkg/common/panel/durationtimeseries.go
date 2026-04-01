package panel

import (
	"time"

	pkgtempo "github.com/foomo/grafana-dashboards/pkg/datasource/tempo"
	"github.com/foomo/grafana-dashboards/pkg/util"
	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
	"github.com/grafana/grafana-foundation-sdk/go/units"
)

func NewDurationTimeSeries(datasource dashboard.DataSourceRef) *timeseries.PanelBuilder {
	return NewTimeSeries("Duration P$percentile", datasource).
		Links([]cog.Builder[dashboard.DashboardLink]{
			dashboard.NewDashboardLinkBuilder("Traces").
				Url(pkgtempo.MustGetExploreURL(
					pkgtempo.NewNamespaceTraceQLFilter("$namespace"),
					pkgtempo.NewServiceTraceQLFilter("$service"),
					pkgtempo.NewMinDurationTraceQLFilter(500*time.Millisecond),
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
					Value: util.ToPtr(0.2),
					Color: "yellow",
				},
				{
					Value: util.ToPtr(0.5),
					Color: "red",
				},
			}),
		).
		ThresholdsStyle(
			common.NewGraphThresholdsStyleConfigBuilder().Mode(
				common.GraphThresholdsStyleModeArea,
			),
		).
		Unit(units.Seconds)
}
