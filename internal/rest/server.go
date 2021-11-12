package rest

import (
	grpcserver "code-challenge/internal/grpc"
	vlanproto "code-challenge/protos"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/grpc"

	"github.com/gorilla/mux"
)

var Server *server

const port = 8081

type server struct {
	port   int
	router *mux.Router
}

func init() {
	r := mux.NewRouter()
	h := &handler{
		grpcClient: grpcClient(),
	}
	r.HandleFunc("/vlans", h.postVlan).Methods("POST", "OPTIONS")
	r.HandleFunc("/vlans", h.getVlans).Methods("GET", "OPTIONS")

	Server = &server{port: port, router: r}
}

func grpcClient() vlanproto.V1Client {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", grpcserver.Port), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return vlanproto.NewV1Client(conn)
}

func (s *server) Run() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.router)
}
