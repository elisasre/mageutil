// Package target exposes Go targets that can be imported in magefile using [import syntax].
// When using this package the user has to set target.BuildTarget.
// For more low level usage the golang package should be preferred.
//
// [import syntax]: https://magefile.org/importing/
package target

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/elisasre/mageutil/golang"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var (
	BuildTarget = ""
	BuildMatrix = golang.DefaultBuildMatrix
	RunArgs     = []string{}
	RunEnvs     = map[string]string{}
)

type Go mg.Namespace

// Build builds binary and calculates sha sum for it
func (Go) Build(ctx context.Context) error {
	_, err := golang.WithSHA(golang.Build(ctx, BuildTarget))
	return err
}

// CrossBuild builds binary for build matrix
func (Go) CrossBuild(ctx context.Context) error {
	_, err := golang.BuildFromMatrixWithSHA(ctx, nil, BuildMatrix, BuildTarget)
	return err
}

// TestBuild builds binary with race detection and coverage collections
func (Go) TestBuild(ctx context.Context) error {
	_, err := golang.WithSHA(golang.BuildForTesting(ctx, BuildTarget, false, golang.TestBinDir))
	return err
}

// E2eBuild builds binary with coverage collections
func (Go) E2eBuild(ctx context.Context) error {
	_, err := golang.WithSHA(golang.BuildForTesting(ctx, BuildTarget, true, golang.BinDir))
	return err
}

// Run builds binary and executes it
func (Go) Run(ctx context.Context) error {
	info, err := golang.WithSHA(golang.Build(ctx, BuildTarget))
	if err != nil {
		return err
	}
	return sh.RunWithV(RunEnvs, info.BinPath, RunArgs...)
}

// Test runs unit and integration tests
func (Go) Test(ctx context.Context) {
	mg.SerialCtxDeps(ctx, Go.UnitTest, Go.IntegrationTest)
}

// UnitTest runs all unit tests
func (Go) UnitTest(ctx context.Context) error {
	err := golang.UnitTest(ctx, golang.UnitTestCoverDir)
	if err != nil {
		return err
	}
	return golang.CreateCoverProfile(ctx, golang.UnitTestCoverProfile, golang.UnitTestCoverDir)
}

// IntegrationTest runs integration tests
func (Go) IntegrationTest(ctx context.Context) error {
	err := golang.RunIntegrationTests(ctx, golang.IntegrationTestPkg)
	if err != nil {
		return err
	}
	return golang.CreateCoverProfile(ctx, golang.IntegrationTestCoverProfile, golang.IntegrationTestCoverDir)
}

// CoverProfile converts binary coverage into text format
func (Go) CoverProfile(ctx context.Context) error {
	return golang.CreateCoverProfile(ctx, golang.CombinedCoverProfile, golang.IntegrationTestCoverDir, golang.UnitTestCoverDir)
}

// ViewCoverage opens test coverage in browser
func (Go) ViewCoverage(ctx context.Context) error {
	return golang.Go(ctx, "tool", "cover", "-html", golang.CombinedCoverProfile)
}

// TestAndCover runs all tests and opens coverage in browser
func (Go) TestAndCover(ctx context.Context) {
	mg.SerialCtxDeps(ctx, Go.UnitTest, Go.IntegrationTest, Go.CoverProfile, Go.ViewCoverage)
}

// Tidy runs go mod tidy
func (Go) Tidy(ctx context.Context) error {
	return golang.Tidy(ctx)
}

// TidyAndVerify verifies that go.mod matches imports
func (Go) TidyAndVerify(ctx context.Context) error {
	return golang.TidyAndVerify(ctx)
}

// RegenIntegrationTestArtifacts regenerates the integration test artifacts
func (Go) RegenIntegrationTestArtifacts(ctx context.Context) {
	os.Setenv("OVERRIDE_TEST_DATA", "true")
	mg.SerialCtxDeps(ctx,
		mg.F(sh.Rm, "./integrationtests/testdata"),
		Go.IntegrationTest,
	)
}

// IntegrationTestAndValidate runs the integration tests and checks that testdata artifacts are unchanged
func (Go) IntegrationTestAndValidate(ctx context.Context) error {
	mg.SerialCtxDeps(ctx, Go.RegenIntegrationTestArtifacts)
	out, err := sh.Output("git", "status", "--porcelain", "--", "./integrationtests/testdata/")
	if err != nil {
		return err
	}

	if len(out) > 0 {
		fmt.Println(out)
		return errors.New("testdata artifacts have changed")
	}

	return nil
}
