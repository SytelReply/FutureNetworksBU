package grpc

import (
	"context"
	"fmt"
	vlanproto "github.com/James-Milligan/FutureNetworksBU/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func Test_HandlerSaveVLAN(t *testing.T) {

	h := Handler{}

	req := &vlanproto.SaveVLANRequest{
		Vlan: &vlanproto.VLAN{
			Id:   "TestIweD",
			Vlan: "TestValue",
		},
	}
	emptyRes := &vlanproto.SaveVLANResponse{}
	res, err := h.SaveVLAN(context.Background(), req)
	if err != nil {
		t.Errorf("Unexpected error from SaveVLAN handler, want: nil, got: %s", err)
	}
	if res.String() != emptyRes.String() {
		t.Errorf("Unexpected response from from SaveVLAN handler, want: nil, got: %s", res)
	}

	res, err = h.SaveVLAN(context.Background(), req)
	if err != nil {
		errStatus, _ := status.FromError(err)
		if errStatus.Code() != codes.InvalidArgument {
			fmt.Println(errStatus.Message())
			t.Errorf("Unexpected error recieved from SaveVLAN handler on duplicated request, want: code 3, got: %s", err)
		}
	}
	if res.String() != emptyRes.String() {
		t.Errorf("Unexpected response from from SaveVLAN handler on duplicated request, want: nil, got: %s", res)
	}
}
