package api

import (
	"net/url"

	"github.com/go-openapi/strfmt"
	"github.com/grafana/grafana-openapi-client-go/client"
)

func NewClient(cfg ClientConfig) *client.GrafanaHTTPAPI {
	return client.NewHTTPClientWithConfig(strfmt.Default, &client.TransportConfig{
		// Host is the domain name or IP address of the host that serves the API.
		Host: cfg.Host,
		// BasePath is the URL prefix for all API paths, relative to the host root.
		BasePath: "/api",
		// Schemes are the transfer protocols used by the API (http or https).
		Schemes: []string{"http"},
		// BasicAuth is contains basic auth credentials.
		BasicAuth: url.UserPassword(cfg.Username, cfg.Password),
	})
}
