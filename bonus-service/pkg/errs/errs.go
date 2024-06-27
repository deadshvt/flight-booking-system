package errs

import "errors"

var (
	ErrPrivilegeNotFound      = errors.New("privilege not found")
	ErrPrivilegeAlreadyExists = errors.New("privilege already exists")
	ErrHistoryNotFound        = errors.New("history not found")
	ErrHistoryAlreadyExists   = errors.New("history already exists")
)
