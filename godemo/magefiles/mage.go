//go:build mage

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/elisasre/mageutil"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"

	//mage:import
	_ "github.com/elisasre/mageutil/git/target"
)

const (
	AppName   = "godemo"
	ImageName = "quay.io/elisaoyj/sre-godemo"
)

type UI mg.Namespace
type CDK mg.Namespace

var (
	uiNpm  = mageutil.NewNpmCmd("--prefix=./ui/")
	cdkNpm = mageutil.NewNpmCmd("--prefix=./manifests/cdk/")
)

// Install installs ui deps
func (UI) Install(ctx context.Context) error {
	return uiNpm.Install(ctx)
}

// CleanInstall performs clean install for ui deps
func (UI) CleanInstall(ctx context.Context) error {
	return uiNpm.CleanInstall(ctx)
}

// Build builds ui
func (UI) Build(ctx context.Context) error {
	return uiNpm.Run(ctx, "build")
}

// Test runs tests for ui
func (UI) Test(ctx context.Context) error {
	return uiNpm.Run(ctx, "test")
}

// TestAndBuild tests and builds ui with clean install
func (UI) TestAndBuild(ctx context.Context) {
	mg.SerialCtxDeps(ctx,
		UI.CleanInstall,
		UI.Test,
		UI.Build,
	)
}

// CleanInstall performs clean install for cdk deps
func (CDK) CleanInstall(ctx context.Context) error {
	return cdkNpm.CleanInstall(ctx)
}

// Install installs cdk deps
func (CDK) Install(ctx context.Context) error {
	return cdkNpm.Install(ctx)
}

// Test runs tests for cdk
func (CDK) Test(ctx context.Context) error {
	return cdkNpm.Run(ctx, "test")
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

// Lint all go files.
func Lint(ctx context.Context) error {
	return mageutil.LintAll(ctx)
}

// VulnCheck all go files.
func VulnCheck(ctx context.Context) error {
	return mageutil.VulnCheckAll(ctx)
}

// LicenseCheck all files.
func LicenseCheck(ctx context.Context) error {
	if err := os.MkdirAll(mageutil.ReportsDir, 0755); err != nil {
		return fmt.Errorf("failed to create reports dir: %w", err)
	}
	licenseFile, err := os.Create(fmt.Sprintf("%s%s", mageutil.ReportsDir, "licenses.csv"))
	if err != nil {
		return err
	}
	return mageutil.LicenseCheck(ctx, licenseFile, mageutil.CmdDir+AppName)
}

// Build docker image
func BuildImage(ctx context.Context) error {
	currentGitCommit, err := sh.Output("git", "rev-parse", "--short", "HEAD")
	if err != nil {
		return err
	}
	os.Setenv("DOCKER_IMAGE_TAGS", currentGitCommit)
	return mageutil.DockerBuildDefault(ctx, ImageName, "")
}

// Push image
func PushImage(ctx context.Context) error {
	return mageutil.DockerPushAllTags(ctx, ImageName)
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

// YamlLint some.yaml file
func YamlLint(ctx context.Context) error {
	return mageutil.YamlLint(ctx, "some.yaml")
}

// YamlFmt some.yaml file
func YamlFmt(ctx context.Context) error {
	return mageutil.YamlFmt(ctx, "some.yaml")
}
