package loki

import (
	"github.com/grafana/grafana-foundation-sdk/go/loki"
)

func NewDataQuery(expr string) *loki.DataqueryBuilder {
	return loki.NewDataqueryBuilder().Datasource(Datasource()).Expr(expr)
}
