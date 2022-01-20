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

COMPLETE
1. Find and refactor any areas that could be improved.
2. Add an rpc to the proto definition that allows the retrieval of a single VLAN by id.
3. Implement the above rpc in our grpc handler & create a gateway REST handler that makes a call to this new rpc (in a similar fashion to the other gateway handlers).
4. Dockerise and make these services configurable (e.g. ports/addresses/timeouts).
5. Create kubernetes config to allow this to be deployed to a cluster.

7. Document how to run the services.

NOT COMPLETE
6. Ensure reasonably production ready (e.g. unit tests).

## How to run the services

#####to deploy to kubernetes cluster:

```
kubectl create -f net-reply.yaml
```

#####to deploy locally:

```
go run cmd/grpcserver/main.go
```
and in a seperate terminal:
```
go run cmd/restserver/main.go
```

## API docs

```
Add a new VLAN:
POST http://{ip}:8081/vlans
Requires JSON body:
{
    "id": string,
    "vlan": string
}
Neither of these 2 fields may be duplicated
```
```
Fetch all available VLANs:
GET http://{ip}:8081/vlans
Expected Response:
{
    "vlans": {
        "id": string,
        "vlan": string
    }[]
}
If no VLANs are stored, this object will be empty
```
```
Fetch VLAN by ID:
GET http://{ip}:8081/vlans/{id}
Expected Response:
{
    "id": string,
    "vlan": string
}
If the provided VLAN ID does not exist in the system, a 404 will be received
```

## Further improvements:
Change method for storing the data (should be using a volume), at present this data is stored in the memory 
of each net-reply-grpc-container meaning that data is lost when the pod is restarted / duplicated when pod is replicated.

Build unit tests for both the grpc handlers and rest handlers

Configure scaling of the kubernetes deployments
