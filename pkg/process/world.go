package process

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/unmango/go/world"
	cmdv1alpha1 "github.com/unstoppablemango/go-protocmd/gen/dev/unmango/cmd/v1alpha1"
)

// type World interface {
// 	Open(*cmdv1alpha1.File) (world.File, error)
// 	Writer(*cmdv1alpha1.Stream) (io.Writer, error)
// 	Reader(*cmdv1alpha1.Stream) (io.Reader, error)
// }

type World struct {
	ctx context.Context
	os  world.Os
}

func In(w world.IO) World {
	return World{
		ctx: context.Background(),
		os:  w.Os(),
	}
}

func FromContext(ctx context.Context) World {
	w := world.FromContext(ctx)

	return World{
		ctx: ctx,
		os:  w.Os(),
	}
}

func (w World) Command(proc *cmdv1alpha1.Process) (*exec.Cmd, error) {
	if !proc.HasPath() {
		return nil, fmt.Errorf("path is required")
	}

	cmd := exec.CommandContext(w.ctx, proc.GetPath())
	cmd.Args = proc.GetArgs()
	cmd.Env = proc.GetEnv()

	if proc.HasCwd() {
		cmd.Dir = proc.GetCwd()
	}
	if proc.HasStdio() {
		if err := w.applyStdio(cmd, proc.GetStdio()); err != nil {
			return nil, err
		}
	}
	if proc.HasTerminal() {
		// TODO
	}
	if proc.HasUser() {
		// TODO
	}
	return cmd, nil
}

func (w World) applyStdio(cmd *exec.Cmd, stdio *cmdv1alpha1.Stdio) error {
	if stdio.HasStdout() {
		stdout, err := w.writer(stdio.GetStdout())
		if err != nil {
			return err
		}
		cmd.Stdout = stdout
	}
	if stdio.HasStderr() {
		stderr, err := w.writer(stdio.GetStderr())
		if err != nil {
			return err
		}
		cmd.Stderr = stderr
	}
	if stdio.HasStdin() {
		// TODO
	}
	return nil
}

func (w World) writer(stream *cmdv1alpha1.Stream) (io.Writer, error) {
	switch kind := stream.WhichKind(); kind {
	case cmdv1alpha1.Stream_Null_case, cmdv1alpha1.Stream_Kind_not_set_case:
		return io.Discard, nil
	case cmdv1alpha1.Stream_File_case:
		return w.open(stream.GetFile())
	case cmdv1alpha1.Stream_Inherit_case:
		fallthrough // TODO
	case cmdv1alpha1.Stream_Pipe_case:
		fallthrough // TODO
	default:
		panic("unsupported stream kind: " + kind.String())
	}
}

func (w World) open(f *cmdv1alpha1.File) (world.File, error) {
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
	return w.os.OpenFile(f.GetPath(), flag, os.ModePerm)
}
