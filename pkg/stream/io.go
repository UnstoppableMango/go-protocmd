package stream

import (
	"io"
	"strings"

	cmdv1alpha1 "github.com/unstoppablemango/go-protocmd/gen/dev/unmango/cmd/v1alpha1"
)

func ToReader(stream *cmdv1alpha1.Stream) io.Reader {
	switch kind := stream.WhichKind(); kind {
	case cmdv1alpha1.Stream_Null_case, cmdv1alpha1.Stream_Kind_not_set_case:
		return strings.NewReader("") // Empty reader
	case cmdv1alpha1.Stream_File_case:
		fallthrough // TODO
	case cmdv1alpha1.Stream_Inherit_case:
		fallthrough // TODO
	case cmdv1alpha1.Stream_Pipe_case:
		fallthrough // TODO
	default:
		panic("unsupported stream kind: " + kind.String())
	}
}

func ToWriter(stream *cmdv1alpha1.Stream) io.Writer {
	switch kind := stream.WhichKind(); kind {
	case cmdv1alpha1.Stream_Null_case, cmdv1alpha1.Stream_Kind_not_set_case:
		return io.Discard
	case cmdv1alpha1.Stream_File_case:
		fallthrough // TODO
	case cmdv1alpha1.Stream_Inherit_case:
		fallthrough // TODO
	case cmdv1alpha1.Stream_Pipe_case:
		fallthrough // TODO
	default:
		panic("unsupported stream kind: " + kind.String())
	}
}
