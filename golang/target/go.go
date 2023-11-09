// Package target exposes Go targets that can be imported in magefile using [import syntax].
// When using this package the user has to set target.BuildTarget.
// For more low level usage the golang package should be preferred.
//
// [import syntax]: https://magefile.org/importing/
package target

import (
	"context"

	"github.com/elisasre/mageutil/golang"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var (
	BuildTarget            = ""
	BuildMatrix            = golang.DefaultBuildMatrix
	RunArgs                = []string{}
	IntegrationTestRunArgs = []string{}
)

type Go mg.Namespace

// Build build binary and calculate sha sum for it
func (Go) Build(ctx context.Context) error {
	_, err := golang.WithSHA(golang.Build(ctx, BuildTarget))
	return err
}

// CrossBuild build binary for build matrix
func (Go) CrossBuild(ctx context.Context) error {
	_, err := golang.BuildFromMatrixWithSHA(ctx, nil, BuildMatrix, BuildTarget)
	return err
}

// Run build binary and execute it
func (Go) Run(ctx context.Context) error {
	info, err := golang.WithSHA(golang.Build(ctx, BuildTarget))
	if err != nil {
		return err
	}

	return sh.RunV(info.BinPath, RunArgs...)
}

// Test run unit and integration tests
func (Go) Test(ctx context.Context) {
	mg.SerialCtxDeps(ctx, Go.UnitTest, Go.IntegrationTest)
}

// UnitTest run all unit tests
func (Go) UnitTest(ctx context.Context) error {
	return golang.UnitTest(ctx, golang.UnitTestCoverDir)
}

// IntegrationTest run integration tests
func (Go) IntegrationTest(ctx context.Context) error {
	return golang.IntegrationTest(ctx, BuildTarget, golang.IntegrationTestPkg, golang.IntegrationTestCoverDir, IntegrationTestRunArgs...)
}

// CoverProfile convert binary coverage into text format
func (Go) CoverProfile(ctx context.Context) error {
	return golang.CreateCoverProfile(ctx, golang.CombinedCoverProfile, golang.IntegrationTestCoverDir, golang.UnitTestCoverDir)
}

// ViewCoverage open test coverage in browser
func (Go) ViewCoverage(ctx context.Context) error {
	return golang.Go(ctx, "tool", "cover", "-html", golang.CombinedCoverProfile)
}

// Tidy run go mod tidy
func (Go) Tidy(ctx context.Context) error {
	return golang.Tidy(ctx)
}

// TidyAndVerify verify that go.mod matches imports
func (Go) TidyAndVerify(ctx context.Context) error {
	return golang.TidyAndVerify(ctx)
}
