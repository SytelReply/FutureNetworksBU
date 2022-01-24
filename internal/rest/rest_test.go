package rest

import (
	"bytes"
	"context"
	"encoding/json"
	grpcServer "github.com/James-Milligan/FutureNetworksBU/internal/grpc"
	vlanproto "github.com/James-Milligan/FutureNetworksBU/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	vlanproto.RegisterV1Server(server, &grpcServer.Handler{})

	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func Test_HandlerGetVLAN(t *testing.T) {
	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithInsecure(), grpc.WithContextDialer(bufDialer))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	h := handler{
		grpcClient: vlanproto.NewV1Client(conn),
	}
	handler := http.HandlerFunc(h.postVlan)

	// test empty body
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/vlans", nil)
	if err != nil {
		t.Fatal(err)
	}
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Unexpected response code from from PostVLAN request with no body, want: %d, got: %d", http.StatusBadRequest, status)
	}

	body := postVlanRequest{
		ID:   "testID",
		VLAN: "testValue",
	}
	b, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	// Test happy path
	rr = httptest.NewRecorder()
	req, err = http.NewRequest("POST", "/vlans", bytes.NewBuffer(b))
	if err != nil {
		t.Fatal(err)
	}
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Unexpected response code from from PostVLAN request with valid body, want: %d, got: %d", http.StatusOK, status)
	}

	// test duplicated data
	rr = httptest.NewRecorder()
	req, err = http.NewRequest("POST", "/vlans", bytes.NewBuffer(b))
	if err != nil {
		t.Fatal(err)
	}
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Unexpected response code from from PostVLAN request with repeated body, want: %d, got: %d", http.StatusBadRequest, status)
	}
}
