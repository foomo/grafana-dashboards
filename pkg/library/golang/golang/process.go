package golang

import (
	pkgpanel "github.com/foomo/grafana-dashboards/pkg/common/panel"
	pkgprometheus "github.com/foomo/grafana-dashboards/pkg/datasource/prometheus"
	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
	"github.com/grafana/grafana-foundation-sdk/go/units"
)

func NewGoRoutines() *timeseries.PanelBuilder {
	return pkgpanel.NewTimeSeries("Goroutines", pkgprometheus.Datasource()).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`go_goroutines{namespace="$namespace", service="$service"}`,
			),
		)
}

func NewGoRoutinesDeriv() *timeseries.PanelBuilder {
	return pkgpanel.NewTimeSeries("Goroutines deriv", pkgprometheus.Datasource()).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`deriv(go_goroutines{namespace="$namespace", service="$service"}[$__interval])`,
			),
		)
}

func NewOpenFDS() *timeseries.PanelBuilder {
	return pkgpanel.NewTimeSeries("Open fds", pkgprometheus.Datasource()).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`process_open_fds{namespace="$namespace", service="$service"}`,
			),
		)
}

func NewOpenFDSDeriv() *timeseries.PanelBuilder {
	return pkgpanel.NewTimeSeries("Open fds deriv", pkgprometheus.Datasource()).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`deriv(process_open_fds{namespace="$namespace", service="$service"}[$__interval])`,
			),
		)
}

func NewMemStats() *timeseries.PanelBuilder {
	return pkgpanel.NewTimeSeries("Memstats", pkgprometheus.Datasource()).
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
				`go_memstats_alloc_bytes{namespace="$namespace", service="$service"}`,
			).LegendFormat("bytes allocated"),
		).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`go_memstats_stack_inuse_bytes{namespace="$namespace", service="$service"}`,
			).LegendFormat("stack inuse"),
		).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`go_memstats_heap_inuse_bytes{namespace="$namespace", service="$service"}`,
			).LegendFormat("heap inuse"),
		)
}

func NewMemStatsDeriv() *timeseries.PanelBuilder {
	return pkgpanel.NewTimeSeries("Memstats deriv", pkgprometheus.Datasource()).
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
				`deriv(go_memstats_alloc_bytes{namespace="$namespace", service="$service"}[$__interval])`,
			).LegendFormat("bytes allocated"),
		).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`rate(go_memstats_alloc_bytes_total{namespace="$namespace", service="$service"}[$__interval])`,
			).LegendFormat("bytes rate"),
		).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`deriv(go_memstats_stack_inuse_bytes{namespace="$namespace", service="$service"}[$__interval])`,
			).LegendFormat("stack inuse"),
		).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`deriv(go_memstats_heap_inuse_bytes{namespace="$namespace", service="$service"}[$__interval])`,
			).LegendFormat("heap inuse"),
		)
}

func NewMemory() *timeseries.PanelBuilder {
	return pkgpanel.NewTimeSeries("Memory", pkgprometheus.Datasource()).
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
				`process_resident_memory_bytes{namespace="$namespace", service="$service"}`,
			).LegendFormat("resident"),
		).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`process_virtual_memory_bytes{namespace="$namespace", service="$service"}`,
			).LegendFormat("virtual"),
		)
}

func NewMemoryDeriv() *timeseries.PanelBuilder {
	return pkgpanel.NewTimeSeries("Memory deriv", pkgprometheus.Datasource()).
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
				`deriv(process_resident_memory_bytes{namespace="$namespace", service="$service"}[$__interval])`,
			).LegendFormat("resident"),
		).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`deriv(process_virtual_memory_bytes{namespace="$namespace", service="$service"}[$__interval])`,
			).LegendFormat("virtual"),
		)
}

func NewResidentMemory() *timeseries.PanelBuilder {
	return pkgpanel.NewTimeSeries("Resident Memory", pkgprometheus.Datasource()).
		Unit(units.BytesIEC).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`process_resident_memory_bytes{namespace="$namespace", service="$service"}`,
			).LegendFormat("resident"),
		)
}

func NewResidentMemoryDeriv() *timeseries.PanelBuilder {
	return pkgpanel.NewTimeSeries("Resident Memory deriv", pkgprometheus.Datasource()).
		Unit(units.BytesIEC).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`deriv(process_resident_memory_bytes{namespace="$namespace", service="$service"}[$__interval])`,
			).LegendFormat("resident"),
		)
}

func NewVirtualMemory() *timeseries.PanelBuilder {
	return pkgpanel.NewTimeSeries("Virtual Memory", pkgprometheus.Datasource()).
		Unit(units.BytesIEC).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`process_virtual_memory_bytes{namespace="$namespace", service="$service"}`,
			).LegendFormat("virtual"),
		)
}

func NewVirtualMemoryDeriv() *timeseries.PanelBuilder {
	return pkgpanel.NewTimeSeries("Virtual Memory deriv", pkgprometheus.Datasource()).
		Unit(units.BytesIEC).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`deriv(process_virtual_memory_bytes{namespace="$namespace", service="$service"}[$__interval])`,
			).LegendFormat("virtual"),
		)
}

func NewGCDurationQuantiles() *timeseries.PanelBuilder {
	return pkgpanel.NewTimeSeries("GC duration quantiles", pkgprometheus.Datasource()).
		Unit(units.Milliseconds).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`go_gc_duration_seconds{namespace="$namespace", service="$service"}`,
			),
		)
}
