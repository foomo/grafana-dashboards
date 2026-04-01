package node

import (
	pkgpanel "github.com/foomo/grafana-dashboards/pkg/common/panel"
	pkgprometheus "github.com/foomo/grafana-dashboards/pkg/datasource/prometheus"
	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
	"github.com/grafana/grafana-foundation-sdk/go/units"
)

func NewCPU() *timeseries.PanelBuilder {
	return pkgpanel.NewTimeSeries("CPU", pkgprometheus.Datasource()).
		Unit(units.PercentUnit).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`irate(process_cpu_time_total{namespace="$namespace", service="$service", process_cpu_state="user"}[$__interval])`,
			),
		)
}

func NewMemory() *timeseries.PanelBuilder {
	return pkgpanel.NewTimeSeries("Memory", pkgprometheus.Datasource()).
		Unit(units.BytesIEC).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`process_memory_usage{namespace="$namespace", service="$service"}`,
			).LegendFormat("resident"),
		)
}

func NewMemoryDeriv() *timeseries.PanelBuilder {
	return pkgpanel.NewTimeSeries("Memory deriv", pkgprometheus.Datasource()).
		Unit(units.BytesIEC).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`deriv(process_memory_usage{namespace="$namespace", service="$service"}[$__interval])`,
			).LegendFormat("resident"),
		)
}

func NewSystemMemory() *timeseries.PanelBuilder {
	return pkgpanel.NewTimeSeries("System memory", pkgprometheus.Datasource()).
		Legend(
			common.NewVizLegendOptionsBuilder().
				ShowLegend(true).
				DisplayMode(common.LegendDisplayModeTable).
				Calcs([]string{"mean", "max"}).
				SortBy("Max").
				SortDesc(true),
		).
		Unit(units.BytesIEC).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`system_memory_usage{namespace="$namespace", service="$service", system_memory_state="used"}`,
			).LegendFormat("used"),
		).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`system_memory_usage{namespace="$namespace", service="$service", system_memory_state="free"}`,
			).LegendFormat("free"),
		)
}

func NewSystemMemoryDeriv() *timeseries.PanelBuilder {
	return pkgpanel.NewTimeSeries("System Memory deriv", pkgprometheus.Datasource()).
		Legend(
			common.NewVizLegendOptionsBuilder().
				ShowLegend(true).
				DisplayMode(common.LegendDisplayModeTable).
				Calcs([]string{"mean", "max"}).
				SortBy("Max").
				SortDesc(true),
		).
		Unit(units.BytesIEC).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`deriv(system_memory_usage{namespace="$namespace", service="$service", system_memory_state="used"}[$__interval])`,
			).LegendFormat("used"),
		).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`deriv(system_memory_usage{namespace="$namespace", service="$service", system_memory_state="free"}[$__interval])`,
			).LegendFormat("free"),
		)
}
