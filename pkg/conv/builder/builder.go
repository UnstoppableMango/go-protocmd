package builder

import "fmt"

type Builder struct {
	cmd []string
}

func New() *Builder {
	return &Builder{}
}

func (b *Builder) AppendAll(values []any) {
	for _, v := range values {
		b.cmd = append(b.cmd, fmt.Sprint(v))
	}
}

func (b *Builder) Build() []string {
	return b.cmd
}
