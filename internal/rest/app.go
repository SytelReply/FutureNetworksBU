package rest

import (
	"fmt"
	"github.com/James-Milligan/FutureNetworksBU/internal/common"
	"github.com/gorilla/mux"
	"net/http"
)

type App struct {
	config     *common.Config `di.inject:"config" di.scope:"singleton"`
	grpcClient *GrpcClient    `di.inject:"grpcClient" di.scope:"singleton"`
}

func (a *App) Run() error {
	fmt.Printf("Starting REST server on port %s\n", a.config.RestPort)

	r := mux.NewRouter()
	h := &handler{
		grpcClient: a.grpcClient.GetInstance(),
	}
	r.HandleFunc("/vlans", h.postVlan).Methods("POST", "OPTIONS")
	r.HandleFunc("/vlans", h.getVlans).Methods("GET", "OPTIONS")
	r.HandleFunc("/vlans/{id}", h.getVlan).Methods("GET", "OPTIONS")

	fmt.Println("Listening...")
	return http.ListenAndServe(fmt.Sprintf(":%s", a.config.RestPort), r)
}
