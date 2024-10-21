package main

import (
	"context"

	"github.com/BinaryArchaism/order-processor/pkg/application"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx := context.Background()
	app, err := application.Init()
	if err != nil {
		log.Fatal().Err(err).Msg("application init error")
	}

	err = app.Start(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("application start error")
	}
}
