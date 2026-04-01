package panel

import (
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
	"github.com/grafana/grafana-foundation-sdk/go/units"
)

func NewRequestRateTimeseries(datasource dashboard.DataSourceRef) *timeseries.PanelBuilder {
	return NewTimeSeries("Requests", datasource).
		Unit(units.RequestsPerSecond)
}
