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

type MyError struct {
	Code int32
	Msg  string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%d %s", e.Code, e.Msg)
}

func (s *ErrorSuite) TestMyError() {
	assert := s.Assert()
	// t := s.T()

	errNotFound := &MyError{
		Code: http.StatusNotFound,
		Msg:  "not found",
	}

	err := errorx.Trace(errNotFound)
	err = errorx.Tracef(err, "more")

	assert.True(errors.Is(err, errNotFound))

	var myErr *MyError
	assert.True(errors.As(err, &myErr))
}

func TestError(t *testing.T) {
	suite.Run(t, new(ErrorSuite))
}
