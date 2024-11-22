# bazel-playground

this is a little playground to better learn how to work with bazel

## setup

#### install bazelisk and bazel

```
# install bazelisk and bazel
brew install bazelisk
bazel
go install github.com/bazelbuild/bazel-gazelle/cmd/gazelle@latest
```

#### setup bazel in repository

```
# setup WORKSPACE with https://github.com/bazelbuild/bazel-gazelle?tab=readme-ov-file#running-gazelle-with-bazel

# setup BUILD.bazel with gazelle go_prefix
# gazelle:prefix github.com/jimmyl02/bazel-playground
gazelle
```

#### setup vscode with bazel

very helpful guide [here](https://github.com/bazelbuild/rules_go/issues/3014)

create scripts/gopackagesdriver.sh and make it executable

```
#!/bin/bash

exec bazel run -- @io_bazel_rules_go//go/tools/gopackagesdriver "$@"
```

edit the workspace preferences

```
{
    "go.toolsEnvVars": {
        "GOPACKAGESDRIVER": "${workspaceFolder}/scripts/gopackagesdriver.sh"
    }
}
```

## golang

#### import a new dependency with gazelle

add an external dependency

```
# anywhere in the project, run to add to WORKSPACE
gazelle update-repos <package>

# use the dependency in the wanted project
# then run gazelle to update BUILD.bazel:
gazelle
```

add an internal dependency

```
# after adding the dependency in code, there is a "metadata missing" error; anywhere run:
gazelle
```

#### run the cmd

```
bazel run //cmd/testcmd
```

#### build and run the cmd

```
bazel build //cmd/testcmd
./bazel-bin/cmd/testcmd/testcmd_/testcmd
```

## protobuf

following guide [here](https://www.tweag.io/blog/2021-09-08-rules_go-gazelle/)

#### setup

write the proto file into a types directory then run gazelle to generate the `BUILD.bazel`

notice that running `bazel build //...` fails because we are missing `@@com_google_protobuf`

we can add it to our WORKSPACE by following [this](https://github.com/bazelbuild/rules_go/tree/5d306c433cebb1ae8a7b72df2a055be2bacbb12b?tab=readme-ov-file#protobuf-and-grpc)

```
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "com_google_protobuf",
    sha256 = "535fbf566d372ccf3a097c374b26896fa044bf4232aef9cab37bd1cc1ba4e850",
    strip_prefix = "protobuf-3.15.0",
    urls = [
        "https://mirror.bazel.build/github.com/protocolbuffers/protobuf/archive/v3.15.0.tar.gz",
        "https://github.com/protocolbuffers/protobuf/archive/v3.15.0.tar.gz",
    ],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()
```

there is another step after rules_go was updated to 48.0, we have to add rules_proto to the MODULE.bazel file
`bazel_dep(name = "rules_proto", version = "7.0.2")`

now running `bazel build //...` works!

#### using protobuf types within golang

take a look at the `BUILD.bazel` of the types directory; notice that we export a `go_library`, this means we can just directly use it within the go code!
