package dashboard

import (
	"github.com/grafana/grafana-foundation-sdk/go/cog"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

type Variables []cog.Builder[dashboard.VariableModel]
