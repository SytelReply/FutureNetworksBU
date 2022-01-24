package rest

import (
	"fmt"
	"github.com/James-Milligan/FutureNetworksBU/internal/common"
	vlanproto "github.com/James-Milligan/FutureNetworksBU/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type GrpcClient struct {
	client vlanproto.V1Client
	config *common.Config `di.inject:"config" di.scope:"singleton"`
}

func (g *GrpcClient) GetInstance() vlanproto.V1Client {
	if g.client == nil {
		conn, err := grpc.Dial(fmt.Sprintf("%s:%s", g.config.GrpcAddress, g.config.GrpcPort), grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		g.client = vlanproto.NewV1Client(conn)
	}
	return g.client
}
