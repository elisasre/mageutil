//go:build mage

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/elisasre/mageutil"
	"github.com/magefile/mage/sh"
)

const (
	AppName   = "godemo"
	ImageName = "quay.io/elisaoyj/sre-godemo"
)

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

// Clean removes all files ignored by git
func Clean(ctx context.Context) error {
	return mageutil.Clean(ctx)
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
