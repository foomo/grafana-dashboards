package dashboard

import (
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

func NewFixedFieldColor(v string) *dashboard.FieldColorBuilder {
	return dashboard.NewFieldColorBuilder().
		Mode(dashboard.FieldColorModeIdFixed).
		FixedColor(v)
}
