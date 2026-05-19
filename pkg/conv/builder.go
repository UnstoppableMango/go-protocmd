package conv

type Builder interface {
	AppendAll(values []any)
	Build() []string
}

func Append(b Builder, values ...any) {
	b.AppendAll(values)
}

func AppendIf(b Builder, pred bool, values ...any) {
	if pred {
		b.AppendAll(values)
	}
}

func AppendAllIf(b Builder, pred bool, values []any) {
	if pred {
		b.AppendAll(values)
	}
}

func AppendMap(b Builder, option string, values map[string]string) {
	for name, value := range values {
		Append(b, option, name, value)
	}
}

func Arg[T any](b Builder, name string, pred func() bool, value func() T) {
	AppendIf(b, pred(), name, value())
}

func Option(b Builder, pred bool, name string, opt func() bool) {
	AppendIf(b, pred && opt(), name)
}

func Opt(b Builder, name string, has, get func() bool) {
	AppendIf(b, has() && get(), name)
}

func AppendAll[S []T, T any](b Builder, values S) {
	for _, v := range values {
		Append(b, v)
	}
}

func AppendOpts[S []T, T any](b Builder, name string, get func() S) {
	for _, v := range get() {
		Append(b, name, v)
	}
}
