package errs

import "errors"

var (
	ErrTicketsNotFound = errors.New("tickets not found")
	ErrTicketNotFound  = errors.New("ticket not found")

	ErrTicketAlreadyExists = errors.New("ticket already exists")

	ErrUnsupportedDBType    = errors.New("unsupported db type")
	ErrUnsupportedCacheType = errors.New("unsupported cache type")

	ErrTicketDoesNotBelongToUser = errors.New("ticket does not belong to user")
)
