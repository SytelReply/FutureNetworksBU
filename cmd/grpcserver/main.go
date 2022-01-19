package main

import (
	"github.com/James-Milligan/FutureNetworksBU/internal/grpc"
	"log"
)

func main() {
	if err := grpc.Server.Run(); err != nil {
		log.Fatal(err)
	}
}
