package manifest

import (
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/resource"
)

func Dashboard(dash dashboard.Dashboard, folderUID string) resource.Manifest {
	return resource.Manifest{
		ApiVersion: "dashboard.grafana.app/v1beta1",
		Kind:       "Dashboard",
		Metadata: resource.Metadata{
			Annotations: map[string]string{
				"grafana.app/folder": folderUID,
			},
			Name: *dash.Uid,
		},
		Spec: dash,
	}
}
