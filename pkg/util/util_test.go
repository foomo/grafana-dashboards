package util_test

import (
	"testing"

	"github.com/foomo/grafana-dashboards/pkg/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestYamlToMap(t *testing.T) {
	s := `
		elimiter: ","
		format: json
		jsonPaths:
		- alias: ""
		  path: squadron
		- alias: ""
		  path: user
		- alias: ""
		  path: commit
		- alias: ""
		  path: branch
		keepTime: false
		replace: false
		source: description
	`
	o, err := util.YamlToMap(s)
	require.NoError(t, err)
	assert.Len(t, o, 6)
}
