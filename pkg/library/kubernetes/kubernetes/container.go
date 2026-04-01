package kubernetes

import (
	pkgpanel "github.com/foomo/grafana-dashboards/pkg/common/panel"
	pkgprometheus "github.com/foomo/grafana-dashboards/pkg/datasource/prometheus"
	"github.com/foomo/grafana-dashboards/pkg/util"
	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"
	"github.com/grafana/grafana-foundation-sdk/go/table"
	"github.com/grafana/grafana-foundation-sdk/go/timeseries"
	"github.com/grafana/grafana-foundation-sdk/go/units"
)

func NewContainerInfos() *table.PanelBuilder {
	return pkgpanel.NewTable("Container details", pkgprometheus.Datasource()).
		NoValue("-").
		Targets([]cog.Builder[variants.Dataquery]{
			pkgprometheus.NewDataQuery(`
				sum(
					kube_pod_container_info{namespace="$namespace", pod=~"$pods"}
					* on(pod, container) group_left(reason) kube_pod_container_status_last_terminated_reason{namespace="$namespace", pod=~"$pods"}
				) by (pod, container, image, reason)
			`).Format(prometheus.PromQueryFormatTable).Instant(),
			pkgprometheus.NewDataQuery(`
				sum(increase(kube_pod_container_status_restarts_total{namespace="$namespace", pod=~"$pods", container!=""}[$__rate_interval])) by (pod, container)
			`),
			// TODO
			// pkgprometheus.NewDataQuery(`
			// 	sum(increase(container_oom_events_total{namespace="$namespace", pod=~"$pods", container!=""}[$__rate_interval])) by (pod, container)
			// `),
			// pkgprometheus.NewDataQuery(`
			// 	label_replace(
			// 		sum(kube_pod_container_status_last_terminated_reason{namespace="$namespace", pod=~"$pods"}) by (pod, container),
			// 		"reason", "unknown", "reason", "^$"
			// 	)
			// `).Format(prometheus.PromQueryFormatTable).Instant(),
		}).
		Transformations([]dashboard.DataTransformerConfig{
			{
				Id:    "timeSeriesTable",
				Topic: util.ToPtr(dashboard.DataTransformerConfigTopicSeries),
				Options: util.MustYamlToMap(`
	        B: { timeField: Time }
				`),
				Filter: &dashboard.MatcherConfig{
					Id:      "byRefId",
					Options: `/^(?:B)$/`,
				},
			},
			{
				Id: "merge",
			},
			{
				Id: "organize",
				Options: util.MustYamlToMap(`
	      excludeByName:
					Time: true
					Value: true
	      indexByName:
					container: 0
					pod: 1
					image: 3
	      renameByName:
					"Trend #B": "Restarts"
					"Ratio": "% of total"
					"reason": "Last Terminated Reason"
			`)},
		}).
		OverrideByName("Restarts", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.Number},
			{Id: "displayName", Value: "Restarts"},
			{Id: "custom.cellOptions", Value: util.MustJSONToMap(`{"type":"sparkline","stat":"total"}`)},
		})
	// OverrideByName("container", []dashboard.DynamicConfigValue{
	// 	{Id: "displayName", Value: "Container"},
	// }).
	// OverrideByName("pod", []dashboard.DynamicConfigValue{
	// 	{Id: "displayName", Value: "Pod"},
	// }).
	// OverrideByName("Value #A", []dashboard.DynamicConfigValue{
	// 	{Id: "unit", Value: units.Short},
	// 	{Id: "displayName", Value: "CPU Used"},
	// }).
	// OverrideByName("Value #B", []dashboard.DynamicConfigValue{
	// 	{Id: "unit", Value: units.Short},
	// 	{Id: "displayName", Value: "CPU Request"},
	// }).
	// OverrideByName("Value #C", []dashboard.DynamicConfigValue{
	// 	{Id: "unit", Value: units.Short},
	// 	{Id: "displayName", Value: "CPU Limit"},
	// }).
	// OverrideByName("Value #D", []dashboard.DynamicConfigValue{
	// 	{Id: "unit", Value: units.BytesIEC},
	// 	{Id: "displayName", Value: "Memory Used"},
	// 	{Id: "custom.align"},
	// }).
	// OverrideByName("Value #E", []dashboard.DynamicConfigValue{
	// 	{Id: "unit", Value: units.BytesIEC},
	// 	{Id: "displayName", Value: "Memory Request"},
	// }).
	// OverrideByName("Value #F", []dashboard.DynamicConfigValue{
	// 	{Id: "unit", Value: units.BytesIEC},
	// 	{Id: "displayName", Value: "Memory Limit"},
	// })
}

func NewQuotasByContainer() *table.PanelBuilder {
	return pkgpanel.NewTable("Quotas by container", pkgprometheus.Datasource()).
		NoValue("-").
		Targets([]cog.Builder[variants.Dataquery]{
			pkgprometheus.NewDataQuery(`
				max_over_time((sum(rate(container_cpu_usage_seconds_total{namespace="$namespace", image!="", container!="", pod=~"$pods"}[$__rate_interval])) by (pod, container))[$__range:])
			`).Format(prometheus.PromQueryFormatTable).Instant(),
			pkgprometheus.NewDataQuery(`
	      sum(kube_pod_container_resource_requests{namespace="$namespace", resource="cpu", pod=~"$pods"}) by (pod, container)
			`).Format(prometheus.PromQueryFormatTable).Instant(),
			pkgprometheus.NewDataQuery(`
				sum(kube_pod_container_resource_limits{namespace="$namespace", resource="cpu", pod=~"$pods"}) by (pod, container)
			`).Format(prometheus.PromQueryFormatTable).Instant(),
			pkgprometheus.NewDataQuery(`
				max_over_time((sum(container_memory_working_set_bytes{namespace="$namespace", image!="", container!="", pod=~"$pods"}) by (pod, container))[$__range:])
			`).Format(prometheus.PromQueryFormatTable).Instant(),
			pkgprometheus.NewDataQuery(`
	      sum(kube_pod_container_resource_requests{namespace="$namespace", resource="memory", pod=~"$pods"}) by (pod, container)
			`).Format(prometheus.PromQueryFormatTable).Instant(),
			pkgprometheus.NewDataQuery(`
				sum(kube_pod_container_resource_limits{namespace="$namespace", resource="memory", pod=~"$pods"}) by (pod, container)
			`).Format(prometheus.PromQueryFormatTable).Instant(),
		}).
		Transformations([]dashboard.DataTransformerConfig{
			{Id: "merge"},
		}).
		OverrideByName("Time", []dashboard.DynamicConfigValue{
			{Id: "custom.hideFrom.viz", Value: true},
		}).
		OverrideByName("container", []dashboard.DynamicConfigValue{
			{Id: "displayName", Value: "Container"},
		}).
		OverrideByName("pod", []dashboard.DynamicConfigValue{
			{Id: "displayName", Value: "Pod"},
		}).
		OverrideByName("Value #A", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.Short},
			{Id: "displayName", Value: "CPU Used"},
		}).
		OverrideByName("Value #B", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.Short},
			{Id: "displayName", Value: "CPU Request"},
		}).
		OverrideByName("Value #C", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.Short},
			{Id: "displayName", Value: "CPU Limit"},
		}).
		OverrideByName("Value #D", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.BytesIEC},
			{Id: "displayName", Value: "Memory Used"},
			{Id: "custom.align"},
		}).
		OverrideByName("Value #E", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.BytesIEC},
			{Id: "displayName", Value: "Memory Request"},
		}).
		OverrideByName("Value #F", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.BytesIEC},
			{Id: "displayName", Value: "Memory Limit"},
		})
}

func NewCPUQuotaByContainer() *table.PanelBuilder {
	return pkgpanel.NewTable("CPU Quota by container", pkgprometheus.Datasource()).
		NoValue("-").
		Targets([]cog.Builder[variants.Dataquery]{
			pkgprometheus.NewDataQuery(`
				max_over_time((sum(rate(container_cpu_usage_seconds_total{namespace="$namespace", image!="", container!="", pod=~"$pods"}[$__rate_interval])) by (pod, container))[$__range:])
			`).Format(prometheus.PromQueryFormatTable).Instant(),
			pkgprometheus.NewDataQuery(`
	      sum(kube_pod_container_resource_requests{namespace="$namespace", resource="cpu", pod=~"$pods"}) by (pod, container)
			`).Format(prometheus.PromQueryFormatTable).Instant(),
			pkgprometheus.NewDataQuery(`
				max_over_time((sum(rate(container_cpu_usage_seconds_total{namespace="$namespace", image!="", container!="", pod=~"$pods"}[$__rate_interval])) by (pod, container))[$__range:])
				/
	      sum(kube_pod_container_resource_requests{namespace="$namespace", resource="cpu", pod=~"$pods"}) by (pod, container)
			`).Format(prometheus.PromQueryFormatTable).Instant(),
			pkgprometheus.NewDataQuery(`
				sum(kube_pod_container_resource_limits{namespace="$namespace", resource="cpu", pod=~"$pods"}) by (pod, container)
			`).Format(prometheus.PromQueryFormatTable).Instant(),
			pkgprometheus.NewDataQuery(`
				max_over_time((sum(rate(container_cpu_usage_seconds_total{namespace="$namespace", image!="", container!="", pod=~"$pods"}[$__rate_interval])) by (pod, container))[$__range:])
				/
				sum(kube_pod_container_resource_limits{namespace="$namespace", resource="cpu", pod=~"$pods"}) by (pod, container)
			`).Format(prometheus.PromQueryFormatTable).Instant(),
		}).
		Transformations([]dashboard.DataTransformerConfig{
			{Id: "merge"},
		}).
		OverrideByName("Time", []dashboard.DynamicConfigValue{
			{Id: "custom.hideFrom.viz", Value: true},
		}).
		OverrideByName("container", []dashboard.DynamicConfigValue{
			{Id: "displayName", Value: "Container"},
		}).
		OverrideByName("pod", []dashboard.DynamicConfigValue{
			{Id: "displayName", Value: "Pod"},
		}).
		OverrideByName("Value #A", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.Short},
			{Id: "displayName", Value: "Used"},
		}).
		OverrideByName("Value #B", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.Short},
			{Id: "displayName", Value: "Request"},
		}).
		OverrideByName("Value #C", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.PercentUnit},
			{Id: "custom.align", Value: "right"},
			{Id: "displayName", Value: "Requests %"},
			{Id: "custom.cellOptions", Value: util.MustJSONToMap(`{"mode":"lcd","type":"gauge","valueDisplayMode":"color"}`)},
			{Id: "thresholds", Value: util.MustJSONToMap(`{"mode":"absolute","steps":[{"color":"red","value":0},{"color":"yellow","value":0.2},{"color":"green","value":0.3},{"color":"yellow","value":0.7},{"color":"red","value":0.8}]}`)},
			{Id: "min", Value: "0"},
			{Id: "max", Value: "1"},
		}).
		OverrideByName("Value #D", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.Short},
			{Id: "displayName", Value: "Limits"},
			{Id: "custom.align"},
		}).
		OverrideByName("Value #E", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.PercentUnit},
			{Id: "custom.align", Value: "right"},
			{Id: "displayName", Value: "Limits %"},
			{Id: "custom.cellOptions", Value: util.MustJSONToMap(`{"mode":"lcd","type":"gauge","valueDisplayMode":"color"}`)},
			{Id: "thresholds", Value: util.MustJSONToMap(`{"mode":"absolute","steps":[{"color":"red","value":0},{"color":"yellow","value":0.2},{"color":"green","value":0.3},{"color":"yellow","value":0.7},{"color":"red","value":0.8}]}`)},
			{Id: "min", Value: "0"},
			{Id: "max", Value: "1"},
		})
}

func NewMemoryQuotaByContainer() *table.PanelBuilder {
	return pkgpanel.NewTable("Memory Quota by container", pkgprometheus.Datasource()).
		NoValue("-").
		Targets([]cog.Builder[variants.Dataquery]{
			pkgprometheus.NewDataQuery(`
				max_over_time((sum(container_memory_working_set_bytes{namespace="$namespace", image!="", container!="", pod=~"$pods"}) by (pod, container))[$__range:])
			`).Format(prometheus.PromQueryFormatTable).Instant(),
			pkgprometheus.NewDataQuery(`
	      sum(kube_pod_container_resource_requests{namespace="$namespace", resource="memory", pod=~"$pods"}) by (pod, container)
			`).Format(prometheus.PromQueryFormatTable).Instant(),
			pkgprometheus.NewDataQuery(`
				max_over_time((sum(container_memory_working_set_bytes{namespace="$namespace", image!="", container!="", pod=~"$pods"}) by (pod, container))[$__range:])
				/
	      sum(kube_pod_container_resource_requests{namespace="$namespace", resource="memory", pod=~"$pods"}) by (pod, container)
			`).Format(prometheus.PromQueryFormatTable).Instant(),
			pkgprometheus.NewDataQuery(`
				sum(kube_pod_container_resource_limits{namespace="$namespace", resource="memory", pod=~"$pods"}) by (pod, container)
			`).Format(prometheus.PromQueryFormatTable).Instant(),
			pkgprometheus.NewDataQuery(`
				max_over_time((sum(container_memory_working_set_bytes{namespace="$namespace", image!="", container!="", pod=~"$pods"}) by (pod, container))[$__range:])
				/
				sum(kube_pod_container_resource_limits{namespace="$namespace", resource="memory", pod=~"$pods"}) by (pod, container)
			`).Format(prometheus.PromQueryFormatTable).Instant(),
		}).
		Transformations([]dashboard.DataTransformerConfig{
			{Id: "merge"},
		}).
		OverrideByName("Time", []dashboard.DynamicConfigValue{
			{Id: "custom.hideFrom.viz", Value: true},
		}).
		OverrideByName("container", []dashboard.DynamicConfigValue{
			{Id: "displayName", Value: "Container"},
		}).
		OverrideByName("pod", []dashboard.DynamicConfigValue{
			{Id: "displayName", Value: "Pod"},
		}).
		OverrideByName("Value #A", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.BytesIEC},
			{Id: "displayName", Value: "Used"},
			{Id: "custom.align"},
		}).
		OverrideByName("Value #B", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.BytesIEC},
			{Id: "displayName", Value: "Requests"},
		}).
		OverrideByName("Value #C", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.PercentUnit},
			{Id: "custom.align", Value: "right"},
			{Id: "displayName", Value: "Requests %"},
			{Id: "custom.cellOptions", Value: util.MustJSONToMap(`{"mode":"lcd","type":"gauge","valueDisplayMode":"color"}`)},
			{Id: "thresholds", Value: util.MustJSONToMap(`{"mode":"absolute","steps":[{"color":"red","value":0},{"color":"yellow","value":0.2},{"color":"green","value":0.3},{"color":"yellow","value":0.7},{"color":"red","value":0.8}]}`)},
			{Id: "min", Value: "0"},
			{Id: "max", Value: "1"},
		}).
		OverrideByName("Value #D", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.BytesIEC},
			{Id: "displayName", Value: "Limits"},
		}).
		OverrideByName("Value #E", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.PercentUnit},
			{Id: "custom.align", Value: "right"},
			{Id: "displayName", Value: "Limits %"},
			{Id: "custom.cellOptions", Value: util.MustJSONToMap(`{"mode":"lcd","type":"gauge","valueDisplayMode":"color"}`)},
			{Id: "thresholds", Value: util.MustJSONToMap(`{"mode":"absolute","steps":[{"color":"red","value":0},{"color":"yellow","value":0.2},{"color":"green","value":0.3},{"color":"yellow","value":0.7},{"color":"red","value":0.8}]}`)},
			{Id: "min", Value: "0"},
			{Id: "max", Value: "1"},
		})
}

func NewCPUResourcesByContainer() *timeseries.PanelBuilder {
	return pkgpanel.NewTimeSeries("CPU Requests & Limits by container", pkgprometheus.Datasource()).
		Unit(units.PercentUnit).
		Legend(
			common.NewVizLegendOptionsBuilder().
				ShowLegend(true).
				DisplayMode(common.LegendDisplayModeTable).
				Calcs([]string{"mean", "min", "max"}).
				SortBy("Mean").
				SortDesc(true),
		).
		Min(0).
		Max(1).
		LineWidth(2.0).
		FillOpacity(0.0).
		GradientMode(common.GraphGradientModeNone).
		Thresholds(dashboard.NewThresholdsConfigBuilder().Steps([]dashboard.Threshold{
			{Value: util.ToPtr(0.0), Color: "red"},
			{Value: util.ToPtr(20.0), Color: "yellow"},
			{Value: util.ToPtr(30.0), Color: "green"},
			{Value: util.ToPtr(70.0), Color: "yellow"},
			{Value: util.ToPtr(80.0), Color: "red"},
		}).Mode(dashboard.ThresholdsModePercentage)).
		ThresholdsStyle(common.NewGraphThresholdsStyleConfigBuilder().Mode(common.GraphThresholdsStyleModeArea)).
		WithTarget(
			pkgprometheus.NewDataQuery(`
				sum(rate(container_cpu_usage_seconds_total{namespace="$namespace", pod=~"$pods", image!=""}[$__rate_interval])) by (container)
				/
       sum(kube_pod_container_resource_requests{namespace="$namespace", pod=~"$pods", resource="cpu"}) by (container)
			`).LegendFormat("{{container}} REQUESTS"),
		).
		WithTarget(
			pkgprometheus.NewDataQuery(`
        sum(rate(container_cpu_usage_seconds_total{namespace="$namespace", pod=~"$pods", image!=""}[$__rate_interval])) by (container)
				/
        sum(kube_pod_container_resource_limits{namespace="$namespace", pod=~"$pods", resource="cpu"}) by (container)
			`).LegendFormat("{{container}} LIMITS"),
		).
		OverrideByQuery("A", []dashboard.DynamicConfigValue{
			{Id: "color", Value: util.MustJSONToMap(`{"fixedColor":"blue","mode":"shades"}`)},
		}).
		OverrideByQuery("B", []dashboard.DynamicConfigValue{
			{Id: "color", Value: util.MustJSONToMap(`{"fixedColor":"semi-dark-orange","mode":"shades"}`)},
		})
}

func NewMemoryResourcesByContainer() *timeseries.PanelBuilder {
	return pkgpanel.NewTimeSeries("Memory Requests & Limits by container", pkgprometheus.Datasource()).
		Unit(units.PercentUnit).
		Legend(
			common.NewVizLegendOptionsBuilder().
				ShowLegend(true).
				DisplayMode(common.LegendDisplayModeTable).
				Calcs([]string{"mean", "min", "max"}).
				SortBy("Mean").
				SortDesc(true),
		).
		Min(0).
		Max(1).
		LineWidth(2.0).
		FillOpacity(0.0).
		GradientMode(common.GraphGradientModeNone).
		Thresholds(dashboard.NewThresholdsConfigBuilder().Steps([]dashboard.Threshold{
			{Value: util.ToPtr(0.0), Color: "red"},
			{Value: util.ToPtr(20.0), Color: "yellow"},
			{Value: util.ToPtr(30.0), Color: "green"},
			{Value: util.ToPtr(70.0), Color: "yellow"},
			{Value: util.ToPtr(80.0), Color: "red"},
		}).Mode(dashboard.ThresholdsModePercentage)).
		ThresholdsStyle(common.NewGraphThresholdsStyleConfigBuilder().Mode(common.GraphThresholdsStyleModeArea)).
		WithTarget(
			pkgprometheus.NewDataQuery(`
        sum(container_memory_working_set_bytes{namespace="$namespace", pod=~"$pods", image!=""}) by (container)
				/
        sum(kube_pod_container_resource_requests{namespace="$namespace", pod=~"$pods", resource="memory"}) by (container)
			`).LegendFormat("{{container}} REQUESTS"),
		).
		WithTarget(
			pkgprometheus.NewDataQuery(`
        sum(container_memory_working_set_bytes{namespace="$namespace", pod=~"$pods", image!=""}) by (container)
				/
        sum(kube_pod_container_resource_limits{namespace="$namespace", pod=~"$pods", resource="memory"}) by (container)
			`).LegendFormat("{{container}} LIMITS"),
		).
		OverrideByQuery("A", []dashboard.DynamicConfigValue{
			{Id: "color", Value: util.MustJSONToMap(`{"fixedColor":"blue","mode":"shades"}`)},
		}).
		OverrideByQuery("B", []dashboard.DynamicConfigValue{
			{Id: "color", Value: util.MustJSONToMap(`{"fixedColor":"semi-dark-orange","mode":"shades"}`)},
		})
}

func NewCPUUsageByContainer() *timeseries.PanelBuilder {
	return pkgpanel.NewTimeSeries("CPU Usage by container", pkgprometheus.Datasource()).
		Unit(units.Short).
		AxisLabel("CPU Cores").
		Legend(
			common.NewVizLegendOptionsBuilder().
				ShowLegend(true).
				DisplayMode(common.LegendDisplayModeTable).
				Calcs([]string{"mean", "max"}).
				SortBy("Max").
				SortDesc(true),
		).
		WithTarget(
			pkgprometheus.NewDataQuery(`
        sum(rate(container_cpu_usage_seconds_total{namespace="$namespace", pod=~"$pods", image!="", container!=""}[$__rate_interval])) by (container)
			`).LegendFormat("{{container}}").IntervalFactor(2),
		)
}

func NewMemoryUsageByContainer() *timeseries.PanelBuilder {
	return pkgpanel.NewTimeSeries("Memory Usage by container", pkgprometheus.Datasource()).
		Unit(units.BytesIEC).
		AxisLabel("Bytes").
		Legend(
			common.NewVizLegendOptionsBuilder().
				ShowLegend(true).
				DisplayMode(common.LegendDisplayModeTable).
				Calcs([]string{"mean", "max"}).
				SortBy("Max").
				SortDesc(true),
		).
		WithTarget(
			pkgprometheus.NewDataQuery(`
        sum(container_memory_working_set_bytes{namespace="$namespace", pod=~"$pods", image!="", container!=""}) by (container)
			`).LegendFormat("{{container}}"),
		)
}

func NewNetworkIOByContainer() *table.PanelBuilder {
	return pkgpanel.NewTable("Network IO by container", pkgprometheus.Datasource()).
		NoValue("-").
		Targets([]cog.Builder[variants.Dataquery]{
			pkgprometheus.NewDataQuery(`
				sum(increase(container_network_receive_bytes_total{namespace="$namespace", pod=~"$pods"}[$__range])) by (pod, container)
			`).Format(prometheus.PromQueryFormatTable).Instant(),
			pkgprometheus.NewDataQuery(`
				sum(increase(container_network_transmit_bytes_total{namespace="$namespace", pod=~"$pods"}[$__range])) by (pod, container)
			`).Format(prometheus.PromQueryFormatTable).Instant(),
			pkgprometheus.NewDataQuery(`
				sum(rate(container_network_receive_bytes_total{namespace="$namespace", pod=~"$pods"}[$__rate_interval])) by (pod, container)
			`),
			pkgprometheus.NewDataQuery(`
				sum(rate(container_network_transmit_bytes_total{namespace="$namespace", pod=~"$pods"}[$__rate_interval])) by (pod, container)
			`),
		}).
		Transformations([]dashboard.DataTransformerConfig{
			{
				Id:    "timeSeriesTable",
				Topic: util.ToPtr(dashboard.DataTransformerConfigTopicSeries),
				Options: util.MustYamlToMap(`
	        C: { timeField: Time }
	        D: { timeField: Time }
				`),
				Filter: &dashboard.MatcherConfig{
					Id:      "byRefId",
					Options: `/^(?:C|D)$/`,
				},
			},
			{
				Id: "merge",
			},
			{
				Id: "organize",
				Options: util.MustYamlToMap(`
	      excludeByName:
					Time: true
	      indexByName:
					pod: 0
					"Value #A": 1
					"Value #B": 3
					"Trend #C": 2
					"Trend #D": 4
	      renameByName:
					pod: Pod
					"Value #A": "Received"
					"Value #B": "Transmitted"
					"Trend #C": "Receive Rate"
					"Trend #D": "Transmit Rate"
			`)},
		}).
		OverrideByName("Received", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.BytesIEC},
			{Id: "custom.align", Value: "right"},
		}).
		OverrideByName("Receive Rate", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.BytesPerSecondSI},
			{Id: "custom.cellOptions", Value: util.MustJSONToMap(`{"type":"sparkline","stat":"mean"}`)},
			{Id: "custom.align", Value: "right"},
		}).
		OverrideByName("Transmitted", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.BytesIEC},
			{Id: "custom.align", Value: "right"},
		}).
		OverrideByName("Transmit Rate", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.BytesPerSecondSI},
			{Id: "custom.cellOptions", Value: util.MustJSONToMap(`{"type":"sparkline","stat":"mean"}`)},
			{Id: "custom.align", Value: "right"},
		})
	// OverrideByName("container", []dashboard.DynamicConfigValue{
	// 	{Id: "displayName", Value: "Container"},
	// }).
	// OverrideByName("pod", []dashboard.DynamicConfigValue{
	// 	{Id: "displayName", Value: "Pod"},
	// }).
	// OverrideByName("Value #A", []dashboard.DynamicConfigValue{
	// 	{Id: "unit", Value: units.Short},
	// 	{Id: "displayName", Value: "CPU Used"},
	// }).
	// OverrideByName("Value #B", []dashboard.DynamicConfigValue{
	// 	{Id: "unit", Value: units.Short},
	// 	{Id: "displayName", Value: "CPU Request"},
	// }).
	// OverrideByName("Value #C", []dashboard.DynamicConfigValue{
	// 	{Id: "unit", Value: units.Short},
	// 	{Id: "displayName", Value: "CPU Limit"},
	// }).
	// OverrideByName("Value #D", []dashboard.DynamicConfigValue{
	// 	{Id: "unit", Value: units.BytesIEC},
	// 	{Id: "displayName", Value: "Memory Used"},
	// 	{Id: "custom.align"},
	// }).
	// OverrideByName("Value #E", []dashboard.DynamicConfigValue{
	// 	{Id: "unit", Value: units.BytesIEC},
	// 	{Id: "displayName", Value: "Memory Request"},
	// }).
	// OverrideByName("Value #F", []dashboard.DynamicConfigValue{
	// 	{Id: "unit", Value: units.BytesIEC},
	// 	{Id: "displayName", Value: "Memory Limit"},
	// })
}
