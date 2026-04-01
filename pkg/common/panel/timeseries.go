package panel

import (
	"github.com/foomo/grafana-dashboards/pkg/util"
	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
)

func NewTimeSeries(title string, datasource dashboard.DataSourceRef) *timeseries.PanelBuilder {
	return timeseries.NewPanelBuilder().
		Title(title).
		LineWidth(1).
		FillOpacity(62).
		Datasource(datasource).
		DrawStyle(common.GraphDrawStyleLine).
		ShowPoints(common.VisibilityModeNever).
		GradientMode(common.GraphGradientModeOpacity).
		// SpanNulls(common.BoolOrFloat64{Bool: cog.ToPtr[bool](false)}).
		AxisBorderShow(false).
		LineInterpolation(common.LineInterpolationSmooth).
		Thresholds(
			dashboard.NewThresholdsConfigBuilder().Steps([]dashboard.Threshold{
				{
					Value: util.ToPtr(0.0),
					Color: "green",
				},
			}),
		).

		// Legend(common.NewVizLegendOptionsBuilder().
		// 	DisplayMode(common.LegendDisplayModeList).
		// 	Placement(common.LegendPlacementBottom).
		// 	ShowLegend(true),
		// ).
		Tooltip(common.NewVizTooltipOptionsBuilder().
			Mode(common.TooltipDisplayModeMulti).
			Sort(common.SortOrderDescending),
		)
}
