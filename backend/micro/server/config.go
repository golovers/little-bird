package server

import (
	"bitbucket.org/disruptive-technologies/micro/auth"
	"bitbucket.org/disruptive-technologies/micro/config"
	"bitbucket.org/disruptive-technologies/micro/health"
	"bitbucket.org/disruptive-technologies/micro/tracing"
	"google.golang.org/grpc"
)

// Config holds the configuration options for the server instance.
type Config struct {
	Address     string `envconfig:"ADDRESS"`
	TLSCertFile string `envconfig:"TLS_CERT_FILE"`
	TLSKeyFile  string `envconfig:"TLS_KEY_FILE"`
	// Internal address used for exposing metrics, health checks etc.
	InternalAddress string `envconfig:"INTERNAL_ADDRESS" default:":9102"`
	// Paths
	LivenessPath  string `envconfig:"LIVENESS_PATH" default:"/liveness"`
	ReadinessPath string `envconfig:"READINESS_PATH" default:"/readiness"`
	MetricsPath   string `envconfig:"METRICS_PATH" default:"/metrics"`

	TraceConfig *tracing.Config

	// Needs to be set manually
	Auth          auth.Authenticator
	HealthChecks  []health.CheckFunc
	ServerOptions []grpc.ServerOption
}

// LoadConfigFromEnv returns an Config object populated
// from environment variables.
func LoadConfigFromEnv() *Config {
	var cfg Config
	config.LoadEnvConfig(&cfg)
	return &cfg
}
