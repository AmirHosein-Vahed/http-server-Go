package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"http-server-Go/handlers"
	"http-server-Go/server"
)

func main() {
	// Create a new server instance
	srv := server.NewServer(8080)

	// Register handlers
	srv.RegisterHandler("/hello", handlers.HelloHandler)
	srv.RegisterHandler("/health", handlers.HealthCheckHandler)

	// Channel to listen for errors coming from the server
	serverErrors := make(chan error, 1)

	// Start server in a goroutine
	go func() {
		log.Printf("Server is starting...")
		serverErrors <- srv.Start()
	}()

	// Channel to listen for interrupt signals
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Block until we receive a signal or an error
	select {
	case err := <-serverErrors:
		log.Fatalf("Error starting server: %v", err)

	case <-shutdown:
		log.Println("Starting shutdown...")

		// Give outstanding requests a deadline for completion
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Shutdown the server
		if err := srv.Stop(ctx); err != nil {
			log.Printf("Error shutting down server: %v", err)
			os.Exit(1)
		}
	}
}
