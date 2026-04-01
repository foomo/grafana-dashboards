package prometheus

import (
	dashboard2 "github.com/foomo/grafana-dashboards/pkg/common/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

// NewLabelValueVariable creates a variable that lists all label values for a given metric and label.
func NewLabelValueVariable(name, metric, label string) *dashboard.QueryVariableBuilder {
	return dashboard.NewQueryVariableBuilder(name).
		Datasource(Datasource()).
		Query(dashboard2.StringOrMapAsString(
			"label_values(" + metric + ", " + label + ")",
		)).
		Refresh(dashboard.VariableRefreshOnTimeRangeChanged).
		Sort(dashboard.VariableSortAlphabeticalAsc).
		IncludeAll(true).
		Multi(true)
}

func NewNamespaceVariable() *dashboard.QueryVariableBuilder {
	return NewLabelValueVariable("namespace", `kube_pod_info`, "namespace").Label("📂")
}

func NewServiceVariable() *dashboard.QueryVariableBuilder {
	return NewLabelValueVariable("service", `namespace_workload_pod:kube_pod_owner:relabel{namespace=~"$namespace"}`, "workload").Label("🛎️")
}

func NewPodVariable() *dashboard.QueryVariableBuilder {
	return NewLabelValueVariable("pod", `namespace_workload_pod:kube_pod_owner:relabel{namespace=~"$namespace", workload=~"$service"}`, "pod").Label("📦")
}

func NewPodsVariable() *dashboard.QueryVariableBuilder {
	return NewLabelValueVariable("pods", `namespace_workload_pod:kube_pod_owner:relabel{namespace=~"$namespace", workload=~"$service"}`, "pod").Label("📦")
}

func NewContainerVariable() *dashboard.QueryVariableBuilder {
	return NewLabelValueVariable("container", `kube_pod_container_info{namespace=~"$namespace", pod=~"$pod"}`, "container").Label("⊙")
}
