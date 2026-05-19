package conv

import (
	"github.com/unstoppablemango/go-protocmd/pkg/conv/builder"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func Args(msg protoreflect.Message) ([]string, error) {
	b := builder.New()
	for fd, v := range msg.Range {
		protoField(b, fd, v)
	}
	return b.Build(), nil
}

func protoField(b Builder, fd protoreflect.FieldDescriptor, v protoreflect.Value) {
	name := fd.TextName()
	switch fd.Kind() {
	case protoreflect.BoolKind:
		boolOpt(b, name, v)
	case protoreflect.StringKind:
		stringArg(b, name, v)
	}
}

func boolOpt(b Builder, name string, v protoreflect.Value) {
	Opt(b, name, v.IsValid, v.Bool)
}

func stringArg(b Builder, name string, v protoreflect.Value) {
	Arg(b, name, v.IsValid, v.String)
}
