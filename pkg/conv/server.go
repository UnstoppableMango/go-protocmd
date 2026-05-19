package conv

import (
	"context"

	cmdv1alpha1 "github.com/unstoppablemango/go-protocmd/gen/dev/unmango/cmd/v1alpha1"
)

type Server struct {
	cmdv1alpha1.UnimplementedConversionServiceServer
}

func NewServer() cmdv1alpha1.ConversionServiceServer {
	return &Server{}
}

func HandleArgs(ctx context.Context, req *cmdv1alpha1.ArgsRequest) (*cmdv1alpha1.ArgsResponse, error) {
	args, err := Args(req.GetSpec().ProtoReflect())
	if err != nil {
		return nil, err
	}
	b := &cmdv1alpha1.ArgsResponse_builder{
		Args: args,
	}
	return b.Build(), nil
}

// Args implements [cmdv1alpha1.ConversionServiceServer].
func (*Server) Args(ctx context.Context, req *cmdv1alpha1.ArgsRequest) (*cmdv1alpha1.ArgsResponse, error) {
	return HandleArgs(ctx, req)
}
