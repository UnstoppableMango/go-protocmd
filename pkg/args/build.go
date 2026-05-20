package args

import (
	"context"

	cmdv1alpha2 "github.com/unstoppablemango/go-protocmd/gen/dev/unmango/cmd/v1alpha2"
	"github.com/unstoppablemango/go-protocmd/pkg/args/builder"
)

func FromJson(ctx context.Context, req *cmdv1alpha2.FromJsonRequest) (*cmdv1alpha2.FromJsonResponse, error) {
	panic("unimplemented")
}

func FromProto(ctx context.Context, req *cmdv1alpha2.FromProtoRequest) (*cmdv1alpha2.FromProtoResponse, error) {
	args := builder.ProtoMessage(req.GetSpec().ProtoReflect())

	res := &cmdv1alpha2.FromProtoResponse_builder{
		Process: toProcess(args),
	}
	return res.Build(), nil
}

func FromYaml(ctx context.Context, req *cmdv1alpha2.FromYamlRequest) (*cmdv1alpha2.FromYamlResponse, error) {
	panic("unimplemented")
}

func toProcess(args []string) *cmdv1alpha2.Process {
	proc := &cmdv1alpha2.Process_builder{
		Args: args,
	}
	return proc.Build()
}
