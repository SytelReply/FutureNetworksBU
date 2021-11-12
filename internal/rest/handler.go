package rest

import (
	vlanproto "code-challenge/protos"
	"encoding/json"
	"log"
	"net/http"

	"google.golang.org/protobuf/types/known/emptypb"
)

type handler struct {
	grpcClient vlanproto.V1Client
}

type postVlanRequest struct {
	ID   string `json:"id"`
	VLAN string `json:"vlan"`
}

func (h *handler) postVlan(w http.ResponseWriter, r *http.Request) {
	var req postVlanRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Fatal(err)
	}

	_, err := h.grpcClient.SaveVLAN(r.Context(), &vlanproto.VLAN{
		Id:   req.ID,
		Vlan: req.VLAN,
	})
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) getVlans(w http.ResponseWriter, r *http.Request) {
	vlans, err := h.grpcClient.GetVLANs(r.Context(), &emptypb.Empty{})
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.Marshal(vlans)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(b)
	if err != nil {
		log.Fatal(err)
	}

}
