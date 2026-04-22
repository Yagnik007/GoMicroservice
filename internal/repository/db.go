package repository

import (
	"fmt"
	"log/slog"

	"github.com/myorg/myservice/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDatabase establishes a pool of connections to the database using GORM
func ConnectDatabase(cfg *config.AppConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode)

	slog.Info("Connecting to postgres database...")
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("Failed to connect to database", "err", err)
		return nil, err
	}

	slog.Info("Successfully connected to database")
	return db, nil
}
