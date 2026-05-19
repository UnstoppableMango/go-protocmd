package protocmd

import (
	"context"
	"fmt"

	cmdv1alpha1 "github.com/unstoppablemango/go-protocmd/gen/dev/unmango/cmd/v1alpha1"
	"github.com/unstoppablemango/go-protocmd/pkg/conv"
)

type ConversionServer struct {
	cmdv1alpha1.UnimplementedConversionServiceServer
}

func NewConversionServer() cmdv1alpha1.ConversionServiceServer {
	return &ConversionServer{}
}

func Args(ctx context.Context, req *cmdv1alpha1.ArgsRequest) (*cmdv1alpha1.ArgsResponse, error) {
	msg, err := req.GetSpec().UnmarshalNew()
	if err != nil {
		return nil, fmt.Errorf("reading spec: %w", err)
	}
	args, err := conv.ProtoArgs(msg.ProtoReflect())
	if err != nil {
		return nil, fmt.Errorf("converting proto args: %w", err)
	}

	res := &cmdv1alpha1.ArgsResponse_builder{
		Args: args,
	}
	return res.Build(), nil
}

// Args implements [cmdv1alpha1.ConversionServiceServer].
func (*ConversionServer) Args(ctx context.Context, req *cmdv1alpha1.ArgsRequest) (*cmdv1alpha1.ArgsResponse, error) {
	return Args(ctx, req)
}
