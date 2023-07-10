package auth

import (
	"github.com/spf13/viper"
)

// Config is the Authorization Service configuration.
type Config struct {
	URL string
}

// ConfigFromEnv constructs Config based on env vars.
func ConfigFromEnv() Config {
	url := viper.GetString("AUTH_SERVICE_URL")

	return Config{
		URL: url,
	}
}
