package main

import (
	"log/slog"
	
	_ "github.com/myorg/myservice/docs"
	"github.com/myorg/myservice/internal/config"
	"github.com/myorg/myservice/internal/handlers"
	"github.com/myorg/myservice/internal/models"
	"github.com/myorg/myservice/internal/repository"
	"github.com/myorg/myservice/internal/server"
	"github.com/myorg/myservice/internal/services"
	"github.com/myorg/myservice/pkg/logger"
)

// @title           MyService API
// @version         1.0
// @description     This is a sample production-ready microservice.
// @host            localhost:8080
// @BasePath        /api/v1
func main() {
	// 1. Load configuration
	cfg := config.LoadConfig()

	// 2. Initialize Logger
	logger.InitLogger(cfg.AppEnv)
	slog.Info("Starting application...", "env", cfg.AppEnv)

	// 3. Connect to Database
	db, err := repository.ConnectDatabase(cfg)
	if err != nil {
		slog.Error("Database connection failed, exiting...")
		return // in production you might want to retry or panic based on policy
	}

	// 4. Run Auto Migration for Gorm models
	slog.Info("Running AutoMigrate...")
	if err := db.AutoMigrate(&models.Item{}); err != nil {
		slog.Error("AutoMigrate failed", "err", err)
	}

	// 5. Dependency Injection Setup
	itemRepo := repository.NewItemRepository(db)
	itemService := services.NewItemService(itemRepo)
	itemHandler := handlers.NewItemHandler(itemService)

	// 6. Initialize and start HTTP server
	srv := server.NewServer(cfg, itemHandler)
	srv.Start() // blocks until graceful shutdown
}
