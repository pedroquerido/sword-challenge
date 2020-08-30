package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	runError = "error starting server: %w"
)

type httpServer struct {
	name    string
	address string
	server  *http.Server
}

// NewHTTPServer ...
func NewHTTPServer(name, address string, handler http.Handler) Server {

	return &httpServer{
		name:    name,
		address: address,
		server: &http.Server{
			Handler: handler,
			Addr:    address,
		},
	}
}

// Run ...
func (s *httpServer) Run() error {

	log.Printf("starting %v server - address %v\n", s.name, s.address)

	err := s.server.ListenAndServe()

	switch err {
	case nil, http.ErrServerClosed:
		return nil
	default:
		return fmt.Errorf(runError, err)
	}
}

// Stop ...
func (s *httpServer) Stop() error {

	log.Printf("stopping %v server\n", s.name)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.server.Shutdown(ctx)
}
