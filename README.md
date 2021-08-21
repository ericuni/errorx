# usage
```go
import "github.com/ericuni/errorx"
```

The ericuni/errorx provides an easy way to trace errors without losing the original error context.

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
  return errorx.Trace(err)
}
```

