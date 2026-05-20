package server

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"syscall"

	cmdv1alpha1 "github.com/unstoppablemango/go-protocmd/gen/dev/unmango/cmd/v1alpha1"
	"github.com/unstoppablemango/go-protocmd/pkg/process"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	cmdv1alpha1.UnimplementedCommandServiceServer
	mu    sync.Mutex
	procs map[string]*exec.Cmd
}

func NewServer() cmdv1alpha1.CommandServiceServer {
	return &Server{
		procs: make(map[string]*exec.Cmd),
	}
}

func (s *Server) Run(ctx context.Context, req *cmdv1alpha1.RunRequest) (*cmdv1alpha1.RunResponse, error) {
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

	exitCode := int32(cmd.ProcessState.ExitCode())
	res := &cmdv1alpha1.RunResponse_builder{
		Stdout:   stdout.Bytes(),
		Stderr:   stderr.Bytes(),
		ExitCode: &exitCode,
	}
	return res.Build(), nil
}

func (s *Server) Exec(req *cmdv1alpha1.ExecRequest, srv grpc.ServerStreamingServer[cmdv1alpha1.ExecResponse]) error {
	if !req.HasProcess() {
		return invalid("process is required")
	}

	cmd := process.CommandContext(srv.Context(), req.GetProcess())

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	var mu sync.Mutex
	send := func(resp *cmdv1alpha1.ExecResponse) error {
		mu.Lock()
		defer mu.Unlock()
		return srv.Send(resp)
	}

	var wg sync.WaitGroup

	pipe := func(r io.Reader, build func([]byte) *cmdv1alpha1.ExecResponse) {
		wg.Go(func() {
			buf := make([]byte, 4096)
			for {
				n, err := r.Read(buf)
				if n > 0 {
					data := make([]byte, n)
					copy(data, buf[:n])
					send(build(data))
				}
				if err != nil {
					return
				}
			}
		})
	}

	pipe(stdout, func(b []byte) *cmdv1alpha1.ExecResponse {
		res := &cmdv1alpha1.ExecResponse_builder{Stdout: b}
		return res.Build()
	})
	pipe(stderr, func(b []byte) *cmdv1alpha1.ExecResponse {
		res := &cmdv1alpha1.ExecResponse_builder{Stderr: b}
		return res.Build()
	})

	wg.Wait()

	if err := cmd.Wait(); err != nil {
		var exitErr *exec.ExitError
		if !errors.As(err, &exitErr) {
			return err
		}
	}

	exitCode := int32(cmd.ProcessState.ExitCode())
	exit := &cmdv1alpha1.ExitResult_builder{Code: &exitCode}
	res := &cmdv1alpha1.ExecResponse_builder{Exit: exit.Build()}
	return srv.Send(res.Build())
}

func (s *Server) Start(ctx context.Context, req *cmdv1alpha1.StartRequest) (*cmdv1alpha1.StartResponse, error) {
	if !req.HasProcess() {
		return nil, invalid("process is required")
	}

	cmd := process.CommandContext(ctx, req.GetProcess())
	if err := cmd.Start(); err != nil {
		return nil, err
	}

	id := strconv.Itoa(cmd.Process.Pid)

	s.mu.Lock()
	s.procs[id] = cmd
	s.mu.Unlock()

	res := &cmdv1alpha1.StartResponse_builder{Id: &id}
	return res.Build(), nil
}

func (s *Server) Wait(ctx context.Context, req *cmdv1alpha1.WaitRequest) (*cmdv1alpha1.WaitResponse, error) {
	if !req.HasId() {
		return nil, invalid("id is required")
	}

	id := req.GetId()

	s.mu.Lock()
	cmd, ok := s.procs[id]
	s.mu.Unlock()

	if !ok {
		return nil, notFound("process %s not found", id)
	}

	if err := cmd.Wait(); err != nil {
		var exitErr *exec.ExitError
		if !errors.As(err, &exitErr) {
			return nil, err
		}
	}

	s.mu.Lock()
	delete(s.procs, id)
	s.mu.Unlock()

	res := &cmdv1alpha1.WaitResponse_builder{
		ExitCode: new(int32(cmd.ProcessState.ExitCode())),
	}
	return res.Build(), nil
}

func (s *Server) Signal(ctx context.Context, req *cmdv1alpha1.SignalRequest) (*cmdv1alpha1.SignalResponse, error) {
	if !req.HasId() {
		return nil, invalid("id is required")
	}

	id := req.GetId()

	s.mu.Lock()
	cmd, ok := s.procs[id]
	s.mu.Unlock()

	if !ok {
		return nil, notFound("process %s not found", id)
	}

	sig, err := osSignal(req.GetSignal())
	if err != nil {
		return nil, err
	}

	if err := cmd.Process.Signal(sig); err != nil {
		return nil, err
	}

	res := &cmdv1alpha1.SignalResponse_builder{}
	return res.Build(), nil
}

func osSignal(sig cmdv1alpha1.Signal) (os.Signal, error) {
	switch sig {
	case cmdv1alpha1.Signal_SIGNAL_HUP:
		return syscall.SIGHUP, nil
	case cmdv1alpha1.Signal_SIGNAL_INT:
		return syscall.SIGINT, nil
	case cmdv1alpha1.Signal_SIGNAL_QUIT:
		return syscall.SIGQUIT, nil
	case cmdv1alpha1.Signal_SIGNAL_KILL:
		return syscall.SIGKILL, nil
	case cmdv1alpha1.Signal_SIGNAL_TERM:
		return syscall.SIGTERM, nil
	case cmdv1alpha1.Signal_SIGNAL_USR1:
		return syscall.SIGUSR1, nil
	case cmdv1alpha1.Signal_SIGNAL_USR2:
		return syscall.SIGUSR2, nil
	case cmdv1alpha1.Signal_SIGNAL_CONT:
		return syscall.SIGCONT, nil
	case cmdv1alpha1.Signal_SIGNAL_STOP:
		return syscall.SIGSTOP, nil
	default:
		return nil, invalid("unsupported signal: %s", sig)
	}
}

func invalid(msg string, args ...any) error {
	return status.Errorf(codes.InvalidArgument, msg, args...)
}

func notFound(msg string, args ...any) error {
	return status.Errorf(codes.NotFound, msg, args...)
}
