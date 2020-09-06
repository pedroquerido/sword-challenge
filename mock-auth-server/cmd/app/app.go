package app

import (
	"log"
	"os"
	"os/signal"

	"github.com/pedroquerido/sword-challenge/mock-auth-server/internal/config"
	"github.com/pedroquerido/sword-challenge/mock-auth-server/internal/router"
	"github.com/pedroquerido/sword-challenge/mock-auth-server/pkg/server"
)

// Auth ...
type Auth struct {
}

// NewAuth ...
func NewAuth() *Auth {
	return &Auth{}
}

// Run ...
func (a *Auth) Run() error {

	// load configs
	cfg := config.Get()

	// Setup router
	router := router.New(cfg.Users.APIKeyUser, cfg.Users.UserRole)

	// Create and start HTTP server
	httpServer := server.NewHTTPServer("auth http", cfg.HTTP.Port, router.GetHTTPHandler())
	go httpServer.Run()

	// Wait for signal and cleanup afterwards
	waitForSignal()
	httpServer.Stop()

	// Wait for signal and cleanup afterwards
	waitForSignal()

	return nil
}

func waitForSignal() {

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	s := <-signalChannel

	log.Printf("received interrupt signal: %v", s)
}
