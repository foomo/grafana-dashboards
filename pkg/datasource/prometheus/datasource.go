package prometheus

import (
	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

// Datasource returns the datasource reference for prometheus.
func Datasource() dashboard.DataSourceRef {
	return dashboard.DataSourceRef{
		Uid:  cog.ToPtr("prometheus"),
		Type: cog.ToPtr("prometheus"),
	}
}
