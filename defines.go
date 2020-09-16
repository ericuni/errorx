package errs

import "errors"

// pre defined errors
var (
	ErrNotFound         = errors.New("not found")
	ErrPermissionDenied = errors.New("persmission denied")
)
