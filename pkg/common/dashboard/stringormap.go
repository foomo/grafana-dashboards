package dashboard

import (
	"github.com/foomo/grafana-dashboards/pkg/util"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

func StringOrMapAsString(s string) dashboard.StringOrMap {
	return dashboard.StringOrMap{
		String: util.ToPtr(s),
		Map:    nil,
	}
}

func StringOrMapAsMap(s map[string]any) dashboard.StringOrMap {
	return dashboard.StringOrMap{
		Map: s,
	}
}
