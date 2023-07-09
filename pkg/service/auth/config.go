package auth

import (
	"github.com/spf13/viper"
)

type Config struct {
	URL string
}

func ConfigFromEnv() Config {
	url := viper.GetString("AUTH_SERVICE_URL")

	return Config{
		URL: url,
	}
}
