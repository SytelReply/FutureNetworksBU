package main

import (
	"github.com/James-Milligan/FutureNetworksBU/internal/rest"
	"log"
)

func main() {
	log.Fatal(rest.Server.Run())
}
