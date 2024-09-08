package main

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/pisarevaa/gophkeeper/internal/server/config"
	"github.com/pisarevaa/gophkeeper/internal/server/handler"
	"github.com/pisarevaa/gophkeeper/internal/server/logger"
	"github.com/pisarevaa/gophkeeper/internal/server/router"
	"github.com/pisarevaa/gophkeeper/internal/server/storage"
	"github.com/pisarevaa/gophkeeper/internal/server/utils"
)

const readTimeout = 5
const writeTimeout = 10
const shutdownTimeout = 10

// @title		Swagger SDK Logger Service API
// @version	1.0
// @host		localhost:7000

func main() {
	ctxCancel, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctxStop, stop := signal.NotifyContext(ctxCancel, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()

	logger := logger.NewLogger()
	config := config.NewConfig(logger)
	validator := utils.NewValidator()
	repo, err := storage.NewDB(config.Database.Dsn, logger)
	if err != nil {
		panic(err)
	}
	defer repo.CloseConnection()

	handlers := handler.NewHandler(
		handler.WithConfig(config),
		handler.WithLogger(logger),
		handler.WithStorage(repo),
		handler.WithValidator(validator),
	)
	srv := &http.Server{
		Addr:         config.Host,
		Handler:      router.NewRouter(handlers),
		ReadTimeout:  readTimeout * time.Second,
		WriteTimeout: writeTimeout * time.Second,
	}
	go func() {
		if errServer := srv.ListenAndServe(); errServer != nil && errServer != http.ErrServerClosed {
			logger.Info("Could not listen on ", config.Host)
		}
	}()
	logger.Info("Server is running on ", config.Host)
	<-ctxStop.Done()
	shutdownCtx, timeout := context.WithTimeout(ctxStop, shutdownTimeout*time.Second)
	defer timeout()
	err = srv.Shutdown(shutdownCtx)
	if err != nil {
		logger.Error(err)
	}
	logger.Info("Server is gracefully shutdown")
}
