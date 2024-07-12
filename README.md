# bazel-playground

this is a little playground to better learn how to work with bazel

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
        "GOPACKAGESDRIVER": "${workspaceFolder}/tools/editor/gopackagesdriver.sh"
    }
}
```

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
