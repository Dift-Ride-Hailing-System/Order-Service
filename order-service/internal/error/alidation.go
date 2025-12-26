package errors

import "fmt"

type ValidationError struct {
	Field   string
	Message string
}

func (v *ValidationError) Error() string {
	return fmt.Sprintf("validation failed: %s - %s", v.Field, v.Message)
}
