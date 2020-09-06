package listener

// EventListener for listening to events published
type EventListener interface {
	Listen() error
}
