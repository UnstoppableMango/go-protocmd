package conv

import (
	"github.com/unstoppablemango/go-protocmd/pkg/conv/builder"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func ProtoArgs(msg protoreflect.Message) ([]string, error) {
	b := builder.New()
	for fd, v := range msg.Range {
		protoField(b, fd, v)
	}
	return b.Build(), nil
}

func protoField(b Builder, fd protoreflect.FieldDescriptor, v protoreflect.Value) {
	switch fd.Kind() {
	case protoreflect.BoolKind:
		boolOpt(b, fd.TextName(), v)
	}
}

func boolOpt(b Builder, name string, v protoreflect.Value) {
	Opt(b, name, v.IsValid, v.Bool)
}
