package server

// Server ...
type Server interface {
	Run() error
	Stop() error
}
