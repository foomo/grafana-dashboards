package panel

import (
	"github.com/grafana/grafana-foundation-sdk/go/bargauge"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

func NewBarGauge(title string, datasource dashboard.DataSourceRef) *bargauge.PanelBuilder {
	return bargauge.NewPanelBuilder().
		Title(title).
		Datasource(datasource)
}
