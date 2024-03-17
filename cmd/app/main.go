package main

import (
	"context"
	config "github.com/Max425/film-library.git/internal/comfig"
	"github.com/Max425/film-library.git/internal/http-server"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title WB API
// @version 1.0
// @description API Server for Film Library

// @host localhost:8000
// @BasePath /
func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

func run() error {
	// read config
	cfg := config.MustLoad()

	// init logger
	logger, err := initLogger()
	if err != nil {
		return err
	}

	// create http server with all handlers & services & repositories
	srv, err := http_server.NewHttpServer(logger, cfg.Postgres, cfg.HttpAddr)
	if err != nil {
		logger.Error("create http server", zap.Error(err))
		return err
	}

	// listen to OS signals and gracefully shutdown HTTP server
	stopped := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sigint
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err = srv.Shutdown(ctx); err != nil {
			logger.Error("HTTP Server Shutdown", zap.Error(err))
		}
		close(stopped)
	}()

	logger.Info("Starting HTTP server", zap.String("addr", cfg.HttpAddr))

	// start HTTP server
	if err = srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		logger.Error("HTTP server ListenAndServe", zap.Error(err))
	}

	<-stopped

	return nil
}

func initLogger() (*zap.Logger, error) {
	return zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Development:      true,
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}.Build()
}
