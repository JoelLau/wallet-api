package config

import "github.com/caarlos0/env/v11"

type Config struct {
	HTTPAddr    string `env:"HTTP_ADDR"`
	MirationDir string `env:"MIG_DIR"`
	PostgresDSN string `env:"PG_DSN"`
}

func FromEnv() (Config, error) {
	return env.ParseAs[Config]()
}
