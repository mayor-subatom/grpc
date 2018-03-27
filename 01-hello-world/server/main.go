package main


import (
	"log"
	"net"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "grpc/01-hello-world/service_definition"
	"flag"
	"fmt"
)

//implementation of the service methods is done via a struct object
type server struct{}



//adding 1st service method to server struct
func (s server) SayHello(ctx context.Context,  req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message:"Server is saying hello to client " + req.Name}, nil
}



func main() {
	port := flag.Int("port", 50052, "tcp port the server will be listening to")
	flag.Parse()

	fmt.Println("Server listening to port ", *port)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()                  //instantiate a gRPC server
	pb.RegisterGreeterServer(grpcServer, &server{}) //register service implementation with the gRPC server
	err = grpcServer.Serve(listener)                //start listening (blocking wait until the process is killed or Stop() is called)
	if( err != nil) {
		log.Fatalf("failed to serve: %v", err)
	}
}




