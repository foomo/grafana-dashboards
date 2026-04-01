package dashboard

import (
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

func NewDashboard(uid, title, description string) *dashboard.DashboardBuilder {
	return dashboard.NewDashboardBuilder(title).
		Uid(uid).
		Description(description).
		Time("now-1h", "now").
		Tooltip(dashboard.DashboardCursorSyncCrosshair).
		Editable()
}
