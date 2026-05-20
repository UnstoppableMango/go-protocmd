package process

import (
	"context"
	"fmt"
	"os/exec"

	"github.com/unmango/go/world"
	cmdv1alpha2 "github.com/unstoppablemango/go-protocmd/gen/dev/unmango/cmd/v1alpha2"
)

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

func (w World) Command(proc *cmdv1alpha2.Process) (*exec.Cmd, error) {
	if !proc.HasPath() {
		return nil, fmt.Errorf("path is required")
	}

	cmd := exec.CommandContext(w.ctx, proc.GetPath())
	cmd.Args = proc.GetArgs()
	cmd.Env = proc.GetEnv()

	if proc.HasCwd() {
		cmd.Dir = proc.GetCwd()
	}
	return cmd, nil
}
