package app

import (
	"log"
	"os"
	"os/signal"
)

// Agent ...
type Agent struct {
}

// NewAgent ...
func NewAgent() *Agent {
	return &Agent{}
}

// Run ...
func (a *Agent) Run() error {

	// load configs
	//cfg := config.Get()

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
