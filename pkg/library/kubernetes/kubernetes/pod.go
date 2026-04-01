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

func NewLastTerminationReason() string {
	return "k8s/last-termination-reason"
}

func NewCPUUsageByPod() *timeseries.PanelBuilder {
	return pkgpanel.NewTimeSeries("CPU Usage by pod", pkgprometheus.Datasource()).
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
		Thresholds(dashboard.NewThresholdsConfigBuilder().Steps([]dashboard.Threshold{
			{
				Value: util.ToPtr(0.0),
				Color: "transparent",
			},
		})).
		ThresholdsStyle(common.NewGraphThresholdsStyleConfigBuilder().Mode(common.GraphThresholdsStyleModeDashedAndArea)).
		Targets([]cog.Builder[variants.Dataquery]{
			pkgprometheus.NewDataQuery(`
				sum(
				    node_namespace_pod_container:container_cpu_usage_seconds_total:sum_irate{namespace="$namespace"}
				  * on(namespace,pod)
				    group_left(workload) namespace_workload_pod:kube_pod_owner:relabel{namespace="$namespace", workload="$service"}
				) by (pod)
			`).LegendFormat("{{pod}}").IntervalFactor(2),
			pkgprometheus.NewDataQuery(`
	      sum(
	          kube_pod_container_resource_requests{namespace="$namespace", resource="cpu"}
	        * on(namespace,pod)
	          group_left(workload) namespace_workload_pod:kube_pod_owner:relabel{namespace="$namespace", workload="$service"}
	      ) by (pod)
			`).Format(prometheus.PromQueryFormatTable).Instant(),
			pkgprometheus.NewDataQuery(`
        sum(
            kube_pod_container_resource_limits{namespace="$namespace", resource="cpu"}
          * on(namespace,pod)
            group_left(workload) namespace_workload_pod:kube_pod_owner:relabel{namespace="$namespace", workload="$service"}
        ) by (pod)
			`).Format(prometheus.PromQueryFormatTable).Instant(),
		}).
		Transformations([]dashboard.DataTransformerConfig{
			{
				Id: "configFromData",
				Options: util.MustYamlToMap(`
        configRefId: B
        mappings:
        - fieldName: "Value #B"
          handlerArguments:
						threshold:
							color: blue
					handlerKey: threshold1
			`),
			},
			{
				Id: "configFromData",
				Options: util.MustYamlToMap(`
        configRefId: C
        mappings:
        - fieldName: "Value #C"
          handlerArguments:
						threshold:
							color: red
					handlerKey: threshold1
			`),
			},
		})
}

func NewCPUQuotaByPod() *table.PanelBuilder {
	return pkgpanel.NewTable("CPU Quota by pod", pkgprometheus.Datasource()).
		NoValue("-").
		Targets([]cog.Builder[variants.Dataquery]{
			pkgprometheus.NewDataQuery(`
        max_over_time((sum(
            node_namespace_pod_container:container_cpu_usage_seconds_total:sum_irate{namespace="$namespace"}
          * on(namespace,pod)
            group_left(workload) namespace_workload_pod:kube_pod_owner:relabel{namespace="$namespace", workload="$service"}
        ) by (pod))[$__range:])
			`).Format(prometheus.PromQueryFormatTable).Instant(),
			pkgprometheus.NewDataQuery(`
	      sum(
	          kube_pod_container_resource_requests{namespace="$namespace", resource="cpu"}
	        * on(namespace,pod)
	          group_left(workload) namespace_workload_pod:kube_pod_owner:relabel{namespace="$namespace", workload="$service"}
	      ) by (pod)
			`).Format(prometheus.PromQueryFormatTable).Instant(),
			pkgprometheus.NewDataQuery(`
        max_over_time((sum(
            node_namespace_pod_container:container_cpu_usage_seconds_total:sum_irate{namespace="$namespace"}
          * on(namespace,pod)
            group_left(workload) namespace_workload_pod:kube_pod_owner:relabel{namespace="$namespace", workload="$service"}
        ) by (pod))[$__range:])
        / sum(
            kube_pod_container_resource_requests{namespace="$namespace", resource="cpu"}
          * on(namespace,pod)
            group_left(workload) namespace_workload_pod:kube_pod_owner:relabel{namespace="$namespace", workload="$service"}
        ) by (pod)
			`).Format(prometheus.PromQueryFormatTable).Instant(),
			pkgprometheus.NewDataQuery(`
        sum(
            kube_pod_container_resource_limits{namespace="$namespace", resource="cpu"}
          * on(namespace,pod)
            group_left(workload) namespace_workload_pod:kube_pod_owner:relabel{namespace="$namespace", workload="$service"}
        ) by (pod)
			`).Format(prometheus.PromQueryFormatTable).Instant(),
			pkgprometheus.NewDataQuery(`
        max_over_time((sum(
            node_namespace_pod_container:container_cpu_usage_seconds_total:sum_irate{namespace="$namespace"}
          * on(namespace,pod)
            group_left(workload) namespace_workload_pod:kube_pod_owner:relabel{namespace="$namespace", workload="$service"}
        ) by (pod))[$__range:])
        / sum(
            kube_pod_container_resource_limits{namespace="$namespace", resource="cpu"}
          * on(namespace,pod)
            group_left(workload) namespace_workload_pod:kube_pod_owner:relabel{namespace="$namespace", workload="$service"}
        ) by (pod)
			`).Format(prometheus.PromQueryFormatTable).Instant(),
		}).
		Transformations([]dashboard.DataTransformerConfig{
			{
				Id: "merge",
			},
		}).
		OverrideByName("pod", []dashboard.DynamicConfigValue{
			{Id: "displayName", Value: "Pod"},
		}).
		OverrideByName("Time", []dashboard.DynamicConfigValue{
			{Id: "custom.hideFrom.viz", Value: true},
		}).
		OverrideByName("Value #A", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.Short},
			{Id: "displayName", Value: "Usage"},
		}).
		OverrideByName("Value #B", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.Short},
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

// TODO
// func NewCPUThrottling() *histogram.PanelBuilder {
//	return pkgpanel.NewHistogram("CPU Throttling", pkgprometheus.Datasource()).
//		WithTarget(
//			pkgprometheus.NewDataQuery(`
//        sum(increase(container_cpu_cfs_throttled_periods_total{namespace="$namespace", service="$service", container!=""}[$__interval])) by (container)
//				/
//				sum(increase(container_cpu_cfs_periods_total{namespace="$namespace", pod="$pod", container!=""}[$__interval])) by (container)
//			`).Format(prometheus.PromQueryFormatTable).Instant(),
//		)
// }

func NewMemoryUsageByPod() *timeseries.PanelBuilder {
	return pkgpanel.NewTimeSeries("Memory Usage by pod", pkgprometheus.Datasource()).
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
		Thresholds(dashboard.NewThresholdsConfigBuilder().Steps([]dashboard.Threshold{
			{
				Value: util.ToPtr(0.0),
				Color: "transparent",
			},
		})).
		ThresholdsStyle(common.NewGraphThresholdsStyleConfigBuilder().Mode(common.GraphThresholdsStyleModeDashedAndArea)).
		WithTarget(
			pkgprometheus.NewDataQuery(`
        sum(
            container_memory_working_set_bytes{namespace="$namespace", container!="", image!=""}
          * on(namespace,pod)
            group_left(workload) namespace_workload_pod:kube_pod_owner:relabel{namespace="$namespace", workload="$service"}
        ) by (pod)
			`).LegendFormat("{{pod}}").IntervalFactor(2),
		).
		WithTarget(
			pkgprometheus.NewDataQuery(`
        sum(
            kube_pod_container_resource_requests{namespace="$namespace", resource="memory"}
          * on(namespace,pod)
            group_left(workload) namespace_workload_pod:kube_pod_owner:relabel{namespace="$namespace", workload="$service"}
        ) by (pod)
			`).Format(prometheus.PromQueryFormatTable).Instant(),
		).
		WithTarget(
			pkgprometheus.NewDataQuery(`
        sum(
            kube_pod_container_resource_limits{namespace="$namespace", resource="memory"}
          * on(namespace,pod)
            group_left(workload) namespace_workload_pod:kube_pod_owner:relabel{namespace="$namespace", workload="$service"}
        ) by (pod)
			`).Format(prometheus.PromQueryFormatTable).Instant(),
		).
		WithTransformation(dashboard.DataTransformerConfig{
			Id: "configFromData",
			Options: util.MustYamlToMap(`
        configRefId: B
        mappings:
        - fieldName: "Value #B"
          handlerArguments:
						threshold:
							color: blue
					handlerKey: threshold1
			`),
		}).
		WithTransformation(dashboard.DataTransformerConfig{
			Id: "configFromData",
			Options: util.MustYamlToMap(`
        configRefId: C
        mappings:
        - fieldName: "Value #C"
          handlerArguments:
						threshold:
							color: red
					handlerKey: threshold1
			`),
		})
}

func NewMemoryQuotaByPod() *table.PanelBuilder {
	return pkgpanel.NewTable("Memory Quota by pod", pkgprometheus.Datasource()).
		WithTarget(
			pkgprometheus.NewDataQuery(`
        max_over_time((sum(
            container_memory_working_set_bytes{namespace="$namespace", container!="", image!=""}
          * on(namespace,pod)
            group_left(workload, workload_type) namespace_workload_pod:kube_pod_owner:relabel{namespace="$namespace", workload="$service"}
        ) by (pod))[$__range:])
			`).Format(prometheus.PromQueryFormatTable).Instant(),
		).
		WithTarget(
			pkgprometheus.NewDataQuery(`
	      sum(
	          kube_pod_container_resource_requests{namespace="$namespace", resource="memory"}
	        * on(namespace,pod)
	          group_left(workload) namespace_workload_pod:kube_pod_owner:relabel{namespace="$namespace", workload="$service"}
	      ) by (pod)
			`).Format(prometheus.PromQueryFormatTable).Instant(),
		).
		WithTarget(
			pkgprometheus.NewDataQuery(`
        max_over_time((sum(
            container_memory_working_set_bytes{namespace="$namespace", container!="", image!=""}
          * on(namespace,pod)
            group_left(workload) namespace_workload_pod:kube_pod_owner:relabel{namespace="$namespace", workload="$service"}
        ) by (pod))[$__range:])
        /
				sum(
            kube_pod_container_resource_requests{namespace="$namespace", resource="memory"}
          * on(namespace,pod)
            group_left(workload) namespace_workload_pod:kube_pod_owner:relabel{namespace="$namespace", workload="$service"}
        ) by (pod)
			`).Format(prometheus.PromQueryFormatTable).Instant(),
		).
		WithTarget(
			pkgprometheus.NewDataQuery(`
        sum(
            kube_pod_container_resource_limits{namespace="$namespace", resource="memory"}
          * on(namespace,pod)
            group_left(workload) namespace_workload_pod:kube_pod_owner:relabel{namespace="$namespace", workload="$service"}
        ) by (pod)
			`).Format(prometheus.PromQueryFormatTable).Instant(),
		).
		WithTarget(
			pkgprometheus.NewDataQuery(`
        max_over_time((sum(
            container_memory_working_set_bytes{namespace="$namespace", container!="", image!=""}
          * on(namespace,pod)
            group_left(workload) namespace_workload_pod:kube_pod_owner:relabel{namespace="$namespace", workload="$service"}
        ) by (pod))[$__range:])
        /
				sum(
            kube_pod_container_resource_limits{namespace="$namespace", resource="memory"}
          * on(namespace,pod)
            group_left(workload) namespace_workload_pod:kube_pod_owner:relabel{namespace="$namespace", workload="$service"}
        ) by (pod)
			`).Format(prometheus.PromQueryFormatTable).Instant(),
		).
		NoValue("-").
		WithTransformation(dashboard.DataTransformerConfig{
			Id: "merge",
		}).
		OverrideByName("Time", []dashboard.DynamicConfigValue{
			{Id: "custom.hideFrom.viz", Value: true},
		}).
		OverrideByName("pod", []dashboard.DynamicConfigValue{
			{Id: "displayName", Value: "Pod"},
		}).
		OverrideByName("Value #A", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.BytesIEC},
			{Id: "decimals", Value: "2"},
			{Id: "displayName", Value: "Usage"},
		}).
		OverrideByName("Value #B", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.BytesIEC},
			{Id: "decimals", Value: "2"},
			{Id: "displayName", Value: "Requests"},
		}).
		OverrideByName("Value #C", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.PercentUnit},
			{Id: "decimals", Value: "2"},
			{Id: "displayName", Value: "Requests %"},
			{Id: "custom.cellOptions", Value: util.MustJSONToMap(`{"mode":"lcd","type":"gauge","valueDisplayMode":"color"}`)},
			{Id: "thresholds", Value: util.MustJSONToMap(`{"mode":"absolute","steps":[{"color":"red","value":0},{"color":"yellow","value":0.2},{"color":"green","value":0.3},{"color":"yellow","value":0.7},{"color":"red","value":0.8}]}`)},
			{Id: "min", Value: "0"},
			{Id: "max", Value: "1"},
		}).
		OverrideByName("Value #D", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.BytesIEC},
			{Id: "decimals", Value: "2"},
			{Id: "displayName", Value: "Limits"},
			{Id: "custom.align"},
		}).
		OverrideByName("Value #E", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.PercentUnit},
			{Id: "decimals", Value: "2"},
			{Id: "displayName", Value: "Limits %"},
			{Id: "custom.cellOptions", Value: util.MustJSONToMap(`{"mode":"lcd","type":"gauge","valueDisplayMode":"color"}`)},
			{Id: "thresholds", Value: util.MustJSONToMap(`{"mode":"absolute","steps":[{"color":"red","value":0},{"color":"yellow","value":0.2},{"color":"green","value":0.3},{"color":"yellow","value":0.7},{"color":"red","value":0.8}]}`)},
			{Id: "min", Value: "0"},
			{Id: "max", Value: "1"},
		})
}

func NewNetworkUsageByPod() *table.PanelBuilder {
	return pkgpanel.NewTable("Network Usage by pod", pkgprometheus.Datasource()).
		WithTarget(
			pkgprometheus.NewDataQuery(`
        (
				sum(irate(container_network_receive_bytes_total{namespace=~"$namespace"}[$__rate_interval]) * on (namespace,pod)
        group_left(workload,workload_type)
        namespace_workload_pod:kube_pod_owner:relabel{namespace=~"$namespace", workload=~"$service"}
				) by (pod))
			`).Format(prometheus.PromQueryFormatTable).Instant(),
		).
		WithTarget(
			pkgprometheus.NewDataQuery(`
	      (
					sum(irate(container_network_transmit_bytes_total{namespace=~"$namespace"}[$__rate_interval]) * on (namespace,pod)
	        group_left(workload,workload_type)
	        namespace_workload_pod:kube_pod_owner:relabel{namespace=~"$namespace", workload=~"$service"}
				) by (pod))
			`).Format(prometheus.PromQueryFormatTable).Instant(),
		).
		WithTarget(
			pkgprometheus.NewDataQuery(`
        (
					sum(irate(container_network_receive_packets_total{namespace=~"$namespace"}[$__rate_interval]) * on (namespace,pod)
					group_left(workload,workload_type)
          namespace_workload_pod:kube_pod_owner:relabel{namespace=~"$namespace", workload=~"$service"}
				) by (pod))
			`).Format(prometheus.PromQueryFormatTable).Instant(),
		).
		WithTarget(
			pkgprometheus.NewDataQuery(`
        (
					sum(irate(container_network_transmit_packets_total{namespace=~"$namespace"}[$__rate_interval]) * on (namespace,pod)
          group_left(workload,workload_type)
					namespace_workload_pod:kube_pod_owner:relabel{namespace=~"$namespace", workload=~"$service"}
				) by (pod))
			`).Format(prometheus.PromQueryFormatTable).Instant(),
		).
		WithTarget(
			pkgprometheus.NewDataQuery(`
        (
					sum(irate(container_network_receive_packets_dropped_total{namespace=~"$namespace"}[$__rate_interval]) * on (namespace,pod)
					group_left(workload,workload_type)
					namespace_workload_pod:kube_pod_owner:relabel{namespace=~"$namespace", workload=~"$service"}
				) by (pod))
			`).Format(prometheus.PromQueryFormatTable).Instant(),
		).
		WithTarget(
			pkgprometheus.NewDataQuery(`
        (
					sum(irate(container_network_transmit_packets_dropped_total{namespace=~"$namespace"}[$__rate_interval]) * on (namespace,pod)
          group_left(workload,workload_type)
	        namespace_workload_pod:kube_pod_owner:relabel{namespace=~"$namespace", workload=~"$service"}
				) by (pod))
			`).Format(prometheus.PromQueryFormatTable).Instant(),
		).
		NoValue("-").
		WithTransformation(dashboard.DataTransformerConfig{
			Id: "merge",
		}).
		OverrideByName("Time", []dashboard.DynamicConfigValue{
			{Id: "custom.hideFrom.viz", Value: true},
		}).
		OverrideByName("pod", []dashboard.DynamicConfigValue{
			{Id: "displayName", Value: "Pod"},
		}).
		OverrideByName("Value #A", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.BitsPerSecondSI},
			{Id: "displayName", Value: "Receive Bandwidth"},
		}).
		OverrideByName("Value #B", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.BitsPerSecondSI},
			{Id: "displayName", Value: "Transmit Bandwidth"},
		}).
		OverrideByName("Value #C", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.PacketsPerSecond},
			{Id: "displayName", Value: "Receive Rate"},
		}).
		OverrideByName("Value #D", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.PacketsPerSecond},
			{Id: "displayName", Value: "Transmit Rate"},
		}).
		OverrideByName("Value #E", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.PacketsPerSecond},
			{Id: "displayName", Value: "Receive Drop Rate"},
		}).
		OverrideByName("Value #F", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.PacketsPerSecond},
			{Id: "displayName", Value: "Transmit Drop Rate"},
		})
}

func NewWorkloadNetworkBandwidth() *table.PanelBuilder {
	return pkgpanel.NewTable("Network Bandwidth by pod", pkgprometheus.Datasource()).
		NoValue("-").
		Targets([]cog.Builder[variants.Dataquery]{
			pkgprometheus.NewDataQuery(`
        (
				sum(irate(container_network_receive_bytes_total{namespace=~"$namespace"}[$__rate_interval]) * on (namespace,pod)
        group_left(workload,workload_type)
        namespace_workload_pod:kube_pod_owner:relabel{namespace=~"$namespace", workload=~"$service"}
				) by (pod))
			`),
			pkgprometheus.NewDataQuery(`
	      (
					sum(irate(container_network_transmit_bytes_total{namespace=~"$namespace"}[$__rate_interval]) * on (namespace,pod)
	        group_left(workload,workload_type)
	        namespace_workload_pod:kube_pod_owner:relabel{namespace=~"$namespace", workload=~"$service"}
				) by (pod))
			`),
		}).
		Transformations([]dashboard.DataTransformerConfig{
			{
				Id:    "timeSeriesTable",
				Topic: util.ToPtr(dashboard.DataTransformerConfigTopicSeries),
				Options: util.MustYamlToMap(`
				A: { timeField: Time }
        B: { timeField: Time }
			`),
			},
			{
				Id: "merge",
			},
		}).
		OverrideByName("Time", []dashboard.DynamicConfigValue{
			{Id: "custom.hideFrom.viz", Value: true},
		}).
		OverrideByName("pod", []dashboard.DynamicConfigValue{
			{Id: "displayName", Value: "Pod"},
		}).
		OverrideByName("Trend #A", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.BitsPerSecondSI},
			{Id: "displayName", Value: "Received"},
			{Id: "custom.cellOptions", Value: util.MustJSONToMap(`{"type":"sparkline","stat":"mean"}`)},
		}).
		OverrideByName("Trend #B", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.BitsPerSecondSI},
			{Id: "displayName", Value: "Transmitted"},
			{Id: "custom.cellOptions", Value: util.MustJSONToMap(`{"type":"sparkline","stat":"mean"}`)},
		})
}

func NewWorkloadNetworkPacketsRate() *table.PanelBuilder {
	return pkgpanel.NewTable("Network Packets Rate by pod", pkgprometheus.Datasource()).
		WithTarget(
			pkgprometheus.NewDataQuery(`
        (
					sum(irate(container_network_receive_packets_total{namespace=~"$namespace"}[$__rate_interval]) * on (namespace,pod)
					group_left(workload,workload_type)
          namespace_workload_pod:kube_pod_owner:relabel{namespace=~"$namespace", workload=~"$service"}
				) by (pod))
			`),
		).
		WithTarget(
			pkgprometheus.NewDataQuery(`
        (
					sum(irate(container_network_transmit_packets_total{namespace=~"$namespace"}[$__rate_interval]) * on (namespace,pod)
          group_left(workload,workload_type)
					namespace_workload_pod:kube_pod_owner:relabel{namespace=~"$namespace", workload=~"$service"}
				) by (pod))
			`),
		).
		NoValue("-").
		WithTransformation(dashboard.DataTransformerConfig{
			Id: "timeSeriesTable",
			Options: util.MustYamlToMap(`
				A: { timeField: Time }
        B: { timeField: Time }
			`),
		}).
		WithTransformation(dashboard.DataTransformerConfig{
			Id: "merge",
		}).
		OverrideByName("Time", []dashboard.DynamicConfigValue{
			{Id: "custom.hideFrom.viz", Value: true},
		}).
		OverrideByName("pod", []dashboard.DynamicConfigValue{
			{Id: "displayName", Value: "Pod"},
		}).
		OverrideByName("Trend #A", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.PacketsPerSecond},
			{Id: "displayName", Value: "Received"},
			{Id: "custom.cellOptions", Value: util.MustJSONToMap(`{"type":"sparkline","stat":"mean"}`)},
		}).
		OverrideByName("Trend #B", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.PacketsPerSecond},
			{Id: "displayName", Value: "Transmitted"},
			{Id: "custom.cellOptions", Value: util.MustJSONToMap(`{"type":"sparkline","stat":"mean"}`)},
		})
}

func NewWorkloadNetworkPacketsDroppedRate() *table.PanelBuilder {
	return pkgpanel.NewTable("Network Packets Dropped Rate pod", pkgprometheus.Datasource()).
		WithTarget(
			pkgprometheus.NewDataQuery(`
        (
					sum(irate(container_network_receive_packets_dropped_total{namespace=~"$namespace"}[$__rate_interval]) * on (namespace,pod)
					group_left(workload,workload_type)
					namespace_workload_pod:kube_pod_owner:relabel{namespace=~"$namespace", workload=~"$service"}
				) by (pod))
			`),
		).
		WithTarget(
			pkgprometheus.NewDataQuery(`
        (
					sum(irate(container_network_transmit_packets_dropped_total{namespace=~"$namespace"}[$__rate_interval]) * on (namespace,pod)
          group_left(workload,workload_type)
	        namespace_workload_pod:kube_pod_owner:relabel{namespace=~"$namespace", workload=~"$service"}
				) by (pod))
			`),
		).
		NoValue("-").
		WithTransformation(dashboard.DataTransformerConfig{
			Id: "timeSeriesTable",
			Options: util.MustYamlToMap(`
				A: { timeField: Time }
        B: { timeField: Time }
			`),
		}).
		WithTransformation(dashboard.DataTransformerConfig{
			Id: "merge",
		}).
		OverrideByName("Time", []dashboard.DynamicConfigValue{
			{Id: "custom.hideFrom.viz", Value: true},
		}).
		OverrideByName("pod", []dashboard.DynamicConfigValue{
			{Id: "displayName", Value: "Pod"},
		}).
		OverrideByName("Trend #A", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.PacketsPerSecond},
			{Id: "displayName", Value: "Received"},
			{Id: "custom.cellOptions", Value: util.MustJSONToMap(`{"type":"sparkline","stat":"mean"}`)},
		}).
		OverrideByName("Trend #B", []dashboard.DynamicConfigValue{
			{Id: "unit", Value: units.PacketsPerSecond},
			{Id: "displayName", Value: "Transmitted"},
			{Id: "custom.cellOptions", Value: util.MustJSONToMap(`{"type":"sparkline","stat":"mean"}`)},
		})
}
