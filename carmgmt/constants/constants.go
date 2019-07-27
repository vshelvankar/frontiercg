package constants

import "errors"

var (
	// ErrNotFound is returned when the car is not found.
	ErrNotFound = errors.New("Car not found")
)
