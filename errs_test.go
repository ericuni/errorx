package errs_test

import (
	"errors"
	"testing"

	"github.com/ericuni/errs"
	"github.com/stretchr/testify/assert"
)

var (
	ErrNotFound         = errors.New("not found")
	ErrPermissionDenied = errors.New("persmission denied")
)

func foo() error {
	return errs.Tracef(ErrNotFound, "user %v", "xxx")
}

func bar() error {
	return errs.Trace(foo())
}

func TestTrace_Function(t *testing.T) {
	t.Log(errs.Trace(bar()))
}

type Foo struct {
}

func (foo *Foo) sayHi() error {
	return errs.Trace(ErrPermissionDenied)
}

func TestTrace_Method(t *testing.T) {
	foo := &Foo{}
	t.Log(errs.Trace(foo.sayHi()))
}

func TestTraceNil(t *testing.T) {
	assert := assert.New(t)

	var err error = nil
	assert.Nil(errs.Trace(err))
	assert.Nil(errs.Tracef(err, "something happened"))
}
