package tempo

import (
	"time"

	"github.com/foomo/grafana-dashboards/pkg/util"
	"github.com/grafana/grafana-foundation-sdk/go/tempo"
)

func NewNamespaceTraceQLFilter(v string) tempo.TraceqlFilter {
	return NewTraceQLFilter("k8s-namespace-name", tempo.TraceqlSearchScopeResource, "k8s.namespace.name", "=", StringOrArrayOfStringAsArrayOfString(v), "string")
}

func NewServiceTraceQLFilter(v string) tempo.TraceqlFilter {
	return NewTraceQLFilter("service-name", tempo.TraceqlSearchScopeResource, "service.name", "=", StringOrArrayOfStringAsArrayOfString(v), "string")
}

func NewMinDurationTraceQLFilter(v time.Duration) tempo.TraceqlFilter {
	return NewTraceQLFilter("min-duration", tempo.TraceqlSearchScopeResource, "duration", ">", StringOrArrayOfStringAsArrayOfString(v.String()), "duration")
}

func NewMaxDurationTraceQLFilter(v time.Duration) tempo.TraceqlFilter {
	return NewTraceQLFilter("max-duration", tempo.TraceqlSearchScopeResource, "duration", "<", StringOrArrayOfStringAsArrayOfString(v.String()), "duration")
}

func NewStatusTraceQLFilter(v string) tempo.TraceqlFilter {
	return NewTraceQLFilter("status", tempo.TraceqlSearchScopeIntrinsic, "status", "=", StringOrArrayOfStringAsString(v), "keyword")
}

func NewTraceQLFilter(id string, scope tempo.TraceqlSearchScope, tag, operator string, value tempo.StringOrArrayOfString, valueType string) tempo.TraceqlFilter {
	return tempo.TraceqlFilter{
		Id:        id,
		Scope:     util.ToPtr(scope),
		Tag:       util.ToPtr(tag),
		Operator:  util.ToPtr(operator),
		Value:     util.ToPtr(value),
		ValueType: util.ToPtr(valueType),
	}
}
