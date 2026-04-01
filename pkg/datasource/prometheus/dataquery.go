package prometheus

import (
	"github.com/grafana/grafana-foundation-sdk/go/prometheus"
)

func NewDataQuery(expr string) *prometheus.DataqueryBuilder {
	return prometheus.NewDataqueryBuilder().Datasource(Datasource()).Expr(expr)
}
