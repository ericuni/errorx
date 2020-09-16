package errs

import (
	"fmt"
)

// New new error
func New(format string, args ...interface{}) error {
	return fmt.Errorf("%v %s", getLocation(2), fmt.Sprintf(format, args...))
}

// Trace trace
func Trace(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%v\n%w", getLocation(2), err)
}

// Tracef trace
func Tracef(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%v %s\n%w", getLocation(2), fmt.Sprintf(format, args...), err)
}
