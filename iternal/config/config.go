package config

import (
	"errors"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	PostgresDatabase  string `env:"POSTGRES_DB"`
	PostgresHost      string `env:"POSTGRES_HOST"`
	PostgresPort      int    `env:"POSTGRES_PORT"`
	PostgresUsername  string `env:"POSTGRES_USER"`
	PostgresPassword  string `env:"POSTGRES_PASSWORD"`
	HttpListenAddress string `env:"HTTP_LISTEN_ADDR"`
}

func GetConfig() (*Config, error) {
	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, errors.New("failed to parse config")
	}

	return cfg, nil
}
