package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// EnvPrefix used as a prefix for environment variables,
// defults to empty.
var EnvPrefix = ""

// LoadEnvConfig loads the environment variables into the provided
// struct, using envconfig.
func LoadEnvConfig(t interface{}) {
	if err := envconfig.Process(EnvPrefix, t); err != nil {
		log.Fatalf("config: Unable to load config for %T: %s", t, err)
	}
}
