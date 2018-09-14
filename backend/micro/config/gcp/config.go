package gcp

import "gitlab.com/7chip/little-bird/backend/micro/config"

// Config for Google Cloud Platform.
type Config struct {
	ProjectID string `envconfig:"GOOGLE_PROJECT_ID"`
	Keyfile   string `envconfig:"GOOGLE_APPLICATION_CREDENTIALS"`
}

// LoadConfigFromEnv returns an Config object populated
// from environment variables.
func LoadConfigFromEnv() *Config {
	var cfg Config
	config.LoadEnvConfig(&cfg)
	return &cfg
}
