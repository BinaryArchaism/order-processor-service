package application

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/BinaryArchaism/order-processor/internal/repository"
	"github.com/BinaryArchaism/order-processor/pkg/application/config"
	"github.com/BinaryArchaism/order-processor/pkg/application/logger"
	"github.com/rs/zerolog/log"
)

const shutdownTimeout = 5 * time.Second

type App struct {
}

func Init() (*App, error) {
	return &App{}, nil
}

func (a *App) Start(ctx context.Context) error {
	cfg, err := config.InitConfig(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse config")
	}

	err = logger.InitLogger(ctx, cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to initialize logger")
	}

	_, err = repository.Connect(ctx, cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	log.Info().Msgf("Service started")
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt)
	<-osSignals
	log.Info().Msg("Shutting down...")

	time.Sleep(shutdownTimeout)
	log.Info().Msg("Service shut down")
	return nil
}
