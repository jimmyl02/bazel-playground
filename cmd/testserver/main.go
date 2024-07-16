package main

import (
	"context"
	"fmt"
	"net"

	"github.com/jimmyl02/bazel-playground/types"
	_ "github.com/jimmyl02/bazel-playground/types"
	"google.golang.org/grpc"
)

type server struct {
	types.UnimplementedTestServer
}

func (s *server) SayHi(ctx context.Context, req *types.SayHiRequest) (*types.SayHiResponse, error) {
	fmt.Println("saying hi!")

	username := ""
	if req.Name != nil {
		username = req.Name.Value
	}

	return &types.SayHiResponse{
		Response: fmt.Sprintf("hi! %s", username),
	}, nil
}

func main() {
	s := grpc.NewServer()
	types.RegisterTestServer(s, &server{})

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	fmt.Println("test srv running!")
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
