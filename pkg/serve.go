package protocmd

import (
	"net"

	"charm.land/log/v2"
	cmdv1alpha2 "github.com/unstoppablemango/go-protocmd/gen/dev/unmango/cmd/v1alpha2"
	convserver "github.com/unstoppablemango/go-protocmd/pkg/conv/server"
	cmdserver "github.com/unstoppablemango/go-protocmd/pkg/process/server"
	"google.golang.org/grpc"
)

func ListenAndServe() error {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}

	srv := grpc.NewServer()
	cmdv1alpha2.RegisterConversionServiceServer(srv, convserver.NewServer())
	cmdv1alpha2.RegisterExecutionServiceServer(srv, cmdserver.NewServer())

	log.Info("Listen and serving", "addr", lis.Addr())
	if err := srv.Serve(lis); err != nil {
		return err
	}
	return nil
}
