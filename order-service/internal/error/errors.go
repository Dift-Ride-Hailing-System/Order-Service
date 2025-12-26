package errors

import "errors"

// -----------------------------
// Common Order Service Errors
// -----------------------------
var (
	ErrInvalidLocation  = errors.New("invalid location data")
	ErrDriverNotFound   = errors.New("driver not found")
	ErrOrderNotFound    = errors.New("order not found")
	ErrAlreadyCancelled = errors.New("order already cancelled or completed")
	ErrInternal         = errors.New("internal server error")
	ErrDuplicateEvent   = errors.New("duplicate event ignored")
)
