package main

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/pisarevaa/fastlog/internal/config"
	"github.com/pisarevaa/fastlog/internal/handler"
	"github.com/pisarevaa/fastlog/internal/logger"
	"github.com/pisarevaa/fastlog/internal/producer"
	"github.com/pisarevaa/fastlog/internal/router"
	"github.com/pisarevaa/fastlog/internal/storage"
	"github.com/pisarevaa/fastlog/internal/utils"
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
	repo := storage.NewDB(config.Database.Dsn, logger)
	defer repo.CloseConnection()
	producer := producer.NewProducer(config.Kafka.Server, logger)
	defer producer.CloseConnection()

	handlers := handler.NewHandler(
		handler.WithConfig(config),
		handler.WithLogger(logger),
		handler.WithStorage(repo),
		handler.WithValidator(validator),
		handler.WithProducer(producer),
	)
	srv := &http.Server{
		Addr:         config.Host,
		Handler:      router.NewRouter(handlers),
		ReadTimeout:  readTimeout * time.Second,
		WriteTimeout: writeTimeout * time.Second,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Info("Could not listen on ", config.Host)
		}
	}()
	logger.Info("Server is running on ", config.Host)
	<-ctxStop.Done()
	shutdownCtx, timeout := context.WithTimeout(ctxStop, shutdownTimeout*time.Second)
	defer timeout()
	err := srv.Shutdown(shutdownCtx)
	if err != nil {
		logger.Error(err)
	}
	logger.Info("Server is gracefully shutdown")
}
