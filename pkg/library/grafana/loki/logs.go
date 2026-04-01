package loki

import (
	pkgdashboard "github.com/foomo/grafana-dashboards/pkg/common/dashboard"
	pkgpanel "github.com/foomo/grafana-dashboards/pkg/common/panel"
	pkgloki "github.com/foomo/grafana-dashboards/pkg/datasource/loki"
	"github.com/grafana/grafana-foundation-sdk/go/stat"
)

func NewTotalLogs() *stat.PanelBuilder {
	return pkgpanel.NewTotalStat("Total Logs", pkgloki.Datasource()).
		WithTarget(
			pkgloki.NewDataQuery(
				`sum(count_over_time({namespace="$namespace", service_name="$service"} [$__auto]))`,
			).Instant(true),
		)
}

func NewTotalWarnings() *stat.PanelBuilder {
	return pkgpanel.NewTotalStat("Total Warnings", pkgloki.Datasource()).
		ColorScheme(pkgdashboard.NewFixedFieldColor("semi-dark-orange")).
		WithTarget(
			pkgloki.NewDataQuery(
				`sum(count_over_time({namespace="$namespace", service_name="$service"} | detected_level="warning" [$__auto]))`,
			).Instant(true),
		)
}

func NewTotalErrors() *stat.PanelBuilder {
	return pkgpanel.NewTotalStat("Total Errors", pkgloki.Datasource()).
		ColorScheme(pkgdashboard.NewFixedFieldColor("red")).
		WithTarget(
			pkgloki.NewDataQuery(
				`sum(count_over_time({namespace="$namespace", service_name="$service"} | detected_level="error" [$__auto]))`,
			).Instant(true),
		)
}
