package protocmd

import cmdv1alpha1 "github.com/unstoppablemango/go-protocmd/gen/dev/unmango/cmd/v1alpha1"

type CommandServer struct {
	cmdv1alpha1.UnimplementedCommandServiceServer
}

func NewCommandServer() cmdv1alpha1.CommandServiceServer {
	return &CommandServer{}
}
