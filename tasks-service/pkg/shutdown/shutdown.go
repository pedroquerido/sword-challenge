package shutdown

import (
	"log"
	"os"
	"os/signal"
	"time"
)

const (
	receivedInterruptSignal = "received interrupt signal\n"
	cleanupResult           = "exiting: cleanup %s\n"
)

// Shutdown ...
type Shutdown struct {
	InitiateChannel chan struct{}
	DoneChannel     chan struct{}
	timeout         time.Duration
}

// NewShutdown ...
func NewShutdown(timeout time.Duration) *Shutdown {

	return &Shutdown{
		InitiateChannel: make(chan struct{}, 1),
		DoneChannel:     make(chan struct{}, 1),
		timeout:         timeout,
	}
}

// WaitForSignal ...
func (shutdown *Shutdown) WaitForSignal() {

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)
	<-signalChannel

	log.Printf(receivedInterruptSignal)
	shutdown.InitiateChannel <- struct{}{}
	timeout := time.After(shutdown.timeout)

	select {
	case <-shutdown.DoneChannel:
		log.Printf(cleanupResult, "successful")
	case <-timeout:
		log.Printf(cleanupResult, "timed out")
	case <-signalChannel:
		log.Fatalf(cleanupResult, "unsuccessful - forcing shutdown")
	}
}
