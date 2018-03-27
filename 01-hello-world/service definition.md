## write the service definition (.proto file aka protobuf)
    - service methods
    - payload messages



## generate sources for stubs (client-side, server-side) and messages
gRPC provides protocol buffer compiler plugins that generate client-side and server-side code.

To generate the gRPC client and server interfaces from .proto service definition (generates go code using the plugin grpc) :
```shell
protoc --go_out=plugins=grpc:. *.proto
```
Generates go code using the plugin grpc.

it creates :
 - an interface class for the server (stub): type XXXServer interface
 - an interface class for the client (stub) : type XXXClient interface
 - messages becomes struct: HelloRequest, HelloResponse
 - XXXClient struct which encapsulates a connection