package panel

import (
	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
)

func NewStatusCodesTimeSeries(datasource dashboard.DataSourceRef) *timeseries.PanelBuilder {
	return NewTimeSeries("Status codes", datasource).
		MaxDataPoints(25.0).
		DrawStyle(common.GraphDrawStyleBars).
		Stacking(common.NewStackingConfigBuilder().Mode(common.StackingModeNormal)).
		GradientMode(common.GraphGradientModeNone).
		ColorScheme(
			dashboard.NewFieldColorBuilder().
				Mode(dashboard.FieldColorModeIdFixed).
				FixedColor("semi-dark-orange"),
		)
}
