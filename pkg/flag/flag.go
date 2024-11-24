// Copyright (c) Bas van Beek 2024.
// Copyright (c) Tetrate, Inc 2021.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
