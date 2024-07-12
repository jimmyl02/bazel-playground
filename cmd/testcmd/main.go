package main

import (
	"fmt"

	"github.com/jimmyl02/bazel-playground/pkg/helper"
	"github.com/moznion/go-optional"
)

func main() {
	fmt.Println("hello world!")

	some := optional.Some(true)
	helpRes := helper.Help()
	fmt.Println(helpRes, some.Unwrap())
}
