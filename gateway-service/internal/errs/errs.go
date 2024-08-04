package errs

import (
	"errors"
	"fmt"
)

var (
	ErrCreateConnection = errors.New("failed to create connection")
	ErrCloseConnection  = errors.New("failed to close connection")
)

func WrapError(wrap error, err error) error {
	if err == nil {
		return wrap
	}

	return fmt.Errorf("%s: %w", wrap.Error(), err)
}
