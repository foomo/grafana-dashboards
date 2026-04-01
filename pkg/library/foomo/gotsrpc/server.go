package gotsrpc

import (
	pkgdashboard "github.com/foomo/grafana-dashboards/pkg/common/dashboard"
	pkgpanel "github.com/foomo/grafana-dashboards/pkg/common/panel"
	pkgvariable "github.com/foomo/grafana-dashboards/pkg/common/variable"
	pkgprometheus "github.com/foomo/grafana-dashboards/pkg/datasource/prometheus"
	"github.com/foomo/grafana-dashboards/pkg/util"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"
	"github.com/grafana/grafana-foundation-sdk/go/stat"
	"github.com/grafana/grafana-foundation-sdk/go/table"
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
	"github.com/grafana/grafana-foundation-sdk/go/units"
)

func NewServerDashboard(namespace, service string) *dashboard.DashboardBuilder {
	return pkgdashboard.NewDashboard(
		"foomo_gotsrpc_server",
		"GoTSRPC Server",
		"Foomo GoTSRPC Server",
	).
		Variables(pkgdashboard.Variables{
			pkgvariable.Percentile(),
			pkgvariable.IntervalVariable(),
			pkgvariable.Constant("namespace", namespace),
			pkgvariable.Constant("service", service),
			pkgprometheus.NewAdHocVariable(),
		}).
		WithPanel(
			NewServerRequestRate().Span(8),
		).
		WithPanel(
			NewServerErrorRate().Span(8),
		).
		WithPanel(
			NewServerDurationPanel().Span(8),
		).
		WithPanel(
			NewServerTable().Span(24),
		)
}

func NewServerRequestRate() *timeseries.PanelBuilder {
	return pkgpanel.NewRequestRateTimeseries(pkgprometheus.Datasource()).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`sum(rate(gotsrpc_execution_duration_seconds_count{namespace="$namespace", service="$service"}[$__interval]))`,
			),
		)
}

func NewServerTotalRequests() *stat.PanelBuilder {
	return pkgpanel.NewTotalStat("Total Requests", pkgprometheus.Datasource()).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`round(sum(increase(gotsrpc_execution_duration_seconds_count{namespace="$namespace", service="$service"}[$__range])))`,
			).Instant(),
		)
}

func NewServerErrorRate() *timeseries.PanelBuilder {
	return pkgpanel.NewErrorRateTimeseries(pkgprometheus.Datasource()).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`(
          sum(rate(gotsrpc_execution_duration_seconds_count{namespace="$namespace", service="$service", gotsrpc_error="true"}[$__interval]))
          /
          sum(rate(gotsrpc_execution_duration_seconds_count{namespace="$namespace", service="$service"}[$__interval]))
          )`,
			),
		)
}

func NewServerDurationPanel() *timeseries.PanelBuilder {
	return pkgpanel.NewDurationTimeSeries(pkgprometheus.Datasource()).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`histogram_quantile($percentile, sum(rate(gotsrpc_execution_duration_seconds_bucket{namespace="$namespace", service="$service"}[$__interval])) by (le))`,
			).Exemplar(true),
		)
}

func NewServerTable() *table.PanelBuilder {
	return pkgpanel.NewTable("Stats", pkgprometheus.Datasource()).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`round(sum(increase(gotsrpc_execution_duration_seconds_count{namespace="$namespace", service="$service"}[$__range])) by (gotsrpc_package, gotsrpc_service, gotsrpc_func))`,
			).Format(prometheus.PromQueryFormatTable).Instant(),
		).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`
					(
          sum(rate(gotsrpc_execution_duration_seconds_count{namespace="$namespace", service="$service", gotsrpc_error="true"}[$__range])) by (gotsrpc_package, gotsrpc_service, gotsrpc_func)
          /
          sum(rate(gotsrpc_execution_duration_seconds_count{namespace="$namespace", service="$service"}[$__interval])) by (gotsrpc_package, gotsrpc_service, gotsrpc_func)
          )
				`,
			).Format(prometheus.PromQueryFormatTable).Instant(),
		).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`
					histogram_quantile($percentile, sum(rate(gotsrpc_execution_duration_seconds_bucket{namespace="$namespace", service="$service"}[$__range])) by (gotsrpc_package, gotsrpc_service, gotsrpc_func, le))
				`,
			).Format(prometheus.PromQueryFormatTable).Instant(),
		).
		WithTransformation(dashboard.DataTransformerConfig{
			Id: "merge",
		}).
		WithTransformation(dashboard.DataTransformerConfig{
			Id: "organize",
			Options: util.MustYamlToMap(`
        excludeByName:
          Time: true
				indexByName:
				  Time: 0
          gotsrpc_package: 1
          gotsrpc_service: 2
          gotsrpc_func: 3
          "Value #A": 4
          "Value #B": 5
          "Value #C": 6
				renameByName:
          Time: ""
          "Value #A": Calls
          "Value #B": Errors
          "Value #C": Duration
          gotsrpc_func: Method
          gotsrpc_package: Package
          gotsrpc_service: Service
		`)}).
		OverrideByName("Calls", []dashboard.DynamicConfigValue{
			{
				Id:    "unit",
				Value: units.Number,
			},
			{
				Id:    "custom.width",
				Value: 100,
			},
		}).
		OverrideByName("Errors", []dashboard.DynamicConfigValue{
			{
				Id:    "unit",
				Value: units.PercentUnit,
			},
			{
				Id:    "custom.width",
				Value: 100,
			},
		}).
		OverrideByName("Duration", []dashboard.DynamicConfigValue{
			{
				Id:    "unit",
				Value: units.Seconds,
			},
			{
				Id:    "custom.width",
				Value: 100,
			},
		}).
		NoValue("-")
}
