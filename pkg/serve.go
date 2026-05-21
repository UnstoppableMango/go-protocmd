package protocmd

import (
	"net"

	"charm.land/log/v2"
	"google.golang.org/grpc"
)

func ListenAndServe() error {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}

	srv := grpc.NewServer()
	// cmdv1alpha1.RegisterExecutionServiceServer(srv, cmdserver.NewServer())

	log.Info("Listen and serving", "addr", lis.Addr())
	if err := srv.Serve(lis); err != nil {
		return err
	}
	return nil
}
