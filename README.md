# errs
```go
import "github.com/ericuni/errs"
```

The ericuni/errs provides an easy way to trace errors without losing the original error context.

The exported New functions are designed to replace the errors.New and fmt.Errorf functions both. The same underlying
error is there, but the package also records the location at which the error was created.

A primary use case for this library is to add extra context any time an error is returned from a function.
```go
if err := SomeFunc(); err != nil {
  return err
}
```
This instead becomes:
```go
if err := SomeFunc(); err != nil {
  return errs.Trace(err)
}
```

# difference with [juju/errors](https://github.com/juju/errors)
Before go1.13, I have used a lot juju/errors to trace error, it's so awesome.
But it is also heavy in comparison with the errors.

Since go1.13, with the new %w syntax, error trace is supported officially.
So I created ericuni/errs to simplify the use.

ericuni/errs is compatible with errors, so we can use errors.Is and errors.As. juju/errors has predefined many errors in
[errtypes.go](https://github.com/juju/errors/blob/master/errortypes.go) with the intention to cover most cases.
It's hard for users to define new errtypes.
