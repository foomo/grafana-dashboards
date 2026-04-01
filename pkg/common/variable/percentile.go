package variable

import (
	dashboard2 "github.com/foomo/grafana-dashboards/pkg/common/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

func Percentile() *dashboard.CustomVariableBuilder {
	return dashboard.NewCustomVariableBuilder("percentile").
		Label("P").
		Current(dashboard.VariableOption{Value: dashboard2.StringOrArrayOfStringAsString("0.95")}).
		AllowCustomValue(false).
		Values(dashboard2.StringOrMapAsString("99 : 0.99,95 : 0.95,90 : 0.90,75 : 0.75,50 : 0.50"))
}
