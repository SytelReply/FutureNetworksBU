package grpc

import (
	vlanproto "code-challenge/protos"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

var Server *server

const Port = 8080

var grpcListenAddress = fmt.Sprintf(":%d", Port)

func init() {
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
