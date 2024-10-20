package tracer

import "go.opentelemetry.io/otel"

func init() {
	tracer := otel.Tracer("tracer")

}
