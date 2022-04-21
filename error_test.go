package errorx_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/ericuni/errorx"
	"github.com/stretchr/testify/suite"
)

type ErrorSuite struct {
	suite.Suite
}

func (s *ErrorSuite) TestNil() {
	assert := s.Assert()
	// t := s.T()

	var err error = nil
	assert.Nil(errorx.Trace(err))
	assert.Nil(errorx.Tracef(err, "something happened"))
}

func (s *ErrorSuite) TestStandardError() {
	assert := s.Assert()
	// t := s.T()

	errNotFound := errors.New("not found")
	err := errorx.Tracef(errNotFound, "something happened")
	assert.True(errors.Is(err, errNotFound))

	errNone := errors.New("none")
	assert.False(errors.Is(err, errNone))
}

type PointerError struct {
	code int32
	msg  string
}

func (e *PointerError) Error() string {
	return fmt.Sprintf("%d %s", e.code, e.msg)
}

type ValueError struct {
	code int32
	msg  string
}

func (e ValueError) Error() string {
	return fmt.Sprintf("%d %s", e.code, e.msg)
}

func (s *ErrorSuite) TestMyError() {
	assert := s.Assert()
	t := s.T()

	t.Run("pointer", func(t *testing.T) {
		errNotFound := &PointerError{
			code: -1,
			msg:  "not found",
		}

		err := errorx.Trace(errNotFound)
		err = errorx.Tracef(err, "more")

		assert.True(errors.Is(err, errNotFound))

		var myErr *PointerError
		assert.True(errors.As(err, &myErr))

		errNone := &PointerError{
			code: -2,
			msg:  "none",
		}
		assert.False(errors.Is(err, errNone))
	})

	t.Run("value", func(t *testing.T) {
		errNotFound := ValueError{
			code: -1,
			msg:  "not found",
		}

		err := errorx.Trace(errNotFound)
		err = errorx.Tracef(err, "more")

		assert.True(errors.Is(err, errNotFound))

		var myErr ValueError
		assert.True(errors.As(err, &myErr))

		errNone := ValueError{
			code: -2,
		}
		assert.False(errors.Is(err, errNone))
	})
}

// https://go.dev/ref/spec#Comparison_operators
// Slice, map, and function values are not comparable.
type PointerErrorWithMap struct {
	code  int32
	msg   string
	extra map[string]string
}

func (e *PointerErrorWithMap) Error() string {
	return fmt.Sprintf("%d %s %v", e.code, e.msg, e.extra)
}

type ValueErrorWithMap struct {
	code  int32
	msg   string
	extra map[string]string
}

func (e ValueErrorWithMap) Error() string {
	return fmt.Sprintf("%d %s %v", e.code, e.msg, e.extra)
}

func (s *ErrorSuite) TestNotComparable() {
	assert := s.Assert()
	t := s.T()

	t.Run("pointer", func(t *testing.T) {
		errNotFound := &PointerErrorWithMap{
			code:  -1,
			msg:   "not found",
			extra: map[string]string{"name": "xxx"},
		}

		err := errorx.Trace(errNotFound)
		err = errorx.Tracef(err, "more")

		assert.True(errors.Is(err, errNotFound))

		var myErr *PointerErrorWithMap
		assert.True(errors.As(err, &myErr))

		errNone := &PointerErrorWithMap{
			code: -2,
			msg:  "none",
		}
		assert.False(errors.Is(err, errNone))
	})

	t.Run("value", func(t *testing.T) {
		errNotFound := ValueErrorWithMap{
			code:  -1,
			msg:   "not found",
			extra: map[string]string{"name": "xxx"},
		}

		err := errorx.Trace(errNotFound)
		err = errorx.Tracef(err, "more")

		// map is not comparable, but errors.Is would use the comparability
		assert.False(errors.Is(err, errNotFound))

		var myErr ValueErrorWithMap
		assert.True(errors.As(err, &myErr))

		errNone := ValueErrorWithMap{
			code: -2,
			msg:  "none",
		}
		assert.False(errors.Is(err, errNone))
	})
}

func TestError(t *testing.T) {
	suite.Run(t, new(ErrorSuite))
}
