package panel

import (
	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/histogram"
)

func NewHistogram(title string, datasource dashboard.DataSourceRef) *histogram.PanelBuilder {
	return histogram.NewPanelBuilder().
		Title(title).
		Datasource(datasource).
		Legend(common.NewVizLegendOptionsBuilder().ShowLegend(false))
}
