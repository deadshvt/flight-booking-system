package errs

import "errors"

var (
	ErrPrivilegesNotFound = errors.New("privileges not found")
	ErrPrivilegeNotFound  = errors.New("privilege not found")
	ErrHistoryNotFound    = errors.New("history not found")

	ErrPrivilegeAlreadyExists = errors.New("privilege already exists")
	ErrOperationAlreadyExists = errors.New("operation already exists")

	ErrUnsupportedDBType    = errors.New("unsupported db type")
	ErrUnsupportedCacheType = errors.New("unsupported cache type")
)
