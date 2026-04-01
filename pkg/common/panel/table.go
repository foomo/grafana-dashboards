package panel

import (
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/table"
)

func NewTable(title string, datasource dashboard.DataSourceRef) *table.PanelBuilder {
	return table.NewPanelBuilder().
		Datasource(datasource).
		Title(title)
}
