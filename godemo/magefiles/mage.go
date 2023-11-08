//go:build mage

package main

import (
	"context"
	"fmt"

	"github.com/elisasre/mageutil"
	"github.com/elisasre/mageutil/npm"

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
)

const AppName = "godemo"

func init() {
	docker.ImageName = "quay.io/elisaoyj/sre-godemo"
	docker.ProjectUrl = "https://github.com/elisasre/mageutil/tree/main/godemo"
	ui.NpmCmd = npm.NewCmd("--prefix=./ui/")
	cdk.NpmCmd = npm.NewCmd("--prefix=./manifests/cdk/")
	yaml.YamlFiles = []string{"some.yaml"}
}

// Build binaries for executables under ./cmd
func Build(ctx context.Context) error {
	return mageutil.BuildAll(ctx)
}

// Build for x64 Linux
func BuildForLinux(ctx context.Context) {
	mageutil.BuildForLinux(ctx, AppName)
}

// Build for amd64 MacOS
func BuildForMac(ctx context.Context) {
	mageutil.BuildForMac(ctx, AppName)
}

// Build for amd64 MacOS
func BuildForWindows(ctx context.Context) {
	mageutil.BuildForWindows(ctx, AppName)
}

// Build for arm64 MacOS
func BuildForArmMac(ctx context.Context) {
	mageutil.BuildForArmMac(ctx, AppName)
}

// List all packages in the module
func GoList(ctx context.Context) error {
	packages, err := mageutil.GoList(ctx, "./cmd")
	if err != nil {
		return err
	}
	for _, pkg := range packages {
		fmt.Println(pkg)
	}
	return nil
}

// UnitTest whole repo
func UnitTest(ctx context.Context) error {
	return mageutil.UnitTest(ctx)
}

// IntegrationTest whole repo
func IntegrationTest(ctx context.Context) error {
	return mageutil.IntegrationTest(ctx, "./cmd/"+AppName)
}

func MergeCoverProfiles(ctx context.Context) error {
	return mageutil.MergeCoverProfiles(ctx)
}

// SwaggerDocs generates swagger documentation files
func SwaggerDocs(ctx context.Context) error {
	return mageutil.SwaggerDocs(ctx, "api", "api.go", "docs")
}

// Ensure dependencies
// Deprecated: run for verifying backward compatibility
func Ensure(ctx context.Context) error {
	return mageutil.Ensure(ctx)
}

// Ensure dependencies are in sync (CI)
// Deprecated: run for verifying backward compatibility
func EnsureInSync(ctx context.Context) error {
	return mageutil.EnsureInSync(ctx)
}

// Tidy dependencies
func Tidy(ctx context.Context) error {
	return mageutil.Tidy(ctx)
}

// TidyAndVerifyNoChanges dependencies
func TidyAndVerifyNoChanges(ctx context.Context) error {
	return mageutil.TidyAndVerifyNoChanges(ctx)
}
