package variable

import (
	dashboard2 "github.com/foomo/grafana-dashboards/pkg/common/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

func Constant(name, value string) *dashboard.ConstantVariableBuilder {
	return dashboard.NewConstantVariableBuilder(name).
		Value(dashboard2.StringOrMapAsString(value))
}
