package squadron

import (
	pkgdashboard "github.com/foomo/grafana-dashboards/pkg/common/dashboard"
	pkgpanel "github.com/foomo/grafana-dashboards/pkg/common/panel"
	pkgvariable "github.com/foomo/grafana-dashboards/pkg/common/variable"
	pkgprometheus "github.com/foomo/grafana-dashboards/pkg/datasource/prometheus"
	"github.com/foomo/grafana-dashboards/pkg/util"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"
	"github.com/grafana/grafana-foundation-sdk/go/table"
	"github.com/grafana/grafana-foundation-sdk/go/units"
)

func NewReleasesDashboard() *dashboard.DashboardBuilder {
	return pkgdashboard.NewDashboard(
		"foomo_squadron_releases",
		"Squadron Releases",
		"Foomo Squadron Releases",
	).
		Variables(pkgdashboard.Variables{
			pkgvariable.NewNamespaceQuery("helm_chart_info"),
			pkgvariable.NewMultiQueryVariable("chart", "helm_chart_info").Label("📜"),
			pkgvariable.NewMultiQueryVariable("release", "helm_chart_info").Label("🚀"),
		}).
		Timepicker(dashboard.NewTimePickerBuilder().Hidden(true)).
		WithPanel(
			NewReleaseTablePanel().Span(24),
		)
}

// NewReleaseTablePanel returns a panel builder for the HTTP server duration panel.
func NewReleaseTablePanel() *table.PanelBuilder {
	return pkgpanel.NewTable("Releases", pkgprometheus.Datasource()).
		Datasource(pkgprometheus.Datasource()).
		WithTarget(
			pkgprometheus.NewDataQuery(
				`min(helm_chart_info{chart=~"$chart", namespace=~"$namespace", release=~"$release"}) by (updated, namespace, chart, release, description, version) != 2`,
			).Format(prometheus.PromQueryFormatTable).Instant(),
		).
		WithTransformation(dashboard.DataTransformerConfig{
			Id: "seriesToRows",
		}).
		WithTransformation(dashboard.DataTransformerConfig{
			Id: "extractFields",
			Options: util.MustYamlToMap(`
				delimiter: ","
				format: json
				jsonPaths:
				- alias: ""
				  path: squadron
				- alias: ""
				  path: user
				- alias: ""
				  path: commit
				- alias: ""
				  path: branch
				keepTime: false
				replace: false
				source: description
		`)}).
		WithTransformation(dashboard.DataTransformerConfig{
			Id: "organize",
			Options: util.MustYamlToMap(`
				excludeByName:
				  Time: true
				  chart: true
				  description: true
				  version: true
				indexByName:
				  Time: 0
				  updated: 1
				  namespace: 2
				  chart: 3
				  release: 4
				  version: 5
				  description: 6
				  squadron: 7
				  user: 8
				  commit: 9
				  branch: 10
				  Value: 11
				renameByName:
				  Value: Status
				  branch: Branch
				  chart: Chart
				  commit: Commit
				  description: DESCRIPTION
				  namespace: Namespace
				  release: Release
				  squadron: Squadron
				  updated: Update
				  user: User
				  version: ""
		`)}).
		OverrideByName("updated", []dashboard.DynamicConfigValue{
			{
				Id:    "unit",
				Value: units.DatetimeDefault,
			},
		}).
		OverrideByName("Value", []dashboard.DynamicConfigValue{
			{
				Id:    "custom.align",
				Value: "center",
			},
			{
				Id: "custom.cellOptions",
				Value: util.MustYamlToMap(`
					type: color-background
				`),
			},
			{
				Id: "thresholds",
				Value: util.MustYamlToMap(`
					mode: absolute
					steps:
					- color: rgba(245, 54, 54, 0.9)
					  value: 0
					- color: "#629e51"
					  value: 0
					- color: "#1f78c1"
					  value: 2
			`),
			},
		}).
		Mappings([]dashboard.ValueMapping{
			pkgdashboard.ValueMappingAsValueMap(map[string]dashboard.ValueMappingResult{
				"-1": {
					Index: util.ToPtr[int32](0),
					Text:  util.ToPtr("Failed"),
					Color: util.ToPtr("red"),
				},
				"0": {
					Index: util.ToPtr[int32](1),
					Text:  util.ToPtr("Unknown"),
					Color: util.ToPtr("yellow"),
				},
				"1": {
					Index: util.ToPtr[int32](2),
					Text:  util.ToPtr("Deployed"),
					Color: util.ToPtr("green"),
				},
				"2": {
					Index: util.ToPtr[int32](3),
					Text:  util.ToPtr("Deleted"),
					Color: util.ToPtr("yellow"),
				},
				"3": {
					Index: util.ToPtr[int32](4),
					Text:  util.ToPtr("Superseded"),
					Color: util.ToPtr("yellow"),
				},
				"5": {
					Index: util.ToPtr[int32](5),
					Text:  util.ToPtr("Deleting"),
					Color: util.ToPtr("yellow"),
				},
				"6": {
					Index: util.ToPtr[int32](6),
					Text:  util.ToPtr("Pending Install"),
					Color: util.ToPtr("orange"),
				},
				"7": {
					Index: util.ToPtr[int32](7),
					Text:  util.ToPtr("Pending Upgrade"),
					Color: util.ToPtr("orange"),
				},
				"8": {
					Index: util.ToPtr[int32](8),
					Text:  util.ToPtr("Pending Rollback"),
					Color: util.ToPtr("orange"),
				},
			}),
		}).
		Height(18)
}
