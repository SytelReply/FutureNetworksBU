package rest

import (
	"encoding/json"
	"fmt"
	vlanproto "github.com/James-Milligan/FutureNetworksBU/protos"
	"github.com/gorilla/mux"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type handler struct {
	grpcClient vlanproto.V1Client
}

type postVlanRequest struct {
	ID   string `json:"id"`
	VLAN string `json:"vlan"`
}

// postVlan stores a new key value pair in the system
func (h *handler) postVlan(w http.ResponseWriter, r *http.Request) {
	var req postVlanRequest

	if r.Body == nil {
		fmt.Println("No body provided in request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if req.VLAN == "" || req.ID == "" {
		fmt.Println("Malformed request received, missing vlan and/or id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf("Received VLAN and Id pair %s %s\n", req.VLAN, req.ID)
	_, err := h.grpcClient.SaveVLAN(r.Context(), &vlanproto.SaveVLANRequest{
		Vlan: &vlanproto.VLAN{
			Id:   req.ID,
			Vlan: req.VLAN,
		},
	})

	if err != nil {
		errStatus, _ := status.FromError(err)
		if errStatus.Code() == codes.InvalidArgument {
			fmt.Println(errStatus.Message())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(errStatus.Message())
		return
	}

	w.WriteHeader(http.StatusOK)
}

// getVlans Fetches all currently stored key value pairs currently stored in the system
func (h *handler) getVlans(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching all VLANs")

	vlans, err := h.grpcClient.GetVLANs(r.Context(), &vlanproto.GetVLANsRequest{})
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(vlans)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(b)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

// getVlan fetches the VLAN object for a given id
func (h *handler) getVlan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["id"] == "" {
		fmt.Println("No id in request")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Printf("Fetching VLAN with id %s\n", vars["id"])

	vlan, err := h.grpcClient.GetVLAN(r.Context(), &vlanproto.GetVLANRequest{
		Id: vars["id"],
	})
	if err != nil {
		errStatus, _ := status.FromError(err)
		if errStatus.Code() == codes.NotFound {
			fmt.Println(errStatus.Message())
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(errStatus.Message())
		return
	}

	b, err := json.Marshal(vlan)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(b)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
