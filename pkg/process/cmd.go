package process

import (
	"context"
	"os/exec"

	cmdv1alpha1 "github.com/unstoppablemango/go-protocmd/gen/dev/unmango/cmd/v1alpha1"
	"github.com/unstoppablemango/go-protocmd/pkg/process/stream"
)

func CommandContext(ctx context.Context, proc *cmdv1alpha1.Process) *exec.Cmd {
	cmd := exec.CommandContext(ctx, proc.GetPath())
	configure(cmd, proc)
	return cmd
}

func Command(proc *cmdv1alpha1.Process) *exec.Cmd {
	cmd := exec.Command(proc.GetPath())
	configure(cmd, proc)
	return cmd
}

func configure(cmd *exec.Cmd, proc *cmdv1alpha1.Process) {
	cmd.Args = proc.GetArgs()
	cmd.Env = proc.GetEnv()

	if proc.HasCwd() {
		cmd.Dir = proc.GetCwd()
	}
	if proc.HasStdio() {
		applyStdio(cmd, proc.GetStdio())
	}
	if proc.HasTerminal() {
		// TODO
	}
	if proc.HasUser() {
		// TODO
	}
}

func applyStdio(cmd *exec.Cmd, stdio *cmdv1alpha1.Stdio) {
	if stdio.HasStdout() {
		cmd.Stdout = stream.ToWriter(stdio.GetStdout())
	}
	if stdio.HasStderr() {
		cmd.Stderr = stream.ToWriter(stdio.GetStderr())
	}
	if stdio.HasStdin() {
		cmd.Stdin = stream.ToReader(stdio.GetStdin())
	}
}
