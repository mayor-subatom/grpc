
package main

import (
	"log"
	"flag"
	"time"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "grpc/01-hello-world/service_definition"
	"fmt"
	"strconv"
)


func main() {
	port := flag.Int("port", 50052, "tcp port the server will be listening to")
	flag.Parse()

	fmt.Println("Starting contacting server on port ", *port)

	connection , err :=  grpc.Dial("localhost:"+strconv.Itoa(*port),  grpc.WithInsecure()) //We need to set transport setting in grpc.Dial
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer connection.Close()

	client := pb.NewGreeterClient(connection)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	request, err := client.SayHello(ctx, &pb.HelloRequest{Name:"me"})
	if err != nil {
		log.Fatalf("could not call server: %v", err)
	}

	log.Printf("Greeting: %s", request.Message)
}
