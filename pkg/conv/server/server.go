package server

import (
	"context"

	cmdv1alpha2 "github.com/unstoppablemango/go-protocmd/gen/dev/unmango/cmd/v1alpha2"
	"github.com/unstoppablemango/go-protocmd/pkg/conv"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	cmdv1alpha2.UnimplementedConversionServiceServer
}

func NewServer() cmdv1alpha2.ConversionServiceServer {
	return &Server{}
}

func (*Server) FromJson(ctx context.Context, req *cmdv1alpha2.FromJsonRequest) (*cmdv1alpha2.FromJsonResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (*Server) FromYaml(ctx context.Context, req *cmdv1alpha2.FromYamlRequest) (*cmdv1alpha2.FromYamlResponse, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (*Server) FromProto(ctx context.Context, req *cmdv1alpha2.FromProtoRequest) (*cmdv1alpha2.FromProtoResponse, error) {
	args, err := conv.Args(req.GetSpec().ProtoReflect())
	if err != nil {
		return nil, err
	}

	proc := &cmdv1alpha2.Process_builder{
		// TODO
		Args: args,
	}
	res := &cmdv1alpha2.FromProtoResponse_builder{
		Process: proc.Build(),
	}
	return res.Build(), nil
}
