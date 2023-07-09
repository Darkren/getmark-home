package config

import (
	"github.com/Darkren/getmark-home/pkg/db"
	"github.com/Darkren/getmark-home/pkg/service/auth"
	"github.com/spf13/viper"
)

type Config struct {
	DB          db.PgSQLConfig
	AuthService auth.Config
}

func FromEnv() (*Config, error) {
	viper.AutomaticEnv()

	return &Config{
		DB:          db.PgSQLConfigFromEnv(),
		AuthService: auth.ConfigFromEnv(),
	}, nil
}
