package grpc

import (
	vlanproto "code-challenge/protos"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

var Server *server

var GRPC_SERVER_PORT = os.Getenv("GRPC_SERVER_PORT")
var REST_SERVER_PORT = os.Getenv("REST_SERVER_PORT")

var grpcListenAddress = fmt.Sprintf(":%s", GRPC_SERVER_PORT)

func init() {
	if GRPC_SERVER_PORT == "" || REST_SERVER_PORT == "" {
		log.Fatal("Missing environment variables")
	}
	fmt.Printf("Starting gRPC server on port %s\n", GRPC_SERVER_PORT)
	srv := grpc.NewServer()
	vlanproto.RegisterV1Server(srv, &handler{})
	Server = &server{
		grpcListenAddress: grpcListenAddress,
		srv:               srv,
	}
}

type server struct {
	grpcListenAddress string
	srv               *grpc.Server
}

func (s *server) Run() error {
	lis, err := net.Listen("tcp", s.grpcListenAddress)
	if err != nil {
		return err
	}

	return s.srv.Serve(lis)
}
