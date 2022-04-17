package errorx_test

import (
	"errors"
	"fmt"
	"net/http"
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
}

type PointerError struct {
	Code int32
	Msg  string
}

func (e *PointerError) Error() string {
	return fmt.Sprintf("%d %s", e.Code, e.Msg)
}

type ValueError struct {
	Code int32
	Msg  string
}

func (e ValueError) Error() string {
	return fmt.Sprintf("%d %s", e.Code, e.Msg)
}

func (s *ErrorSuite) TestMyError() {
	assert := s.Assert()
	t := s.T()

	t.Run("pointer", func(t *testing.T) {
		errNotFound := &PointerError{
			Code: http.StatusNotFound,
			Msg:  "not found",
		}

		err := errorx.Trace(errNotFound)
		err = errorx.Tracef(err, "more")

		assert.True(errors.Is(err, errNotFound))

		var myErr *PointerError
		assert.True(errors.As(err, &myErr))
	})

	t.Run("value", func(t *testing.T) {
		errNotFound := ValueError{
			Code: http.StatusNotFound,
			Msg:  "not found",
		}

		err := errorx.Trace(errNotFound)
		err = errorx.Tracef(err, "more")

		assert.True(errors.Is(err, errNotFound))

		var myErr ValueError
		assert.True(errors.As(err, &myErr))
	})
}

func TestError(t *testing.T) {
	suite.Run(t, new(ErrorSuite))
}
