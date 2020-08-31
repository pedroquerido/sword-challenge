package service

import (
	"context"
	pkgError "tasks-service/pkg/error"
)

const (
	detailParsingContext  = "could not parse"
	detailMissingUserID   = "missing UserID"
	detailMissingUserRole = "missing UserRole"
)

var (
	// ContextKey represents the Context Key in which Context is expected at context.Context
	ContextKey contextKey = "service-context"
)

type contextKey string

// Context represents the context structure to be used at the service package
type Context struct {
	UserID   string
	UserRole string
}

func parseContext(ctx context.Context) (*Context, error) {

	context := ctx.Value(ContextKey)

	// Cast
	serviceContext, ok := context.(Context)
	if !ok {
		return nil, pkgError.NewError(ErrorMissingContext).WithDetails(detailParsingContext)
	}

	// Validate designed required fields - could later use validator as well
	details := make([]string, 0, 2)
	if serviceContext.UserID == "" {
		details = append(details, detailMissingUserID)
	}

	if serviceContext.UserRole == "" {
		details = append(details, detailMissingUserRole)
	}

	if len(details) > 0 {
		return nil, pkgError.NewError(ErrorMissingContext).WithDetails(details...)

	}

	return &serviceContext, nil
}
