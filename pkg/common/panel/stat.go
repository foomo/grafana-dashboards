package panel

import (
	pkgdashboard "github.com/foomo/grafana-dashboards/pkg/common/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/stat"
	"github.com/grafana/grafana-foundation-sdk/go/units"
)

func NewStat(title string, datasource dashboard.DataSourceRef) *stat.PanelBuilder {
	return stat.NewPanelBuilder().
		Title(title).
		Datasource(datasource).
		Span(4).
		Height(3)
}

func NewTotalStat(title string, datasource dashboard.DataSourceRef) *stat.PanelBuilder {
	return NewStat(title, datasource).
		ColorScheme(pkgdashboard.NewFixedFieldColor("white")).
		NoValue("-").
		Unit(units.Number)
}
