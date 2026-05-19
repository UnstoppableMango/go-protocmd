package stream

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/unmango/go/world"
	cmdv1alpha1 "github.com/unstoppablemango/go-protocmd/gen/dev/unmango/cmd/v1alpha1"
)

func ToReader(stream *cmdv1alpha1.Stream, w world.IO) io.Reader {
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

type World struct{ world.IO }

func From(w world.IO) World {
	return World{w}
}

func (w World) Open(f *cmdv1alpha1.File) (world.File, error) {
	return open(w.Os(), f)
}

func (w World) Reader(stream *cmdv1alpha1.Stream) (io.Reader, error) {
	switch kind := stream.WhichKind(); kind {
	case cmdv1alpha1.Stream_Null_case, cmdv1alpha1.Stream_Kind_not_set_case:
		return strings.NewReader(""), nil // Empty reader
	case cmdv1alpha1.Stream_File_case:
		return w.Open(stream.GetFile())
	case cmdv1alpha1.Stream_Inherit_case:
		fallthrough // TODO
	case cmdv1alpha1.Stream_Pipe_case:
		fallthrough // TODO
	default:
		panic("unsupported stream kind: " + kind.String())
	}
}

func (w World) Writer(stream *cmdv1alpha1.Stream) (io.Writer, error) {
	switch kind := stream.WhichKind(); kind {
	case cmdv1alpha1.Stream_Null_case, cmdv1alpha1.Stream_Kind_not_set_case:
		return io.Discard, nil
	case cmdv1alpha1.Stream_File_case:
		return w.Open(stream.GetFile())
	case cmdv1alpha1.Stream_Inherit_case:
		fallthrough // TODO
	case cmdv1alpha1.Stream_Pipe_case:
		fallthrough // TODO
	default:
		panic("unsupported stream kind: " + kind.String())
	}
}

func open(w world.Os, f *cmdv1alpha1.File) (world.File, error) {
	if !f.HasPath() {
		return nil, fmt.Errorf("path is required")
	}

	var flag int
	switch f.GetMode() {
	case cmdv1alpha1.OpenMode_OPEN_MODE_APPEND:
		flag |= os.O_APPEND
	case cmdv1alpha1.OpenMode_OPEN_MODE_READ:
		flag |= os.O_RDONLY
	case cmdv1alpha1.OpenMode_OPEN_MODE_WRITE:
		flag |= os.O_WRONLY
	}
	return w.OpenFile(f.GetPath(), flag, os.ModePerm)
}
