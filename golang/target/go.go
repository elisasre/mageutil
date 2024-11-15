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

type Go mg.Namespace

// Build builds binary and calculates sha sum for it
func (Go) Build(ctx context.Context) error { return BuildFn.Run(ctx) }

// CrossBuild builds binary for build matrix
func (Go) CrossBuild(ctx context.Context) error { return CrossBuildFn.Run(ctx) }

// TestBuild builds binary with race detection and coverage collections
func (Go) TestBuild(ctx context.Context) error { return TestBuildFn.Run(ctx) }

// E2eBuild builds binary with coverage collections
func (Go) E2eBuild(ctx context.Context) error { return E2eBuildFn.Run(ctx) }

// Run builds binary and executes it
func (Go) Run(ctx context.Context) error { return RunFn.Run(ctx) }

// Test runs unit and integration tests
func (Go) Test(ctx context.Context) error { return TestFn.Run(ctx) }

// UnitTest runs all unit tests
func (Go) UnitTest(ctx context.Context) error { return UnitTestFn.Run(ctx) }

// IntegrationTest runs integration tests
func (Go) IntegrationTest(ctx context.Context) error { return IntegrationTestFn.Run(ctx) }

// CoverProfile converts binary coverage into text format
func (Go) CoverProfile(ctx context.Context) error { return CoverProfileFn.Run(ctx) }

// ViewCoverage opens test coverage in browser
func (Go) ViewCoverage(ctx context.Context) error { return ViewCoverageFn.Run(ctx) }

// TestAndCover runs all tests and opens coverage in browser
func (Go) TestAndCover(ctx context.Context) error { return TestAndCoverFn.Run(ctx) }

// Tidy runs go mod tidy
func (Go) Tidy(ctx context.Context) error { return TidyFn.Run(ctx) }

// TidyAndVerify verifies that go.mod matches imports
func (Go) TidyAndVerify(ctx context.Context) error { return TidyAndVerifyFn.Run(ctx) }

// RegenIntegrationTestArtifacts regenerates the integration test artifacts
func (Go) RegenIntegrationTestArtifacts(ctx context.Context) error {
	return RegenIntegrationTestArtifactsFn.Run(ctx)
}

// IntegrationTestAndValidate runs the integration tests and checks that testdata artifacts are unchanged
func (Go) IntegrationTestAndValidate(ctx context.Context) error {
	return IntegrationTestAndValidateFn.Run(ctx)
}

var (
	BuildTarget    = ""
	ExtraBuildArgs = []string{}
	BuildMatrix    = golang.DefaultBuildMatrix
	RunArgs        = []string{}
	RunEnvs        = map[string]string{}
)

var (
	BuildFn mg.Fn = mg.F(func(ctx context.Context) error {
		_, err := golang.WithSHA(golang.Build(ctx, BuildTarget, ExtraBuildArgs...))
		return err
	})

	CrossBuildFn mg.Fn = mg.F(func(ctx context.Context) error {
		_, err := golang.BuildFromMatrixWithSHA(ctx, nil, BuildMatrix, BuildTarget)
		return err
	})

	TestBuildFn mg.Fn = mg.F(func(ctx context.Context) error {
		_, err := golang.WithSHA(golang.BuildForTesting(ctx, BuildTarget, false, golang.TestBinDir))
		return err
	})

	E2eBuildFn mg.Fn = mg.F(func(ctx context.Context) error {
		_, err := golang.WithSHA(golang.BuildForTesting(ctx, BuildTarget, true, golang.BinDir))
		return err
	})

	RunFn mg.Fn = mg.F(func(ctx context.Context) error {
		info, err := golang.WithSHA(golang.Build(ctx, BuildTarget, ExtraBuildArgs...))
		if err != nil {
			return err
		}
		return sh.RunWithV(RunEnvs, info.BinPath, RunArgs...)
	})

	TestFn mg.Fn = mg.F(func(ctx context.Context) {
		mg.SerialCtxDeps(ctx, Go.UnitTest, Go.IntegrationTest)
	})

	UnitTestFn mg.Fn = mg.F(func(ctx context.Context) error {
		err := golang.UnitTest(ctx, golang.UnitTestCoverDir)
		if err != nil {
			return err
		}
		return golang.CreateCoverProfile(ctx, golang.UnitTestCoverProfile, golang.UnitTestCoverDir)
	})

	IntegrationTestFn mg.Fn = mg.F(func(ctx context.Context) error {
		err := golang.RunIntegrationTests(ctx, golang.IntegrationTestPkg)
		if err != nil {
			return err
		}
		return golang.CreateCoverProfile(ctx, golang.IntegrationTestCoverProfile, golang.IntegrationTestCoverDir)
	})

	CoverProfileFn mg.Fn = mg.F(func(ctx context.Context) error {
		return golang.CreateCoverProfile(ctx, golang.CombinedCoverProfile, golang.IntegrationTestCoverDir, golang.UnitTestCoverDir)
	})

	ViewCoverageFn mg.Fn = mg.F(func(ctx context.Context) error {
		return golang.Go(ctx, "tool", "cover", "-html", golang.CombinedCoverProfile)
	})

	TestAndCoverFn mg.Fn = mg.F(func(ctx context.Context) {
		mg.SerialCtxDeps(ctx, Go.UnitTest, Go.IntegrationTest, Go.CoverProfile, Go.ViewCoverage)
	})

	TidyFn mg.Fn = mg.F(func(ctx context.Context) error {
		return golang.Tidy(ctx)
	})

	TidyAndVerifyFn mg.Fn = mg.F(func(ctx context.Context) error {
		return golang.TidyAndVerify(ctx)
	})

	RegenIntegrationTestArtifactsFn mg.Fn = mg.F(func(ctx context.Context) {
		os.Setenv("OVERRIDE_TEST_DATA", "true")
		mg.SerialCtxDeps(ctx,
			mg.F(sh.Rm, "./integrationtests/testdata"),
			Go.IntegrationTest,
		)
	})

	IntegrationTestAndValidateFn mg.Fn = mg.F(func(ctx context.Context) error {
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
	})
)
