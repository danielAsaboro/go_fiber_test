package main

import (
	"context"
	"fiber_sample/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)


func main() {
	// first we must setup new Instance of FIber
	app := fiber.New()

	// Initialize trace provider
	tp := InitTracerProvider()
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()
	// Set global tracer provider & text propagators
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	// Initialize tracer
	// then we must define our routes function
	routes.Routes(app)

	// create http connection for this api
	app.Listen(":3000")
}
