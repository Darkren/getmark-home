// Package db contains structures related to DB usage.
package db

import (
	"fmt"
	"github.com/spf13/viper"
)

// PgSQLConfig is the configuration for the PostgreSQL connection.
type PgSQLConfig struct {
	Host     string
	Port     int
	DBName   string
	User     string
	Password string
	SSLMode  string
}

// PgSQLConfigFromEnv constructs PgSQLConfig based on env vars.
func PgSQLConfigFromEnv() PgSQLConfig {
	host := viper.GetString("DB_HOST")
	port := viper.GetInt("DB_PORT")
	dbName := viper.GetString("DB_NAME")
	user := viper.GetString("DB_USER")
	password := viper.GetString("DB_PASSWORD")
	sslMode := viper.GetString("DB_SSL_MODE")

	return PgSQLConfig{
		Host:     host,
		Port:     port,
		DBName:   dbName,
		User:     user,
		Password: password,
		SSLMode:  sslMode,
	}
}

// ToDriverDSN formats PgSQLConfig as the DSN string.
func (c PgSQLConfig) ToDriverDSN() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s port=%d sslmode=%s host=%s", c.User, c.Password, c.DBName,
		c.Port, c.SSLMode, c.Host)
}
