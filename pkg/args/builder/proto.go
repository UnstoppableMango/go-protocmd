package builder

import "google.golang.org/protobuf/reflect/protoreflect"

func ProtoField(fd protoreflect.FieldDescriptor, v protoreflect.Value) (args []string) {
	name := fd.TextName()

	switch {
	case fd.IsList():
		return ProtoList(fd, v.List())
	case fd.IsMap():
		return ProtoMap(name, v.Map())
	}
	return ProtoValue(fd, v)
}

func ProtoValue(fd protoreflect.FieldDescriptor, value protoreflect.Value) []string {
	switch kind := fd.Kind(); kind {
	case protoreflect.BoolKind:
		return Opt(name(fd), value.IsValid, value.Bool)
	case protoreflect.StringKind:
		return Arg(name(fd), value.IsValid, value.String)
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind,
		protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return Arg(name(fd), value.IsValid, value.Int)
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind,
		protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return Arg(name(fd), value.IsValid, value.Uint)
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		return Arg(name(fd), value.IsValid, value.Float)
	case protoreflect.EnumKind:
		return Arg(name(fd), value.IsValid, value.Enum)
	case protoreflect.BytesKind:
		return Arg(name(fd), value.IsValid, value.Bytes)
	case protoreflect.MessageKind, protoreflect.GroupKind:
		return ProtoMessage(value.Message())
	default:
		panic("unsupported kind: " + kind.String())
	}
}

func ProtoList(fd protoreflect.FieldDescriptor, list protoreflect.List) (args []string) {
	for i := range list.Len() {
		v := ProtoValue(fd, list.Get(i))
		args = append(args, v...)
	}
	return
}

func ProtoMap(fd protoreflect.FieldDescriptor, m protoreflect.Map) (args []string) {
	kt := fd.MapKey().Kind()
	vt := fd.MapValue().Kind()

	for key, value := range m.Range {
		args = append(args, key, value)
	}
	return
}

func ProtoMessage(msg protoreflect.Message) (args []string) {
	for fd, v := range msg.Range {
		args = append(args, ProtoField(fd, v)...)
	}
	return
}

func name(fd protoreflect.FieldDescriptor) string {
	return fd.TextName()
}
