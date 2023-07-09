package db

import "fmt"

type PgSQLConfig struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     int    `mapstructure:"DB_PORT"`
	DBName   string `mapstructure:"DB_NAME"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	SSLMode  string `mapstructure:"DB_SSL_MODE"`
}

func (c PgSQLConfig) ToDriverDSN() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s port=%d sslmode=%s host=%s", c.User, c.Password, c.DBName,
		c.Port, c.SSLMode, c.Host)
}
