package main

import (
	"fmt"

	"github.com/jimmyl02/bazel-playground/pkg/helper"
	"github.com/jimmyl02/bazel-playground/proto/testproto"
	"github.com/moznion/go-optional"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {
	fmt.Println("hello world!")

	// try out the protobuf generated type!
	req := testproto.SayHiRequest{
		Name: &wrapperspb.StringValue{
			Value: "hi! my name is test!",
		},
	}

	data, err := proto.Marshal(&req)
	if err != nil {
		fmt.Println("Error marshaling protobuf message:", err)
		return
	}

	fmt.Println("Serialized protobuf message:", data)

	// unmarshal the data back into a new instance of the message struct
	var parsedReq testproto.SayHiRequest
	if err := proto.Unmarshal(data, &parsedReq); err != nil {
		fmt.Println("Error unmarshaling protobuf message:", err)
		return
	}

	// access the parsed data
	if parsedReq.Name != nil {
		fmt.Println("Parsed Name:", parsedReq.Name.Value)
	} else {
		fmt.Println("Name is not set")
	}

	// attempt invalid unmarshal
	var incorrectReq testproto.SayByeRequest
	if err := proto.Unmarshal(data, &incorrectReq); err != nil {
		fmt.Println("Unexpected error when unmarshaling protobuf message:", err)
	} else {
		// sadly, this still works which means it's imperative that the type is correct
		// this is so sad :sob:
		fmt.Println("Expected success but this is so sad ;-;", incorrectReq.Bye)
	}

	some := optional.Some(true)
	helpRes := helper.Help()
	fmt.Println(helpRes, some.Unwrap())
}
