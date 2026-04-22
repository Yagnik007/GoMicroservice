package config

import (
	"log/slog"
	"strings"

	"github.com/spf13/viper"
)

// AppConfig represents the root configuration
type AppConfig struct {
	AppName string `mapstructure:"APP_NAME"`
	AppEnv  string `mapstructure:"APP_ENV"`
	AppPort string `mapstructure:"APP_PORT"`

	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBSSLMode  string `mapstructure:"DB_SSLMODE"`
}

var Cfg *AppConfig

// LoadConfig loads the configuration from .env file or environment variables
func LoadConfig() *AppConfig {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	// Replace dot with underscore in environment variables (e.g., DB.HOST -> DB_HOST)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		slog.Warn("No .env file found, relying on environment variables only", "err", err)
	}

	var config AppConfig
	if err := viper.Unmarshal(&config); err != nil {
		slog.Error("Failed to unmarshal config", "err", err)
		panic("Failed to load generic config")
	}

	Cfg = &config
	return Cfg
}
