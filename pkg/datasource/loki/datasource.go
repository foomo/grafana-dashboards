package loki

import (
	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

func Datasource() dashboard.DataSourceRef {
	return dashboard.DataSourceRef{
		Uid:  cog.ToPtr("loki"),
		Type: cog.ToPtr("loki"),
	}
}
