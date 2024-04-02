//go:build mage

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/elisasre/mageutil"
	goutil "github.com/elisasre/mageutil/golang"
	"github.com/elisasre/mageutil/npm"
	"github.com/magefile/mage/mg"

	//mage:import
	_ "github.com/elisasre/mageutil/git/target"
	//mage:import
	_ "github.com/elisasre/mageutil/golangcilint/target"
	//mage:import
	cdx "github.com/elisasre/mageutil/cyclonedx/target"
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
	//mage:import
	lambda "github.com/elisasre/mageutil/lambda/target"
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
	lambda.BuildTargets = []string{"./cmd/godemo"}

	// Target overwriting
	golang.BuildFn = mg.F(customBuild)

	// Add Pre and Post hooks
	cdx.SBOMFn = mageutil.SerialFns{
		mg.F(preSBOM),
		cdx.SBOMFn,
		mg.F(postSBOM),
	}
}

func preSBOM(context.Context)  { fmt.Println("Running pre hook") }
func postSBOM(context.Context) { fmt.Println("Running post hook") }

// customBuild overwrites default GoBuild target.
func customBuild(ctx context.Context) error {
	fmt.Println("Running custom build")
	info, err := goutil.WithSHA(goutil.Build(ctx, golang.BuildTarget))
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", info)
	return nil
}
