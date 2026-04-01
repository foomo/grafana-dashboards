package panel

import (
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/histogram"
)

func NewRequestDurationDistribution(datasource dashboard.DataSourceRef) *histogram.PanelBuilder {
	return NewHistogram("Duration distribution", datasource).
		DisplayName("Requests").
		ColorScheme(
			dashboard.NewFieldColorBuilder().
				Mode(dashboard.FieldColorModeIdFixed).
				FixedColor("semi-dark-orange"),
		)
}
