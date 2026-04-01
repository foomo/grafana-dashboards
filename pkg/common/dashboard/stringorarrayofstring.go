package dashboard

import (
	"github.com/foomo/grafana-dashboards/pkg/util"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

func StringOrArrayOfStringAsString(s string) dashboard.StringOrArrayOfString {
	return dashboard.StringOrArrayOfString{
		String: util.ToPtr(s),
	}
}

func StringOrArrayOfStringAsArrayOfString(s ...string) dashboard.StringOrArrayOfString {
	return dashboard.StringOrArrayOfString{
		ArrayOfString: s,
	}
}
