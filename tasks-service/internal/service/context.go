package service

// Context represents the context structure to be used at the service package
type Context struct {
	UserID   string
	UserRole string
}

const (
	// ContextKey represents the key in which Context is expected at context.Context
	ContextKey = "service-context"
)
