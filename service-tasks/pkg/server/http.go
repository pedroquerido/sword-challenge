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
	name          string
	listenAddress string
	server        *http.Server
}

// NewHTTPServer ...
func NewHTTPServer(name, address string, handler *http.Handler) Server {

	srv := &http.Server{
		Handler: *handler,
		Addr:    address,
	}

	return &httpServer{
		name:          name,
		listenAddress: srv.Addr,
		server:        srv,
	}
}

// Run ...
func (s *httpServer) Run() error {

	log.Printf("starting %v server - address %v\n", s.name, s.listenAddress)

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

	log.Printf("stopping %v server", s.name)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.server.Shutdown(ctx)
}
