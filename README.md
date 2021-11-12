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
2. Dockerise and make these services configurable (e.g. which port).
3. Create kubernetes config to allow this to be deployed to a cluster.

(BONUS - write tests)
