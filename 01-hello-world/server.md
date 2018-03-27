## Implementing server (endpoints)
2 Steps process:
 1) implements the service methods (the what to do). 
 The server interface code has been generated when compiling service definition file.
 Implement server by creating a structure implementing these methods.
 2) start a grpc server listening for requests and calling the service methods





#### Implement service method
Unary RPC: client sends a (blocking) request and waits until a response come back (just like a normal function call).



## Starting server
[using the grpc library](https://grpc.io/docs/tutorials/basic/go.html#starting-the-server)



## Run server
compile + run:
```shell
go run src/hello-world/server/main.go
```
