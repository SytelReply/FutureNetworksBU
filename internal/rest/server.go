package rest

import (
	"fmt"
	vlanproto "github.com/James-Milligan/FutureNetworksBU/protos"
	"log"
	"net/http"
	"os"
	"strconv"

	"google.golang.org/grpc"

	"github.com/gorilla/mux"
)

var Server *server

var GRPC_SERVER_PORT = os.Getenv("GRPC_SERVER_PORT")
var REST_SERVER_PORT = os.Getenv("REST_SERVER_PORT")

type server struct {
	port   int
	router *mux.Router
}

func init() {
	if GRPC_SERVER_PORT == "" || REST_SERVER_PORT == "" {
		log.Fatal("Missing environment variables")
	}
	restPortParsed, err := strconv.Atoi(REST_SERVER_PORT)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Starting rest server on port %s\n", REST_SERVER_PORT)
	r := mux.NewRouter()
	h := &handler{
		grpcClient: grpcClient(),
	}
	r.HandleFunc("/vlans", h.postVlan).Methods("POST", "OPTIONS")
	r.HandleFunc("/vlans", h.getVlans).Methods("GET", "OPTIONS")
	r.HandleFunc("/vlans/{id}", h.getVlan).Methods("GET", "OPTIONS")

	Server = &server{port: restPortParsed, router: r}
}

func grpcClient() vlanproto.V1Client {
	conn, err := grpc.Dial(fmt.Sprintf("net-reply-service:%s", GRPC_SERVER_PORT), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return vlanproto.NewV1Client(conn)
}

func (s *server) Run() error {
	fmt.Println("Listening...")
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.router)
}
