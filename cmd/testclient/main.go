package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jimmyl02/bazel-playground/types"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// use the conn variable to interact with the gRPC server
	c := types.NewTestClient(conn)

	// contact the server and print the response
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHi(ctx, &types.SayHiRequest{
		Name: &wrapperspb.StringValue{
			Value: "my name is jeff",
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("resp", r.Response)

	fmt.Println("ran test client!")
}