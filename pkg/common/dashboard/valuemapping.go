package dashboard

import (
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

func ValueMappingAsValueMap(v map[string]dashboard.ValueMappingResult) dashboard.ValueMapping {
	return dashboard.ValueMapping{
		ValueMap: &dashboard.ValueMap{
			Type:    dashboard.MappingTypeValueToText,
			Options: v,
		},
	}
}
