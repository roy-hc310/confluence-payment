package infrastructure

import (
	"confluence-payment/core-internal/utils"
	"context"

	"github.com/rs/zerolog/log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
)

type OtelInfra struct {
	Tracer trace.Tracer
}

func NewOtelInfra() *OtelInfra {
	ctx := context.Background()

	exporter, err := otlptracegrpc.New(
		ctx, otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(utils.GlobalEnv.OtelGrpcExporter),
	)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	res := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName("confluence-payment"),
	)

	tp := sdktrace.NewTracerProvider(sdktrace.WithBatcher(exporter), sdktrace.WithResource(res))
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	tracer := otel.Tracer("confluence-payment-tracer")

	return &OtelInfra{
		Tracer: tracer,
	}
}
