package golang

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Default values used by mageutil/golang/target.Go targets.
const (
	IntegrationTestPkg      = "./integrationtests"
	TestBinDir              = "./target/tests/bin/"
	UnitTestCoverDir        = "./target/tests/cover/unit/"
	IntegrationTestCoverDir = "./target/tests/cover/int/"
	CombinedCoverProfile    = "./target/tests/cover/combined/cover.txt"
)

// IntegrationTestRunner executes integration tests in 4 phases:
//
//  1. Build application binary with coverage collection support.
//  2. Start application binary in background.
//  3. Execute testFn.
//  4. Send SIGINT to application and wait for it to exit.
//
// For example usage see golang.IntegrationTest function.
func IntegrationTestRunner(ctx context.Context, name, coverDir string, testFn func(ctx context.Context) error, runArgs ...string) error {
	buildInfo, err := BuildForTesting(ctx, name)
	if err != nil {
		return fmt.Errorf("builing application failed: %w", err)
	}

	stop, err := StartAppForIntegrationTests(ctx, buildInfo.BinPath, coverDir, runArgs...)
	if err != nil {
		return fmt.Errorf("starting application failed: %w", err)
	}

	if err := testFn(ctx); err != nil {
		_ = stop()
		return fmt.Errorf("running integration tests failed: %w", err)
	}

	if err := stop(); err != nil {
		return fmt.Errorf("running application failed: %w", err)
	}
	return nil
}

// IntegrationTest executes integration tests in 4 phases:
//
//  1. Build application binary with coverage collection support.
//  2. Start application binary in background.
//  3. Execute integration tests.
//  4. Send SIGINT to application and wait for it to exit.
func IntegrationTest(ctx context.Context, name, testPkg, coverDir string, runArgs ...string) error {
	return IntegrationTestRunner(ctx, name, coverDir, func(ctx context.Context) error {
		return RunIntegrationTests(ctx, testPkg)
	}, runArgs...)
}

// UnitTest runs all tests and collects coverage in coverDir.
func UnitTest(ctx context.Context, coverDir string) error {
	err := os.MkdirAll(coverDir, 0755)
	if err != nil {
		return err
	}

	dir, err := filepath.Abs(coverDir)
	if err != nil {
		return err
	}

	env := map[string]string{"CGO_ENABLED": "1"}
	return GoWith(ctx, env, "test", "-race", "-cover", "-covermode", "atomic", "./...", "-test.gocoverdir="+dir)
}

// RunIntegrationTests runs tests inside given package with integration tag.
// To prevent caching -count=1 argument is also provided.
func RunIntegrationTests(ctx context.Context, integrationTestPkg string) error {
	_, err := os.Stat(IntegrationTestPkg)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("No integration tests to run")
		return nil
	}

	return Go(ctx, "test", "-tags=integration", "-count=1", integrationTestPkg)
}

// StartAppForIntegrationTests starts application for integration testing in background.
func StartAppForIntegrationTests(ctx context.Context, bin, coverDir string, args ...string) (stop func() error, err error) {
	err = os.MkdirAll(coverDir, 0755)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command(bin, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = append(cmd.Env, "GOCOVERDIR="+coverDir)

	fmt.Printf("exec: %s %s\n", bin, args)
	if err := cmd.Start(); err != nil {
		return nil, err
	}

	stop = func() error {
		if err := cmd.Process.Signal(os.Interrupt); err != nil {
			return err
		}
		return cmd.Wait()
	}

	return stop, nil
}

// CreateCoverProfile creates combined coverage profile in text format.
func CreateCoverProfile(ctx context.Context, output string, inputDirs ...string) error {
	err := os.MkdirAll(filepath.Dir(output), 0755)
	if err != nil {
		return err
	}
	return Go(ctx, "tool", "covdata", "textfmt", "-i="+strings.Join(inputDirs, ","), "-o", output)
}
