package flag

import "fmt"

// ValidationError provides the ability to create constant errors for run.Group
// validation errors, e.g. incorrect flag values.
type ValidationError string

// Error implements the built-in error interface.
func (v ValidationError) Error() string { return string(v) }

// NewValidationError provides a convenient helper function to create flag
// validation errors usable by run.Config implementations.
func NewValidationError(flag string, reason error) error {
	return fmt.Errorf(FlagErr, flag, reason)
}

const (
	// FlagErr can be used as formatting string for flag related validation
	// errors where the first variable lists the flag name and the second
	// variable is the actual error.
	FlagErr = "--%s error: %w"

	// ErrRequired is returned when required config options are not provided.
	ErrRequired ValidationError = "required"

	// ErrInvalidPath is returned when a path config option is invalid.
	ErrInvalidPath ValidationError = "invalid path"

	// ErrInvalidVal is returned when the value passed into a flag argument is invalid.
	ErrInvalidVal ValidationError = "invalid value"
)
