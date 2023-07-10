// Package config contains configuration for the service.
package config

import (
	"github.com/spf13/viper"

	"github.com/Darkren/getmark-home/pkg/db"
	"github.com/Darkren/getmark-home/pkg/service/auth"
)

// Config is the service configuration.
type Config struct {
	DB          db.PgSQLConfig
	AuthService auth.Config
}

// FromEnv build configuration based on env vars.
func FromEnv() (*Config, error) {
	viper.AutomaticEnv()

	return &Config{
		DB:          db.PgSQLConfigFromEnv(),
		AuthService: auth.ConfigFromEnv(),
	}, nil
}
