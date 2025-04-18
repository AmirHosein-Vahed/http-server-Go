package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// Server represents our HTTP server
type Server struct {
	server   *http.Server
	port     int
	handlers map[string]http.HandlerFunc
	mu       sync.RWMutex
}

// NewServer creates a new instance of our custom server
func NewServer(port int) *Server {
	return &Server{
		port:     port,
		handlers: make(map[string]http.HandlerFunc),
	}
}

// RegisterHandler registers a new handler for a specific path
func (s *Server) RegisterHandler(path string, handler http.HandlerFunc) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers[path] = handler
}

// Start starts the HTTP server
func (s *Server) Start() error {
	mux := http.NewServeMux()

	// Register all handlers
	s.mu.RLock()
	for path, handler := range s.handlers {
		mux.HandleFunc(path, handler)
	}
	s.mu.RUnlock()

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: mux,
	}

	log.Printf("Server starting on port %d", s.port)
	return s.server.ListenAndServe()
}

// Stop gracefully stops the server
func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
