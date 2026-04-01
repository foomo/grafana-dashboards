package prometheus

import (
	pkgpanel "github.com/foomo/grafana-dashboards/pkg/common/panel"
	pkgprometheus "github.com/foomo/grafana-dashboards/pkg/datasource/prometheus"
	"github.com/foomo/grafana-dashboards/pkg/util"
	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"
	"github.com/grafana/grafana-foundation-sdk/go/stat"
	"github.com/grafana/grafana-foundation-sdk/go/table"
	"github.com/grafana/grafana-foundation-sdk/go/units"
)

func NewTotalMetrics() *stat.PanelBuilder {
	return pkgpanel.NewTotalStat("Metrics", pkgprometheus.Datasource()).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`count(count({namespace="$namespace", service="$service", __name__!=""}) by (__name__))`,
			).Instant(),
		)
}

func NewTotalSeries() *stat.PanelBuilder {
	return pkgpanel.NewTotalStat("Metrics", pkgprometheus.Datasource()).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`count({namespace="$namespace", service="$service", __name__!=""})`,
			).Instant(),
		)
}

func NewSeriesByMetrics() *table.PanelBuilder {
	return pkgpanel.NewTable("Series by metric", pkgprometheus.Datasource()).
		Footer(common.NewTableFooterOptionsBuilder().EnablePagination(true)).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`count({namespace="$namespace", service="$service", __name__!=""}) by (__name__)`,
			).Format(prometheus.PromQueryFormatTable).Instant(),
		).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`count({namespace="$namespace", service="$service", __name__!=""})`,
			).Format(prometheus.PromQueryFormatTable).Instant(),
		).
		WithTransformation(dashboard.DataTransformerConfig{
			Id: "merge",
		}).
		WithTransformation(dashboard.DataTransformerConfig{
			Id: "calculateField",
			Options: util.MustYamlToMap(`
        alias: Ratio
        binary:
          left: "Value #A"
          operator: /
          reducer: sum
          right: "Value #B"
        mode: binary
        reduce:
          reducer: sum
        replaceFields: false
		`)}).
		OverrideByName("Ratio", []dashboard.DynamicConfigValue{
			{
				Id:    "unit",
				Value: units.PercentUnit,
			},
			{
				Id:    "decimals",
				Value: "1",
			},
		}).
		WithTransformation(dashboard.DataTransformerConfig{
			Id: "organize",
			Options: util.MustYamlToMap(`
	      excludeByName:
					Time: true
					"Value #B": true
	      indexByName: {}
	      renameByName:
	        __name__: Metric
					"Value #A": "Series"
					"Ratio": "% of total"
			`)}).
		SortBy([]cog.Builder[common.TableSortByFieldState]{
			common.NewTableSortByFieldStateBuilder().DisplayName("% of total").Desc(true),
		})
}
