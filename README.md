# Code challenge

This repo contains 2 entry points.

### 1. GRPC Server ([cmd/grpcserver/main.go](cmd/grpcserver/main.go))

The grpc server exposes two rpcs:

```
rpc SaveVLAN (VLAN) returns (SaveVLANResponse) {}
rpc GetVLANs (google.protobuf.Empty) returns (GetVLANsResponse) {}
```

The first stores the given VLAN in memory in an alphabetically ordered array (by id).
The second retrieves the aforementioned array.

### 2. REST Server ([cmd/restserver/main.go](cmd/restserver/main.go))

The REST server exposes a simple gateway to the GRPC server.

## Task

The code to achieve the above is intentionally written impractically. 

1. Find and refactor any areas that could be improved.
2. Add an rpc to the proto definition that allows the retrieval of a single VLAN by id.
3. Implement the above rpc in our grpc handler & create a gateway REST handler that makes a call to this new rpc (in a similar fashion to the other gateway handlers).
4. Dockerise and make these services configurable (e.g. ports/addresses/timeouts).
5. Create kubernetes config to allow this to be deployed to a cluster.
6. Ensure reasonably production ready (e.g. unit tests).
7. Document how to run the services.
