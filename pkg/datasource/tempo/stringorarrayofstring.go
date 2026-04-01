package tempo

import (
	"github.com/foomo/grafana-dashboards/pkg/util"
	"github.com/grafana/grafana-foundation-sdk/go/tempo"
)

func StringOrArrayOfStringAsString(s string) tempo.StringOrArrayOfString {
	return tempo.StringOrArrayOfString{
		String: util.ToPtr(s),
	}
}

func StringOrArrayOfStringAsArrayOfString(s ...string) tempo.StringOrArrayOfString {
	return tempo.StringOrArrayOfString{
		ArrayOfString: s,
	}
}
