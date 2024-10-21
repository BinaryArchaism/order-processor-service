package tracer

import (
	"context"

	"github.com/BinaryArchaism/order-processor/pkg/application/config"
)

func InitTracer(_ context.Context, _ config.Config) error {
	// tracer := otel.Tracer("tracer")
	return nil
}
