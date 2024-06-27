package errs

import "errors"

var (
	ErrFlightNotFound       = errors.New("flight not found")
	ErrFlightAlreadyExists  = errors.New("flight already exists")
	ErrAirportNotFound      = errors.New("airport not found")
	ErrAirportAlreadyExists = errors.New("airport already exists")
)
