package panel

import (
	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
	"github.com/grafana/grafana-foundation-sdk/go/units"
)

func NewNetworkTimeseries(datasource dashboard.DataSourceRef) *timeseries.PanelBuilder {
	return NewTimeSeries("Network", datasource).
		Legend(common.NewVizLegendOptionsBuilder().ShowLegend(false)).
		AxisCenteredZero(true).
		Unit(units.BytesIEC)
}
