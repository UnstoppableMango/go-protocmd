package proto

import (
	"fmt"

	cliv1alpha1 "github.com/unstoppablemango/go-protocmd/gen/unmango/cli/v1alpha1"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
)

type Parser struct {
	b *cliv1alpha1.Utility_builder
}

func Parse(msg *anypb.Any) (*cliv1alpha1.Utility, error) {
	p := &Parser{&cliv1alpha1.Utility_builder{}}
	return p.b.Build(), nil
}

func (p *Parser) message(msg protoreflect.Message) {
	for fd, v := range msg.Range {
		p.field(fd, v)
	}
}

func (p *Parser) option(name string, has, get func() bool) {
	if has() && get() {
		opt := &cliv1alpha1.Option_builder{
			Name: &name,
		}
		p.b.Options = append(p.b.Options, opt.Build())
	}
}

func (p *Parser) operand(has func() bool, v any) {
	if has() {
		p.b.Operands = append(p.b.Operands, fmt.Sprint(v))
	}
}

func (p *Parser) field(fd protoreflect.FieldDescriptor, v protoreflect.Value) {
	switch fd.Kind() {
	case protoreflect.BoolKind:
		p.option(fd.TextName(), v.IsValid, v.Bool)
	case protoreflect.StringKind:
		p.operand(v.IsValid, v.String())
	case protoreflect.MessageKind:
		if fd.IsMap() {
			// TODO
		} else if fd.IsList() {
			// TODO
		} else {
			p.message(v.Message())
		}
	}
}
