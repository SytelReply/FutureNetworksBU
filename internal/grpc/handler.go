package grpc

import (
	"code-challenge/internal/vlan"
	vlanproto "code-challenge/protos"
	"context"
	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"
)

type handler struct {
	vlanproto.UnimplementedV1Server
}

func (h *handler) SaveVLAN(ctx context.Context, v *vlanproto.SaveVLANRequest) (*vlanproto.SaveVLANResponse, error) {

	// Check if provided VLAN / Id have previously been stored
	if vlan.IsDuplicate(v.Vlan) {
		return &vlanproto.SaveVLANResponse{}, status.Errorf(codes.InvalidArgument,
			"VLAN and/or Id has already been stored")
	}

	vlan.Save(v.Vlan.Id, v.Vlan.Vlan)

	return &vlanproto.SaveVLANResponse{}, nil
}

func (h *handler) GetVLANs(ctx context.Context, _ *vlanproto.GetVLANsRequest) (*vlanproto.GetVLANsResponse, error) {
	return &vlanproto.GetVLANsResponse{Vlans: vlan.VLANs()}, nil
}

func (h *handler) GetVLAN(ctx context.Context, v *vlanproto.GetVLANRequest) (*vlanproto.GetVLANResponse, error) {
	res := vlan.GetVLAN(v.Id)
	if res == nil {
		return &vlanproto.GetVLANResponse{}, status.Errorf(codes.NotFound,
			"VLAN with specified Id does not exist in the system")
	}
	return &vlanproto.GetVLANResponse{Vlan: res}, nil
}
