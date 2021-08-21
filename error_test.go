package errorx_test

import (
	"errors"
	"testing"

	"github.com/ericuni/errorx"
	"github.com/stretchr/testify/assert"
)

var (
	ErrNotFound         = errors.New("not found")
	ErrPermissionDenied = errors.New("persmission denied")
)

func foo() error {
	return errorx.Tracef(ErrNotFound, "user %v", "xxx")
}

func bar() error {
	return errorx.Trace(foo())
}

func TestTrace_Function(t *testing.T) {
	t.Log(errorx.Trace(bar()))
}

type Foo struct {
}

func (foo *Foo) sayHi() error {
	return errorx.Trace(ErrPermissionDenied)
}

func TestTrace_Method(t *testing.T) {
	foo := &Foo{}
	t.Log(errorx.Trace(foo.sayHi()))
}

func TestTraceNil(t *testing.T) {
	assert := assert.New(t)

	var err error = nil
	assert.Nil(errorx.Trace(err))
	assert.Nil(errorx.Tracef(err, "something happened"))
}
