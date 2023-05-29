package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ScyllaHost       []string `envconfig:"SCYLLA_HOST"`
	ScyllaKeySpace   string   `envconfig:"SCYLLA_KEYSPACE"`
	ScyllaMigrations string   `envconfig:"SCYLLA_MIGRATIONS"`
}

func LoadEnv() (Config, error) {
	var env Config
	err := envconfig.Process("", &env)
	return env, err
}
