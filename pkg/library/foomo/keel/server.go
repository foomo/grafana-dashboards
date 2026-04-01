package keel

import (
	pkgdashboard "github.com/foomo/grafana-dashboards/pkg/common/dashboard"
	pkgvariable "github.com/foomo/grafana-dashboards/pkg/common/variable"
	pkgprometheus "github.com/foomo/grafana-dashboards/pkg/datasource/prometheus"
	"github.com/foomo/grafana-dashboards/pkg/library/kubernetes/kubernetes"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

func NewServerDashboard(namespace, service string) *dashboard.DashboardBuilder {
	return pkgdashboard.NewDashboard(
		"foomo_keel_server",
		"Keel Server",
		"Foomo Keel Server",
	).
		Variables(pkgdashboard.Variables{
			pkgvariable.Percentile(),
			// pkgvariable.IntervalVariable(),
			pkgvariable.Constant("namespace", namespace),
			pkgvariable.Constant("service", service),
			pkgprometheus.NewPodsVariable().Hide(dashboard.VariableHideHideVariable),
			// pkgprometheus.NewPodVariable().IncludeAll(false).Multi(false),
			// pkgprometheus.NewContainerVariable().IncludeAll(false).Multi(false),
			pkgprometheus.NewAdHocVariable(),
		}).
		// WithPanel(otelhttplib.NewServerTotalRequests()).
		// WithPanel(otelhttplib.NewClientTotalRequests()).
		// WithPanel(foomogotsrpclib.NewServerTotalRequests()).
		WithPanel(kubernetes.NewNetworkIOByContainer().Span(24)).
		WithPanel(kubernetes.NewContainerInfos().Span(24)).
		// WithPanel(kubernetes.NewQuotasByContainer().Span(24)).
		// WithPanel(kubernetes.NewCPUQuotaByContainer().Span(24)).
		// WithPanel(kubernetes.NewMemoryQuotaByContainer().Span(24)).
		// WithPanel(kubernetes.NewCPUQuotaByPod().Span(24)).
		// WithPanel(kubernetes.NewMemoryQuotaByPod().Span(24)).
		// WithPanel(kubernetes.NewCPUUsageByPod().Span(12)).
		// WithPanel(kubernetes.NewCPUUsageByContainer().Span(12)).
		// WithPanel(kubernetes.NewMemoryUsageByPod().Span(12)).
		// WithPanel(kubernetes.NewMemoryUsageByContainer().Span(12)).
		// WithPanel(kubernetes.NewPodCPUResourcesByContainer().Span(24)).
		// WithPanel(kubernetes.NewPodResourcesByContainer().Span(24)).
		// WithPanel(kubernetes.NewPodMemoryResourcesByContainer().Span(24)).
		// WithPanel(kubernetes.NewWorkloadCPUQuota().Span(24)).
		// WithPanel(kubernetes.NewWorkloadMemoryQuota().Span(24)).
		// WithPanel(kubernetes.NewWorkloadNetwork().Span(24)).
		// WithPanel(kubernetes.NewWorkloadNetworkBandwidth().Span(24)).
		// WithPanel(kubernetes.NewWorkloadNetworkPacketsRate().Span(24)).
		// WithPanel(kubernetes.NewWorkloadNetworkPacketsDroppedRate().Span(24)).
		// WithPanel(nodelib.NewCPU()).
		// WithPanel(nodelib.NewMemory()).
		// WithPanel(nodelib.NewMemoryDeriv()).
		// WithPanel(nodelib.NewSystemMemory()).
		// WithPanel(nodelib.NewSystemMemoryDeriv()).
		// WithPanel(lokilib.NewTotalLogs()).
		// WithPanel(lokilib.NewTotalWarnings()).
		// WithPanel(lokilib.NewTotalErrors()).
		// WithPanel(tempolib.NewTotalSpans()).
		// WithPanel(tempolib.NewTotalErrorSpans()).
		// WithPanel(tempolib.NewTotalErrorSpans()).
		// WithPanel(prometheuslib.NewTotalMetrics()).
		// WithPanel(prometheuslib.NewTotalSeries()).
		// WithPanel(prometheuslib.NewTotalSeries()).
		// WithPanel(prometheuslib.NewSeriesByMetrics().Span(24)).
		// WithPanel(golanglib.NewGoRoutines().Span(12)).
		// WithPanel(golanglib.NewGoRoutinesDeriv().Span(12)).
		// WithPanel(golanglib.NewOpenFDS().Span(12)).
		// WithPanel(golanglib.NewOpenFDSDeriv().Span(12)).
		// WithPanel(golanglib.NewMemStats().Span(12)).
		// WithPanel(golanglib.NewMemStatsDeriv().Span(12)).
		// WithPanel(golanglib.NewMemory().Span(12)).
		// WithPanel(golanglib.NewMemoryDeriv().Span(12)).
		// WithPanel(golanglib.NewGCDurationQuantiles().Span(12)).
		Editable()
}
