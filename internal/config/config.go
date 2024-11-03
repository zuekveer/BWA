package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	DB     DB
	HTTP   HTTPConfig
	Logger Logger
}

type DB struct {
	URI string `env:"DB_URI" envDefault:"postgresql://postgres:password@localhost:5432/auth"`
}

type HTTPConfig struct {
	Port int `env:"HTTP_PORT" envDefault:"8080"`
}

type Logger struct {
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
}

func Parse() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("load config from env: %w", err)
	}
	return &cfg, nil
}
