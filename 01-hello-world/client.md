
3) the client has a local object known as stub (aka client) implementing same methods as the service.

   [gRPC take care of sending requests to server and returning the server’s protocol buffer responses]


3a) create a gRPC channel to the server
    import "google.golang.org/grpc"
    conn, err := grpc.Dial(address, grpc.WithInsecure()) //Customize channel behaviour (like switching on and off message compression). Mandatory to set transport setting in grpc.Dial
        if err != nil {
            log.Fatalf("did not connect: %v", err)
        }
        defer conn.Close()

        A gRPC channel provides a connection to a gRPC server on a specified host and port.
        Channed is needed when creating a client stub/client.

3b) create a client/stub with the channel using the compiled service definition proto file.
    import pb "google.golang.org/grpc/examples/helloworld/helloworld"
    client := pb.NewGreeterClient(conn)

3c)Calling service methods
   gRPC-Go, RPCs operate in a blocking/synchronous mode, which means that the RPC call waits for the server to respond, and will either return a response or an error.

 - We also pass a context: lets us change our RPC’s behaviour if necessary ( such as cancelling the call or having it time-out).
 ==>  create context is using context.Background() which returns an empty context (no deadline)
        we derive a new context with its deadline updated
        ctx, cancel := context.WithDeadline(parentContext, time)
            // or
        ctx, cancel := context.WithTimeout(parentContext, duration)
        ==> it returns the cancel() function too ! (The Done method returns a read-only channel that is closed on cancelation. Here’s a simple example for checking if the context is canceled.
                    select {
                            case <-ctx.Done():
                                // ctx is canceled
                                return
                            default:
                                // ctx is not canceled, continue immediately
                            }

                            https://siadat.github.io/post/context

## Run cllient
compile + run:
```shell
go run src/hello-world/client/main.go
```
