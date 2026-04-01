package http

import (
	pkgdashboard "github.com/foomo/grafana-dashboards/pkg/common/dashboard"
	pkgpanel "github.com/foomo/grafana-dashboards/pkg/common/panel"
	pkgvariable "github.com/foomo/grafana-dashboards/pkg/common/variable"
	pkgprometheus "github.com/foomo/grafana-dashboards/pkg/datasource/prometheus"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/histogram"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"
	"github.com/grafana/grafana-foundation-sdk/go/stat"
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
)

func NewClientDashboard(namespace, service string) *dashboard.DashboardBuilder {
	return pkgdashboard.NewDashboard(
		"opentelemetry_http_client",
		"HTTP Client",
		"OpenTelemetry HTTP Client",
	).
		Variables(pkgdashboard.Variables{
			pkgvariable.Percentile(),
			pkgvariable.IntervalVariable(),
			pkgvariable.Constant("namespace", namespace),
			pkgvariable.Constant("service", service),
			pkgprometheus.NewAdHocVariable(),
		}).
		WithPanel(
			NewClientRequestRate().Span(8),
		).
		WithPanel(
			NewClientErrorRate().Span(8),
		).
		WithPanel(
			NewClientDuration().Span(8),
		).
		WithPanel(
			NewClientRequestDurationDistribution().Span(12),
		).
		WithPanel(
			NewClientStatusCodes().Span(12),
		).
		WithPanel(
			NewClientNetwork().Span(24),
		)
}

func NewClientRequestRate() *timeseries.PanelBuilder {
	return pkgpanel.NewRequestRateTimeseries(pkgprometheus.Datasource()).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`sum(rate(http_client_request_duration_seconds_count{namespace="$namespace", service="$service"}[$__interval]))`,
			),
		)
}

func NewClientTotalRequests() *stat.PanelBuilder {
	return pkgpanel.NewTotalStat("Total Requests", pkgprometheus.Datasource()).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`round(sum(increase(http_client_request_duration_seconds_count{namespace="$namespace", service="$service"}[$__range])))`,
			).Instant(),
		)
}

func NewClientErrorRate() *timeseries.PanelBuilder {
	return pkgpanel.NewErrorRateTimeseries(pkgprometheus.Datasource()).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`(
          sum(rate(http_client_request_duration_seconds_count{namespace="$namespace", service="$service", http_status_code=~"(4|5).*"}[$__interval]))
          /
          sum(rate(http_client_request_duration_seconds_count{namespace="$namespace", service="$service"}[$__interval]))
          )`,
			),
		)
}

func NewClientDuration() *timeseries.PanelBuilder {
	return pkgpanel.NewDurationTimeSeries(pkgprometheus.Datasource()).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`histogram_quantile($percentile, sum(rate(http_client_request_duration_seconds_bucket{namespace="$namespace", service="$service"}[$__interval])) by (le))`,
			).Exemplar(true),
		)
}

func NewClientStatusCodes() *timeseries.PanelBuilder {
	return pkgpanel.NewStatusCodesTimeSeries(pkgprometheus.Datasource()).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`round(sum(increase(http_client_request_duration_seconds_count{namespace="$namespace", service="$service"}[$__interval])) by (http_response_status_code))`,
			).LegendFormat("{{http_response_status_code}}"),
		)
}

func NewClientRequestDurationDistribution() *histogram.PanelBuilder {
	return pkgpanel.NewRequestDurationDistribution(pkgprometheus.Datasource()).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`sum(increase(http_client_request_duration_seconds_bucket{namespace="$namespace", service="$service"}[$__range])) by (le)`,
			).
				LegendFormat("{{le}}").
				Format(prometheus.PromQueryFormatHeatmap).
				Instant(),
		).
		Mappings([]dashboard.ValueMapping{NewBucketValueMapping()})
}

func NewClientNetwork() *timeseries.PanelBuilder {
	return pkgpanel.NewNetworkTimeseries(pkgprometheus.Datasource()).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`sum(increase(http_client_request_body_size_bytes_sum{namespace="$namespace", service="$service"}[$__interval]))`,
			).LegendFormat("Request"),
		).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`- sum(increase(http_client_response_body_size_bytes_sum{namespace="$namespace", service="$service"}[$__interval]))`,
			).LegendFormat("Response"),
		)
}
