package process

import (
	"context"
	"os/exec"

	cmdv1alpha2 "github.com/unstoppablemango/go-protocmd/gen/dev/unmango/cmd/v1alpha2"
)

func CommandContext(ctx context.Context, proc *cmdv1alpha2.Process) *exec.Cmd {
	cmd := exec.CommandContext(ctx, proc.GetPath())
	configure(cmd, proc)
	return cmd
}

func Command(proc *cmdv1alpha2.Process) *exec.Cmd {
	cmd := exec.Command(proc.GetPath())
	configure(cmd, proc)
	return cmd
}

func configure(cmd *exec.Cmd, proc *cmdv1alpha2.Process) {
	cmd.Args = proc.GetArgs()
	cmd.Env = proc.GetEnv()

	if proc.HasCwd() {
		cmd.Dir = proc.GetCwd()
	}
}
