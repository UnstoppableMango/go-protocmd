package server

import (
	"bytes"
	"context"
	"errors"
	"os/exec"

	cmdv1alpha2 "github.com/unstoppablemango/go-protocmd/gen/dev/unmango/cmd/v1alpha2"
	"github.com/unstoppablemango/go-protocmd/pkg/process"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	cmdv1alpha2.UnimplementedExecutionServiceServer
}

func NewServer() cmdv1alpha2.ExecutionServiceServer {
	return &Server{}
}

func (s *Server) Run(ctx context.Context, req *cmdv1alpha2.RunRequest) (*cmdv1alpha2.RunResponse, error) {
	if !req.HasProcess() {
		return nil, invalid("process is required")
	}

	var stdout, stderr bytes.Buffer
	cmd := process.CommandContext(ctx, req.GetProcess())
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		var exitErr *exec.ExitError
		if !errors.As(err, &exitErr) {
			return nil, err
		}
	}

	exit := cmdv1alpha2.ExitResult_builder{
		Code: new(int32(cmd.ProcessState.ExitCode())),
	}
	res := &cmdv1alpha2.RunResponse_builder{
		Stdout: stdout.Bytes(),
		Stderr: stderr.Bytes(),
		Exit:   exit.Build(),
	}
	return res.Build(), nil
}

func invalid(msg string, args ...any) error {
	return status.Errorf(codes.InvalidArgument, msg, args...)
}
