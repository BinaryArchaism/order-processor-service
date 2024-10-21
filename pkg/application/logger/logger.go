package logger

import (
	"context"
	"os"
	"time"

	"github.com/BinaryArchaism/order-processor/pkg/application/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogger(_ context.Context, _ config.Config) error {
	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Logger()

	return nil
}
