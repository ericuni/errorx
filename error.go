package errorx

import (
	"errors"
	"fmt"
)

// New new error
func New(format string, args ...interface{}) error {
	return fmt.Errorf("%s %s", getLocation(), fmt.Sprintf(format, args...))
}

// Trace trace
func Trace(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s\n%w", getLocation(), err)
}

// Tracef trace
func Tracef(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s %s\n%w", getLocation(), fmt.Sprintf(format, args...), err)
}

func Cause(err error) error {
	if err == nil {
		return nil
	}

	for {
		next := errors.Unwrap(err)
		if next == nil {
			return err
		}

		err = next
	}
}
