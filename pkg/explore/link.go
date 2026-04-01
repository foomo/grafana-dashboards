package explore

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

type Link struct {
	SchemaVersion string          `json:"schemaVersion"`
	Panes         map[string]Pane `json:"panes"`
}

func NewLink(datasource dashboard.DataSourceRef, queries ...any) *Link {
	return &Link{
		SchemaVersion: "1",
		Panes: map[string]Pane{
			"20p": {
				Compact: false,
				Range: Range{
					From: "$__from",
					To:   "$__to",
				},
				Datasource: *datasource.Uid,
				Queries:    queries,
			},
		},
	}
}

func (l *Link) URL() (string, error) {
	v := &url.Values{}

	out, err := json.Marshal(l.Panes)
	if err != nil {
		return "", err
	}

	v.Set("panes", string(out))
	v.Set("schemaVersion", l.SchemaVersion)

	return "/explore?" + strings.ReplaceAll(v.Encode(), "%24", "$"), nil
}
