package errs

import "errors"

var (
	ErrTicketNotFound      = errors.New("ticket not found")
	ErrTicketAlreadyExists = errors.New("ticket already exists")
)
