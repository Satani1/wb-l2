package main

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"http-calendar/api"
	"http-calendar/internal/db"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	//logger init
	logger, err := zap.NewProduction()
	if err != nil {
		os.Exit(1)
	}
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

	//read environmental variables
	logger.Info("Reading config")
	cfg, err := NewConfig()
	if err != nil {
		logger.Error("Can't decode config", zap.Error(err))
		return
	}

	//connect to db
	logger.Info("Connection to DB")
	dbURL := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresDB)
	repo, err := db.NewPostgres(dbURL)
	if err != nil {
		logger.Error("Can't connect to DB", zap.Error(err))
	}
	defer repo.Close()

	App := api.NewApplication(logger, repo)
	Router := App.NewRouter()

	srv := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: Router,
	}

	logger.Info("Running HTTP server", zap.String("server address", cfg.ServerAddress))
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("Can't run the server", zap.Error(err), zap.String("server address", cfg.ServerAddress))
		}
	}()

	//graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10*time.Second))
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Can't shutdown the server", zap.Error(err))
	}

	logger.Info("Shutdown the server")

}
