# mageutil

[![Go Reference](https://pkg.go.dev/badge/github.com/elisasre/mageutil.svg)](https://pkg.go.dev/github.com/elisasre/mageutil) ![build](https://github.com/elisasre/mageutil/actions/workflows/main.yml/badge.svg)

Mageutil provides ready made targets and helper functions for [Magefile](https://magefile.org/). For usage please refer to [documentation](https://magefile.org/importing/) provided by Magefile. For autocompletions see [completions](https://github.com/elisasre/mageutil/tree/main/completions).

### tool pkg

Each package under tool uses `go tool` to execute actual commands. This allows using the "native" implementation without maintaining the main package wrapper code in this repo. To import tool commands under mage namespace check [importing options](https://magefile.org/importing/).

#### Migrating from non tool packages

1. change import path in your `magefile.go`
2. add namespace to import comment eg. `//mage:import go`
    - in case of golangci-lint read also [v2 migration guide](https://golangci-lint.run/product/migration-guide/)
3. run added mage commands locally eg. `mage go:lint`
4. run `go mod tidy`

**From:**

```go
import (
	//mage:import
	_ "github.com/elisasre/mageutil/golangcilint/target"
)
```

**To:**

```go
import (
	//mage:import go
	_ "github.com/elisasre/mageutil/tool/golangcilint"
)
```

### Example

Example magefile:

```go
//go:build mage

package main

import (
	"os"

	"github.com/magefile/mage/mg"

	//mage:import
	_ "github.com/elisasre/mageutil/git/target"
	//mage:import go
	_ "github.com/elisasre/mageutil/tool/golangcilint"
	//mage:import
	docker "github.com/elisasre/mageutil/docker/target"
	//mage:import
	golang "github.com/elisasre/mageutil/golang/target"
)

// Configure imported targets
func init() {
	os.Setenv(mg.VerboseEnv, "1")
	os.Setenv("CGO_ENABLED", "0")

	golang.BuildTarget = "./cmd/myapp"
	golang.RunArgs = []string{"--loglevel", "0", "--development", "true"}
	docker.ImageName = "docker.io/myorg/myapp"
	docker.ProjectUrl = "https://github.com/myorg/myapp"
}
```

Output with example magefile:

```sh
$ mage

Targets:
  docker:build          builds docker image
  docker:push           pushes all tags for image
  git:clean             removes all untracked files from workspace
  go:build              build binary and calculate sha sum for it
  go:coverProfile       convert binary coverage into text format
  go:crossBuild         build binary for build matrix
  go:integrationTest    run integration tests
  go:lint               runs golangci-lint for all go files
  go:lintAndFix         runs golangci-lint for all go files with --fix flag
  go:run                build binary and execute it
  go:test               run unit and integration tests
  go:tidy               run go mod tidy
  go:tidyAndVerify      verify that go.mod matches imports
  go:unitTest           run all unit tests
  go:viewCoverage       open test coverage in browser
```

## Integration tests

Running `mage go:integrationTest` has couple expectations from test code:

1. Test files must be placed under `./integrationtests`
2. Test must produce coverage files in binary format
3. Coverage files must be placed under `./target/tests/cover/int/`

To comply with these rules library like [this](https://pkg.go.dev/github.com/elisasre/go-common@v1.4.6/integrationtest) could be used.
