package cmd

import (
	"context"
	"fmt"

	cmdv1alpha1 "github.com/unstoppablemango/go-protocmd/gen/dev/unmango/cmd/v1alpha1"
	"github.com/unstoppablemango/go-protocmd/pkg/process"
	"google.golang.org/grpc"
)

type Server struct {
	cmdv1alpha1.UnimplementedCommandServiceServer
}

func NewServer() cmdv1alpha1.CommandServiceServer {
	return &Server{}
}

func (Server) Run(ctx context.Context, req *cmdv1alpha1.RunRequest) (*cmdv1alpha1.RunResponse, error) {
	if !req.HasProcess() {
		return nil, fmt.Errorf("process is required")
	}

	cmd, err := process.CommandContext(ctx, req.GetProcess())
	if err != nil {
		return nil, fmt.Errorf("converting process to cmd: %w", err)
	}
	if err = cmd.Run(); err != nil {
		return nil, err
	}

	res := &cmdv1alpha1.RunResponse_builder{
		// Stdout: , // TODO
		// Stderr: , // TODO
		ExitCode: new(int32(cmd.ProcessState.ExitCode())),
	}
	return res.Build(), nil
}

func (Server) Exec(req *cmdv1alpha1.ExecRequest, stream grpc.ServerStreamingServer[cmdv1alpha1.ExecResponse]) error {
	panic("unimplemented")
}

func (Server) Start(ctx context.Context, req *cmdv1alpha1.StartRequest) (*cmdv1alpha1.StartResponse, error) {
	panic("unimplemented")
}

func (Server) Wait(ctx context.Context, req *cmdv1alpha1.WaitRequest) (*cmdv1alpha1.WaitResponse, error) {
	panic("unimplemented")
}

func (Server) Signal(ctx context.Context, req *cmdv1alpha1.SignalRequest) (*cmdv1alpha1.SignalResponse, error) {
	panic("unimplemented")
}
