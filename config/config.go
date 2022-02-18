package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	DBPath string `env:"CB_DB_PATH"`
}


func ParseEnv(cfg *Config) error {
	if err := env.Parse(cfg); err != nil {
		return fmt.Errorf("Could not parse config from env: %w", err)
	}
	return nil
}


func New() (cfg *Config, err error) {
	cfg = new(Config)
	err = ParseEnv(cfg)
	return
}
