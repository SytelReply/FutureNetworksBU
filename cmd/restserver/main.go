package main

import (
	"code-challenge/internal/rest"
	"log"
)

func main() {
	log.Fatal(rest.Server.Run())
}
