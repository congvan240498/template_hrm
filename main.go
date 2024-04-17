package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"hrm/api/http"
	"hrm/api/http/handler/auth"
	"hrm/api/http/handler/user"
	"hrm/config"
	"hrm/internal/service"
	"hrm/pkg/logger"
	"hrm/pkg/mongodb"
	"hrm/repository"

	_ "github.com/labstack/echo/v4"
)

func main() {
	// Initialize logger
	logger.InitLog("lp-admin")
	log := logger.GetLogger()
	log.Info().Msg("Start lp-admin services")

	// Error code Init

	// Load configuration
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Msgf("Load config failed! %s", err)
	}

	// Set up Echo
	e := echo.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Kill, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	database, err := mongodb.ConnectMongoDB(ctx, &config.MongoDBConfig)
	if err != nil {
		log.Fatal().Msgf("load database fail! %s", err)
	}

	userRepo := repository.NewUserRepository(database)
	authRepo := repository.NewAuthRepository(database)

	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(authRepo, userRepo)

	authHandler := auth.NewAuthHandler(authService)
	userHandler := user.NewUserHandler(userService)

	// Start HTTP server
	srv := http.NewHttpServe(config, authHandler, userHandler)
	srv.Start(e)

	// Handle graceful shutdown
	<-ctx.Done()
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()
	if err := e.Shutdown(shutdownCtx); err != nil {
		log.Fatal().Msgf("Force shutdown services")
	}
}
