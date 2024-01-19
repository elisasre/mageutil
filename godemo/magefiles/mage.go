//go:build mage

package main

import (
	"os"

	goutil "github.com/elisasre/mageutil/golang"
	"github.com/elisasre/mageutil/npm"
	"github.com/magefile/mage/mg"

	//mage:import
	_ "github.com/elisasre/mageutil/git/target"
	//mage:import
	_ "github.com/elisasre/mageutil/golangcilint/target"
	//mage:import
	_ "github.com/elisasre/mageutil/govulncheck/target"
	//mage:import
	_ "github.com/elisasre/mageutil/golicenses/target"
	//mage:import
	docker "github.com/elisasre/mageutil/docker/target"
	//mage:import
	cdk "github.com/elisasre/mageutil/cdk/target"
	//mage:import
	ui "github.com/elisasre/mageutil/npm/target"
	//mage:import
	yaml "github.com/elisasre/mageutil/yamlfmt/target"
	//mage:import
	swaggo "github.com/elisasre/mageutil/swaggo/target"
	//mage:import
	golang "github.com/elisasre/mageutil/golang/target"
)

const AppName = "godemo"

// Configure imported targets
func init() {
	os.Setenv(mg.VerboseEnv, "1")
	os.Setenv("CGO_ENABLED", "0")

	docker.ImageName = "europe-north1-docker.pkg.dev/sose-sre-5737/sre-public/godemo"
	docker.ProjectUrl = "https://github.com/elisasre/mageutil/tree/main/godemo"
	ui.NpmCmd = npm.NewCmd("--prefix=./ui/")
	cdk.NpmCmd = npm.NewCmd("--prefix=./manifests/cdk/")
	yaml.YamlFiles = []string{"some.yaml"}
	swaggo.SearchDir = "api"
	swaggo.ApiFile = "api.go"
	swaggo.OutputDir = "docs"
	golang.BuildTarget = "./cmd/godemo"
	golang.BuildMatrix = append(golang.BuildMatrix, goutil.BuildPlatform{OS: "windows", Arch: "amd64"})
}
