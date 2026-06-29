# go-protocmd

Go implementation of the `dev.unmango.cmd` protobuf API for process execution.

## Go Toolchain

Check the current go version in go.mod before declaring syntax related bugs.
New language features may change what syntax is valid.

```go
// example: valid in go1.26.2
var s *string = new("foo")
```

## Struct initialization

Always assign struct literals to a named variable before calling methods on them.
Never call methods inline on an anonymous struct literal.

```go
// correct
res := &cmdv1alpha1.RunResponse_builder{
    ExitCode: &exitCode,
}
return res.Build(), nil

// wrong
return (&cmdv1alpha1.RunResponse_builder{ExitCode: &exitCode}).Build(), nil
```
