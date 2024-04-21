package errorx

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	err := New("oops")
	t.Log(err)
}

func TestTrace(t *testing.T) {
	assert := assert.New(t)

	t.Run("nil", func(t *testing.T) {
		var err error = nil
		assert.Nil(Trace(err))
		assert.Nil(Tracef(err, "something happened"))
	})

	t.Run("standard error", func(t *testing.T) {
		errNotFound := errors.New("not found")

		err := Tracef(errNotFound, "something happened")
		assert.True(errors.Is(err, errNotFound))
		t.Logf("\n%+v", err)

		got := errors.Unwrap(err)
		t.Log(got)

		errNone := errors.New("none")
		assert.False(errors.Is(err, errNone))
	})

	t.Run("this error", func(t *testing.T) {
		errNotFound := New("not found")

		err := Tracef(errNotFound, "something happened")
		t.Logf("\n%+v", err)
	})
}

func TestCause(t *testing.T) {
	assert := assert.New(t)

	tds := []struct {
		name string
		err  error
		want error
	}{
		{
			name: "nil",
			err:  nil,
			want: nil,
		},
		{
			name: "standard error",
			err:  errors.New("err1"),
			want: errors.New("err1"),
		},
		{
			name: "trace err",
			err:  Trace(errors.New("err1")),
			want: errors.New("err1"),
		},
		{
			name: "more trace err",
			err:  Trace(Trace(errors.New("err1"))),
			want: errors.New("err1"),
		},
		{
			name: "tracef err",
			err:  Tracef(errors.New("err1"), "description"),
			want: errors.New("err1"),
		},
	}

	for _, td := range tds {
		t.Run(td.name, func(t *testing.T) {
			got := Cause(td.err)
			assert.Equal(td.want, got)
		})
	}
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

func TestPointerAndValueError(t *testing.T) {
	assert := assert.New(t)

	t.Run("pointer", func(t *testing.T) {
		errNotFound := &PointerError{
			code: -1,
			msg:  "not found",
		}

		err := Trace(errNotFound)
		err = Tracef(err, "more")

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

		err := Trace(errNotFound)
		err = Tracef(err, "more")

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

func TestNotComparableError(t *testing.T) {
	assert := assert.New(t)

	t.Run("pointer", func(t *testing.T) {
		errNotFound := &PointerErrorWithMap{
			code:  -1,
			msg:   "not found",
			extra: map[string]string{"name": "xxx"},
		}

		err := Trace(errNotFound)
		err = Tracef(err, "more")

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

		err := Trace(errNotFound)
		err = Tracef(err, "more")

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
