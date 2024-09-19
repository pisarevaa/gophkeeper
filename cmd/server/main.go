package main

import (
	"context"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/pisarevaa/gophkeeper/internal/server/config"
	"github.com/pisarevaa/gophkeeper/internal/server/handler"
	"github.com/pisarevaa/gophkeeper/internal/server/logger"
	"github.com/pisarevaa/gophkeeper/internal/server/router"
	"github.com/pisarevaa/gophkeeper/internal/server/service/auth"
	"github.com/pisarevaa/gophkeeper/internal/server/service/keeper"
	"github.com/pisarevaa/gophkeeper/internal/server/storage/db"
	"github.com/pisarevaa/gophkeeper/internal/server/storage/minio"
	sharedUtils "github.com/pisarevaa/gophkeeper/internal/shared/utils"
)

const readTimeout = 5
const writeTimeout = 10
const shutdownTimeout = 10

// @title		Swagger Gophkeeper API
// @version	1.0
// @host		localhost:8080

func main() {
	ctxCancel, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctxStop, stop := signal.NotifyContext(ctxCancel, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()

	logger.NewLogger()
	config := config.NewConfig()
	validator := sharedUtils.NewValidator()
	repo, err := db.NewDB(config.Database.Dsn)
	if err != nil {
		panic(err)
	}
	defer db.CloseConnection(repo)
	s3, err := minio.NewMinio(config.Minio)
	if err != nil {
		panic(err)
	}

	authService := auth.NewService(
		auth.WithConfig(config),
		auth.WithStorage(repo),
	)

	keeperService := keeper.NewService(
		keeper.WithConfig(config),
		keeper.WithStorage(repo),
		keeper.WithMinio(s3),
	)

	handlers := handler.NewHandler(
		handler.WithConfig(config),
		handler.WithValidator(validator),
		handler.WithAuthService(authService),
		handler.WithKeeperService(keeperService),
	)
	srv := &http.Server{
		Addr:         config.Host,
		Handler:      router.NewRouter(handlers),
		ReadTimeout:  readTimeout * time.Second,
		WriteTimeout: writeTimeout * time.Second,
	}
	go func() {
		if errServer := srv.ListenAndServe(); errServer != nil && errServer != http.ErrServerClosed {
			slog.Info("Could not listen on " + config.Host)
		}
	}()
	slog.Info("Server is running on " + config.Host)
	<-ctxStop.Done()
	shutdownCtx, timeout := context.WithTimeout(ctxStop, shutdownTimeout*time.Second)
	defer timeout()
	err = srv.Shutdown(shutdownCtx)
	if err != nil {
		slog.Error(err.Error())
	}
	slog.Info("Server is gracefully shutdown")
}
