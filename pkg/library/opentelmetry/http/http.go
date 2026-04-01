package http

import (
	pkgdashboard "github.com/foomo/grafana-dashboards/pkg/common/dashboard"
	"github.com/foomo/grafana-dashboards/pkg/util"
	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

func NewBucketValueMapping() dashboard.ValueMapping {
	return pkgdashboard.ValueMappingAsValueMap(map[string]dashboard.ValueMappingResult{
		"0.005": {
			Index: util.ToPtr[int32](0),
			Text:  util.ToPtr("5ms"),
		},
		"0.01": {
			Index: util.ToPtr[int32](1),
			Text:  util.ToPtr("10ms"),
		},
		"0.025": {
			Index: util.ToPtr[int32](2),
			Text:  util.ToPtr("25ms"),
		},
		"0.05": {
			Index: util.ToPtr[int32](3),
			Text:  util.ToPtr("50ms"),
		},
		"0.075": {
			Index: util.ToPtr[int32](4),
			Text:  util.ToPtr("75ms"),
		},
		"0.1": {
			Index: util.ToPtr[int32](5),
			Text:  util.ToPtr("100ms"),
		},
		"0.25": {
			Index: util.ToPtr[int32](6),
			Text:  util.ToPtr("250ms"),
		},
		"0.5": {
			Index: util.ToPtr[int32](7),
			Text:  util.ToPtr("500ms"),
		},
		"0.75": {
			Index: util.ToPtr[int32](8),
			Text:  util.ToPtr("750ms"),
		},
		"1.0": {
			Index: util.ToPtr[int32](9),
			Text:  util.ToPtr("1s"),
		},
		"2.5": {
			Index: util.ToPtr[int32](10),
			Text:  util.ToPtr("2.5s"),
		},
		"5.0": {
			Index: util.ToPtr[int32](11),
			Text:  util.ToPtr("5s"),
		},
		"7.5": {
			Index: util.ToPtr[int32](12),
			Text:  util.ToPtr("7.5s"),
		},
		"10.0": {
			Index: util.ToPtr[int32](13),
			Text:  util.ToPtr("10s"),
		},
	})
}
