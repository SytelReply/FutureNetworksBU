package grpc

import (
	"fmt"
	"github.com/James-Milligan/FutureNetworksBU/internal/common"
	vlanproto "github.com/James-Milligan/FutureNetworksBU/protos"
	"google.golang.org/grpc"
	"net"
)

type App struct {
	config *common.Config `di.inject:"config" di.scope:"singleton"`
}

func (a *App) Run() error {
	fmt.Printf("Starting gRPC server on port %s\n", a.config.GrpcPort)

	srv := grpc.NewServer()
	vlanproto.RegisterV1Server(srv, &Handler{})

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", a.config.GrpcPort))
	if err != nil {
		return err
	}
	fmt.Println("Listening...")
	return srv.Serve(lis)
}
