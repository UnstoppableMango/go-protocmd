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
		Opt(b, name, v.IsValid, v.Bool)
	case protoreflect.StringKind:
		Arg(b, name, v.IsValid, v.String)
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind,
		protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		Arg(b, name, v.IsValid, v.Int)
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind,
		protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		Arg(b, name, v.IsValid, v.Uint)
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		Arg(b, name, v.IsValid, v.Float)
	case protoreflect.EnumKind:
		Arg(b, name, v.IsValid, v.Enum)
	case protoreflect.BytesKind:
		Arg(b, name, v.IsValid, v.Bytes)
	case protoreflect.MessageKind, protoreflect.GroupKind:
		msgArg(b, v.Message())
	}

	switch {
	case fd.IsList():
		listArg(b, name, v.List())
	case fd.IsMap():
		mapArg(b, name, v.Map())
	}
}

func listArg(b Builder, name string, l protoreflect.List) {
	for i := range l.Len() {
		Append(b, name, l.Get(i))
	}
}

func mapArg(b Builder, name string, m protoreflect.Map) {
	for k, v := range m.Range {
		Append(b, name, k, v)
	}
}

func msgArg(b Builder, msg protoreflect.Message) {
	for fd, v := range msg.Range {
		// How does protobuf handle recursion?
		// Do we need to a base case?
		protoField(b, fd, v)
	}
}
