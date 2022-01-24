package main

import (
	"github.com/James-Milligan/FutureNetworksBU/internal/grpc"
	"github.com/goioc/di"
	"log"
)

func init() {
	grpc.BuildDependencyContainer()
}

func main() {
	app := di.GetInstance("app").(*grpc.App)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
