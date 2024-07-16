package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/jimmyl02/bazel-playground/pkg/helper"
	"github.com/jimmyl02/bazel-playground/types"
	"github.com/moznion/go-optional"
)

func main() {
	fmt.Println("hello world!")

	req := types.SayHiRequest{
		Name: "hi! my name is test",
	}

	data, err := proto.Marshal(&req)
	if err != nil {
		fmt.Println("Error marshaling protobuf message:", err)
		return
	}

	fmt.Println("Serialized protobuf message:", data)

	some := optional.Some(true)
	helpRes := helper.Help()
	fmt.Println(helpRes, some.Unwrap())
}
