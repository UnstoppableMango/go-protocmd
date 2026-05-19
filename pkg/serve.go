package protocmd

import (
	"net"

	"charm.land/log/v2"
	cmdv1alpha1 "github.com/unstoppablemango/go-protocmd/gen/dev/unmango/cmd/v1alpha1"
	"github.com/unstoppablemango/go-protocmd/pkg/cmd"
	"github.com/unstoppablemango/go-protocmd/pkg/conv"
	"google.golang.org/grpc"
)

func ListenAndServe() error {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}

	srv := grpc.NewServer()
	cmdv1alpha1.RegisterConversionServiceServer(srv, conv.NewServer())
	cmdv1alpha1.RegisterCommandServiceServer(srv, cmd.NewServer())

	log.Info("Listen and serving", "addr", lis.Addr())
	if err := srv.Serve(lis); err != nil {
		return err
	}
	return nil
}
