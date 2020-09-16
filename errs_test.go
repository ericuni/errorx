package errs_test

import (
	"testing"

	"github.com/ericuni/errs"
	"github.com/stretchr/testify/assert"
)

func foo() error {
	return errs.Tracef(errs.ErrNotFound, "user %v", "xxx")
}

func bar() error {
	return errs.Trace(foo())
}

func TestTrace(t *testing.T) {
	t.Log(errs.Trace(bar()))
}

func TestTraceNil(t *testing.T) {
	assert := assert.New(t)

	var err error = nil
	assert.Nil(errs.Trace(err))
	assert.Nil(errs.Tracef(err, "something happened"))
}
