package common

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

//LoadCfgFromEnv load configuration from environment
func LoadCfgFromEnv(t interface{}) {
	if err := envconfig.Process("", t); err != nil {
		log.Fatalf("config: unable to load config for %T: %s", t, err)
	}
}
