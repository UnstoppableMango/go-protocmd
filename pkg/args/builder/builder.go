package builder

import "fmt"

var empty = []any{}

type Builder struct {
	cmd []any
}

func New() *Builder {
	return &Builder{}
}

func (b *Builder) AppendAll(values []any) {
	b.cmd = append(b.cmd, values...)
}

func (b *Builder) Build() (args []string) {
	for _, x := range b.cmd {
		args = append(args, fmt.Sprint(x))
	}
	return
}

// func (b *Builder) AppendIf(pred bool, values ...any) {
// 	b.AppendAll(AllIf(pred, values))
// }

// func (b *Builder) AppendAllIf(pred bool, values []any) {
// 	b.AppendAll(AllIf(pred, values))
// }

// func (b *Builder) AppendMap(option string, values map[string]any) {
// 	b.AppendAll(OptMap(option, values))
// }

// func (b *Builder) AppendArg(name string, pred func() bool, value func() any) {
// 	b.AppendAll(Arg(name, pred, value))
// }

// func (b *Builder) AppendOpt(name string, has, get func() bool) {
// 	b.AppendAll(Opt(name, has, get))
// }

func If[T any](pred bool, values ...T) []T {
	if pred {
		return values
	}
	return []T{}
}

func AllIf[T any](pred bool, values []T) []T {
	if pred {
		return values
	}
	return []T{}
}

func OptMap[T any](option string, values map[string]T) (args []string) {
	for name, value := range values {
		args = append(args, option, name, fmt.Sprint(value))
	}
	return
}

func Arg[T any](name string, pred func() bool, value func() T) []string {
	return If(pred(), name, fmt.Sprint(value()))
}

func Opt(name string, has, get func() bool) []string {
	return If(has() && get(), name)
}
