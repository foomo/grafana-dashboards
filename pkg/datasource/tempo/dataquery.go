package tempo

import (
	"github.com/grafana/grafana-foundation-sdk/go/tempo"
)

func NewDataQuery(expr string) *tempo.TempoQueryBuilder {
	return tempo.NewTempoQueryBuilder().Datasource(Datasource()).Query(expr)
}
