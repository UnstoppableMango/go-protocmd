package protocmd

import (
	"net"

	cmdv1alpha1 "github.com/unstoppablemango/go-protocmd/gen/dev/unmango/cmd/v1alpha1"
	"google.golang.org/grpc"
)

func ListenAndServe() error {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}

	srv := grpc.NewServer()
	cmdv1alpha1.RegisterConversionServiceServer(srv, NewConversionServer())
	cmdv1alpha1.RegisterCommandServiceServer(srv, NewCommandServer())

	if err := srv.Serve(lis); err != nil {
		return err
	}
	return nil
}
