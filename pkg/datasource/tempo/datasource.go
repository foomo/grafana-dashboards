package tempo

import (
	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

func Datasource() dashboard.DataSourceRef {
	return dashboard.DataSourceRef{
		Uid:  cog.ToPtr("tempo"),
		Type: cog.ToPtr("tempo"),
	}
}
