package config

import (
	"fmt"
	"github.com/Darkren/getmark-home/pkg/db"
	"github.com/Darkren/getmark-home/pkg/service/auth"
	"github.com/spf13/viper"
)

type Config struct {
	DB          db.PgSQLConfig
	AuthService auth.Config
}

func FromEnv() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("viper.ReadInConfig: %w", err)
	}

	var dbCfg db.PgSQLConfig
	err = viper.Unmarshal(&dbCfg)
	if err != nil {
		return nil, fmt.Errorf("viper.Unmarshal (DB Config): %w", err)
	}

	var authServiceCfg auth.Config
	err = viper.Unmarshal(&authServiceCfg)
	if err != nil {
		return nil, fmt.Errorf("viper.Unmarhsal (Auth Service Config): %w", err)
	}

	return &Config{
		DB:          dbCfg,
		AuthService: authServiceCfg,
	}, nil
}
