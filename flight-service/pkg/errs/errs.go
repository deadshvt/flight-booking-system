package errs

import "errors"

var (
	ErrFlightsNotFound = errors.New("flights not found")
	ErrFlightNotFound  = errors.New("flight not found")
	ErrAirportNotFound = errors.New("airport not found")

	ErrFlightAlreadyExists  = errors.New("flight already exists")
	ErrAirportAlreadyExists = errors.New("airport already exists")

	ErrUnsupportedDBType    = errors.New("unsupported db type")
	ErrUnsupportedCacheType = errors.New("unsupported cache type")

	ErrInvalidPage = errors.New("invalid page")
	ErrInvalidSize = errors.New("invalid size")
)
