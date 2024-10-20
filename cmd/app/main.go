package main

import (
	"context"
	"github.com/BinaryArchaism/order-processor/internal/config"
	"github.com/BinaryArchaism/order-processor/internal/logger"
	_ "github.com/BinaryArchaism/order-processor/internal/logger"
	"github.com/BinaryArchaism/order-processor/internal/repository"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	cfg, err := config.Parse()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse config")
	}

	err = logger.InitLogger(ctx, cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to initialize logger")
	}

	fx.WithLogger(log.Logger)

	_, err = repository.Connect(ctx, cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	log.Info().Msgf("Service started")
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt)
	<-osSignals
	log.Info().Msg("Shutting down...")
	cancel()
	time.Sleep(5 * time.Second)
	log.Info().Msg("Service shut down")
}
