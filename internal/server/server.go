package server

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/myorg/myservice/internal/config"
	"github.com/myorg/myservice/internal/handlers"
	"github.com/myorg/myservice/internal/routes"
)

// Server holds the gin engine and configuration
type Server struct {
	engine      *gin.Engine
	config      *config.AppConfig
	itemHandler *handlers.ItemHandler
}

// NewServer factory
func NewServer(cfg *config.AppConfig, itemHandler *handlers.ItemHandler) *Server {
	
	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.Default()

	// Apply Middlewares (e.g. CORS, Logger, Recovery)
	engine.Use(gin.Recovery())

	return &Server{
		engine:      engine,
		config:      cfg,
		itemHandler: itemHandler,
	}
}

// Start runs the HTTP server and handles graceful shutdown
func (s *Server) Start() {
	routes.RegisterRoutes(s.engine, s.itemHandler)

	srv := &http.Server{
		Addr:    ":" + s.config.AppPort,
		Handler: s.engine,
	}

	go func() {
		slog.Info("Starting server", "port", s.config.AppPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("ListenAndServe err", "err", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("Server forced to shutdown", "err", err)
	}

	slog.Info("Server exiting")
}
