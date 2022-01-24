package main

import (
	"github.com/James-Milligan/FutureNetworksBU/internal/rest"
	"github.com/goioc/di"
	"log"
)

func init() {
	rest.BuildDependencyContainer()
}

func main() {
	app := di.GetInstance("app").(*rest.App)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
