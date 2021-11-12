package grpc

import (
	"code-challenge/internal/vlan"
	vlanproto "code-challenge/protos"
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

type handler struct {
	vlanproto.UnimplementedV1Server
}

func (h *handler) SaveVLAN(ctx context.Context, v *vlanproto.VLAN) (*vlanproto.SaveVLANResponse, error) {
	vlan.Save(v.Id, v.Vlan)

	return &vlanproto.SaveVLANResponse{State: vlanproto.State_OK}, nil
}

func (h *handler) GetVLANs(ctx context.Context, _ *emptypb.Empty) (*vlanproto.GetVLANsResponse, error) {
	return &vlanproto.GetVLANsResponse{Vlans: vlan.VLANs()}, nil
}
