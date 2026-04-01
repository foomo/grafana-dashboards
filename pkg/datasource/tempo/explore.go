package tempo

import (
	"github.com/foomo/grafana-dashboards/pkg/explore"
	"github.com/foomo/grafana-dashboards/pkg/util"
	"github.com/grafana/grafana-foundation-sdk/go/tempo"
)

func MustGetExploreURL(filters ...tempo.TraceqlFilter) string {
	ret, err := GetExploreURL(filters...)
	if err != nil {
		panic(err)
	}
	return ret
}

func GetExploreURL(filters ...tempo.TraceqlFilter) (string, error) {
	datasource := Datasource()
	return explore.NewLink(datasource, &tempo.TempoQuery{
		RefId:            "A",
		QueryType:        util.ToPtr("traceqlSearch"),
		Limit:            util.ToPtr[int64](20),
		TableType:        util.ToPtr(tempo.SearchTableTypeTraces),
		MetricsQueryType: util.ToPtr(tempo.MetricsQueryTypeRange),
		Filters:          filters,
		Datasource:       util.ToPtr(datasource),
	}).URL()
}
