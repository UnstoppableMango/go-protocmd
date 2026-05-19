package process

import (
	"context"
	"fmt"
	"os/exec"

	cmdv1alpha1 "github.com/unstoppablemango/go-protocmd/gen/dev/unmango/cmd/v1alpha1"
	"github.com/unstoppablemango/go-protocmd/pkg/stream"
)

func CommandContext(ctx context.Context, proc *cmdv1alpha1.Process) (*exec.Cmd, error) {
	if !proc.HasPath() {
		return nil, fmt.Errorf("path is required")
	}

	cmd := exec.CommandContext(ctx, proc.GetPath())
	Apply(cmd, proc)
	return cmd, nil
}

func Apply(cmd *exec.Cmd, proc *cmdv1alpha1.Process) {
	cmd.Args = proc.GetArgs()
	if proc.HasCwd() {
		cmd.Dir = proc.GetCwd()
	}
	if proc.HasStdio() {
		ApplyStdio(cmd, proc.GetStdio())
	}
}

func ApplyStdio(cmd *exec.Cmd, stdio *cmdv1alpha1.Stdio) error {
	var err error
	if stdio.HasStdout() {
		if cmd.Stdout, err = stream.ToWriter(stdio.GetStdout()); err != nil {
			return err
		}
	}
	if stdio.HasStderr() {
		if cmd.Stderr, err = stream.ToWriter(stdio.GetStderr()); err != nil {
			return err
		}
	}
	if stdio.HasStdin() {
		if cmd.Stdin, err = stream.ToReader(stdio.GetStdin()); err != nil {
			return err
		}
	}
	return nil
}
