package tempo

import (
	pkgdashboard "github.com/foomo/grafana-dashboards/pkg/common/dashboard"
	pkgpanel "github.com/foomo/grafana-dashboards/pkg/common/panel"
	pkgprometheus "github.com/foomo/grafana-dashboards/pkg/datasource/prometheus"
	"github.com/grafana/grafana-foundation-sdk/go/stat"
)

func NewTotalSpans() *stat.PanelBuilder {
	return pkgpanel.NewTotalStat("Total Spans", pkgprometheus.Datasource()).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`round(sum(increase(traces_spanmetrics_calls_total{span_kind=~"SPAN_KIND_SERVER|SPAN_KIND_CONSUMER", k8s_namespace_name="$namespace", service="$service"}[$__range])))`,
			).Instant())
}

func NewTotalErrorSpans() *stat.PanelBuilder {
	return pkgpanel.NewTotalStat("Total Errors", pkgprometheus.Datasource()).
		ColorScheme(pkgdashboard.NewFixedFieldColor("red")).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`round(sum(increase(traces_spanmetrics_calls_total{span_kind=~"SPAN_KIND_SERVER|SPAN_KIND_CONSUMER", k8s_namespace_name="$namespace", service="$service", status_code="STATUS_CODE_ERROR"}[$__range])))`,
			).Instant(),
		)
}
