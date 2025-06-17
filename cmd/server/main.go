package main

import (
	"log"
	"net/http"
	"weatherdredger/internal/server"
)

func main() {
	// Setup logging (you could replace log with zap, zerolog, etc.)
	log.Println("Starting WeatherDredger server...")

	// Init tracing (stubbed here — add OpenTelemetry/etc if needed)
	// initTracing()

	// Register all HTTP handlers
	server.RegisterHandlers()

	// Start server
	addr := ":8080"
	log.Printf("Listening on http://localhost%s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
