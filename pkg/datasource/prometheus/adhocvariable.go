package prometheus

import (
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

// NewAdHocVariable creates a new adhoc variable builder for prometheus.
func NewAdHocVariable() *dashboard.AdHocVariableBuilder {
	return dashboard.NewAdHocVariableBuilder("filter").
		Hide(dashboard.VariableHideHideLabel).
		Datasource(Datasource()).
		AllowCustomValue(false)
}
