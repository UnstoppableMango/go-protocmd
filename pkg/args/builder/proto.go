package builder

import "google.golang.org/protobuf/reflect/protoreflect"

func ProtoField(fd protoreflect.FieldDescriptor, v protoreflect.Value) (args []string) {
	name := fd.TextName()

	switch {
	case fd.IsList():
		return ProtoList(name, v.List())
	case fd.IsMap():
		return ProtoMap(name, v.Map())
	}

	switch fd.Kind() {
	case protoreflect.BoolKind:
		return Opt(name, v.IsValid, v.Bool)
	case protoreflect.StringKind:
		return Arg(name, v.IsValid, v.String)
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind,
		protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return Arg(name, v.IsValid, v.Int)
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind,
		protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return Arg(name, v.IsValid, v.Uint)
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		return Arg(name, v.IsValid, v.Float)
	case protoreflect.EnumKind:
		return Arg(name, v.IsValid, v.Enum)
	case protoreflect.BytesKind:
		return Arg(name, v.IsValid, v.Bytes)
	case protoreflect.MessageKind, protoreflect.GroupKind:
		return ProtoMessage(v.Message())
	}

	panic("unsupported: " + fd.FullName())
}

func ProtoList(name string, l protoreflect.List) (args []string) {
	for i := range l.Len() {
		args = append(args, name, l.Get(i))
	}
	return
}

func ProtoMap(name string, m protoreflect.Map) (args []string) {
	for k, v := range m.Range {
		args = append(args, k, v)
	}
	return
}

func ProtoMessage(msg protoreflect.Message) (args []string) {
	for fd, v := range msg.Range {
		args = append(args, ProtoField(fd, v)...)
	}
	return
}
