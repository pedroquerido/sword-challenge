package listener

// Listener for listening to events published
type Listener interface {
	Listen() error
}
