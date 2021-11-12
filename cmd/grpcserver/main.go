package main

import (
	"code-challenge/internal/grpc"
	"log"
)

func main() {
	if err := grpc.Server.Run(); err != nil {
		log.Fatal(err)
	}
}
