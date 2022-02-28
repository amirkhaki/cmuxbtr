package config

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

var Cfg Config

type Config struct {
	DBPath   string  `env:"CB_DB_PATH"`
	WPURL    string  `env:"CB_WP_URL"`
	WPKey    string  `env:"CB_WP_KEY"`
	WPSecret string  `env:"CB_WP_SECRET"`
	AEDPrice float64 `env:"CB_AED_PRICE"`
	Tax      float64 `env:"CB_TAX_PERCENT"`
}

func ParseEnv(cfg *Config) error {
	if err := env.Parse(cfg); err != nil {
		return fmt.Errorf("Could not parse config from env: %w", err)
	}
	return nil
}

/*
func New() (cfg *Config, err error) {
	cfg = &Config{}
	err = ParseEnv(cfg)
	return
}
*/

func Parse() error {
	cfg := Config{}
	err := ParseEnv(&cfg)
	Cfg = cfg
	return err
}
