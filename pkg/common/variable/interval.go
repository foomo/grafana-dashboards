package variable

import (
	dashboard2 "github.com/foomo/grafana-dashboards/pkg/common/dashboard"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

func IntervalVariable() *dashboard.IntervalVariableBuilder {
	return dashboard.NewIntervalVariableBuilder("interval").
		Label("⏱︎").
		AllowCustomValue(false).
		Current(dashboard.VariableOption{Value: dashboard2.StringOrArrayOfStringAsString("2m")}).
		Values(dashboard2.StringOrMapAsString("1m,2m,3m,5m,10m,30m"))
}
