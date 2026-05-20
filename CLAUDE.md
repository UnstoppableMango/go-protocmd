# go-protocmd

Go implementation of the `dev.unmango.cmd` protobuf API for process execution.

## Struct initialization

Always assign struct literals to a named variable before calling methods on them. Never call methods inline on an anonymous struct literal.

```go
// correct
res := &cmdv1alpha1.RunResponse_builder{
    ExitCode: &exitCode,
}
return res.Build(), nil

// wrong
return (&cmdv1alpha1.RunResponse_builder{ExitCode: &exitCode}).Build(), nil
```
