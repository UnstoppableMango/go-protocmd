package stream

import (
	"fmt"
	"io"
	"strings"

	cmdv1alpha1 "github.com/unstoppablemango/go-protocmd/gen/dev/unmango/cmd/v1alpha1"
)

func ToReader(stream *cmdv1alpha1.Stream) (io.Reader, error) {
	switch kind := stream.WhichKind(); kind {
	case cmdv1alpha1.Stream_Null_case:
		return strings.NewReader(""), nil
	case cmdv1alpha1.Stream_File_case:
		fallthrough // TODO
	case cmdv1alpha1.Stream_Inherit_case:
		fallthrough // TODO
	case cmdv1alpha1.Stream_Pipe_case:
		fallthrough // TODO
	default:
		return nil, fmt.Errorf("unsupported stream kind: %s", kind)
	}
}

func ToWriter(stream *cmdv1alpha1.Stream) (io.Writer, error) {
	switch kind := stream.WhichKind(); kind {
	case cmdv1alpha1.Stream_Null_case:
		return io.Discard, nil
	case cmdv1alpha1.Stream_File_case:
		fallthrough // TODO
	case cmdv1alpha1.Stream_Inherit_case:
		fallthrough // TODO
	case cmdv1alpha1.Stream_Pipe_case:
		fallthrough // TODO
	default:
		return nil, fmt.Errorf("unsupported stream kind: %s", kind)
	}
}
