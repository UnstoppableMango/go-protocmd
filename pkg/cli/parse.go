package cli

import (
	"fmt"

	cliv1alpha1 "github.com/unstoppablemango/go-protocmd/gen/unmango/cli/v1alpha1"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
)

func Parse(req *cliv1alpha1.ParseRequest) (*cliv1alpha1.ParseResponse, error) {
	switch in := req.WhichInput(); in {
	case cliv1alpha1.ParseRequest_Proto_case:
		return ParseProto(req.GetProto(), req.GetConvention())
	default:
		return nil, fmt.Errorf("unsupported input: %s", in)
	}
}

func ParseProto(x *anypb.Any, c cliv1alpha1.Convention) (*cliv1alpha1.ParseResponse, error) {
	var err error

	b := &protoParser{}
	res := &cliv1alpha1.ParseResponse_builder{}
	if res.Utility, err = b.Parse(x.ProtoReflect()); err != nil {
		return nil, err
	}
	return res.Build(), nil
}

type protoParser struct {
	cliv1alpha1.Utility_builder
}

func (b *protoParser) Parse(msg protoreflect.Message) (*cliv1alpha1.Utility, error) {
	b.message(msg)
	return b.Build(), nil
}

func (b *protoParser) message(msg protoreflect.Message) {
	for fd, v := range msg.Range {
		b.field(fd, v)
	}
}

func (b *protoParser) option(name string, has, get func() bool) {
	if has() && get() {
		opt := &cliv1alpha1.Option_builder{
			// TODO
			Name: &name,
		}
		b.Options = append(b.Options, opt.Build())
	}
}

func (b *protoParser) operand(has func() bool, v any) {
	if has() {
		b.Operands = append(b.Operands, fmt.Sprint(v))
	}
}

func (b *protoParser) field(fd protoreflect.FieldDescriptor, v protoreflect.Value) {
	switch fd.Kind() {
	case protoreflect.BoolKind:
		b.option(fd.TextName(), v.IsValid, v.Bool)
	case protoreflect.StringKind:
		b.operand(v.IsValid, v.String())
	case protoreflect.MessageKind:
		if fd.IsMap() {
			// TODO
		} else if fd.IsList() {
			// TODO
		} else {
			b.message(v.Message())
		}
	}
}
