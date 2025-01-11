package main

import (
	"context"
	"fmt"
	"net"

	"github.com/jimmyl02/bazel-playground/proto/testproto"
	"google.golang.org/grpc"
)

type server struct {
	testproto.UnimplementedTestServer
}

func (s *server) SayHi(ctx context.Context, req *testproto.SayHiRequest) (*testproto.SayHiResponse, error) {
	fmt.Println("saying hi!")

	username := ""
	if req.Name != nil {
		username = req.Name.Value
	}

	return &testproto.SayHiResponse{
		Response: fmt.Sprintf("hi! %s", username),
	}, nil
}

func main() {
	s := grpc.NewServer()
	testproto.RegisterTestServer(s, &server{})

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
