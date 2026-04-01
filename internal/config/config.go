package config

import (
	"github.com/foomo/grafana-dashboards/pkg/api"
)

type Config struct {
	// Version of the config file.
	Version string `yaml:"version"`
	// Grafana API client configuration.
	Grafana api.ClientConfig `yaml:"grafana"`
}
