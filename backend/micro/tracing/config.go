package tracing

import "bitbucket.org/disruptive-technologies/micro/config"

// Config for Google Cloud Platform.
type Config struct {
	// ProjectID defines which Google cloud project we trace to.
	// If this is not defined, no tracing will be performed.
	// The service needs a service account with trace permission on this project.
	ProjectID string `envconfig:"TRACING_GOOGLE_PROJECT_ID"`
	// TracingPeriod defines how large fraction of requests are traced
	//
	// 0 : No tracing
	//
	// 1 : Every request is traced
	//
	// n : Every nth request is traced
	TracingPeriod uint64 `envconfig:"TRACING_PERIOD" default:"0"`
}

// LoadConfigFromEnv returns an Config object populated
// from environment variables.
func LoadConfigFromEnv() *Config {
	var cfg Config
	config.LoadEnvConfig(&cfg)
	return &cfg
}
