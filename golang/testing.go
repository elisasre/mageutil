package golang

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/magefile/mage/mg"
)

// Default values used by mageutil/golang/target.Go targets.
const (
	IntegrationTestPkg          = "./integrationtests"
	TestBinDir                  = "./target/tests/bin/"
	UnitTestCoverDir            = "./target/tests/cover/unit/"
	UnitTestCoverProfile        = "./target/tests/cover/unit/cover.txt"
	IntegrationTestCoverDir     = "./target/tests/cover/int/"
	IntegrationTestCoverProfile = "./target/tests/cover/int/cover.txt"
	CombinedCoverProfile        = "./target/tests/cover/combined/cover.txt"
)

var DefaultTestCmd = GoTest

type Cmd func(ctx context.Context, env map[string]string, args ...string) error

func GoTest(ctx context.Context, env map[string]string, args ...string) error {
	if mg.Verbose() {
		args = append([]string{"-v"}, args...)
	}
	args = append([]string{"test"}, args...)
	return GoWith(ctx, env, args...)
}

// IntegrationTestRunner executes integration tests in 4 phases:
//
//  1. Build application binary with coverage collection support.
//  2. Start application binary in background.
//  3. Execute testFn.
//  4. Send SIGINT to application and wait for it to exit.
//
// For example usage see golang.IntegrationTest function.
//
// NOTE: RunIntegrationTests model is now preferred and used by target since it gives more control to test code.
func IntegrationTestRunner(ctx context.Context, name, coverDir string, testFn func(ctx context.Context) error, runArgs ...string) error {
	buildInfo, err := BuildForTesting(ctx, name, false, TestBinDir)
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
//
// NOTE: RunIntegrationTests model is now preferred and used by target since it gives more control to test code.
func IntegrationTest(ctx context.Context, name, testPkg, coverDir string, runArgs ...string) error {
	return IntegrationTestRunner(ctx, name, coverDir, func(ctx context.Context) error {
		return RunIntegrationTests(ctx, testPkg)
	}, runArgs...)
}

// UnitTest runs all tests and collects coverage in coverDir.
func UnitTest(ctx context.Context, coverDir string) error {
	return UnitTestWithCmd(ctx, coverDir, DefaultTestCmd)
}

// UnitTestUnitTestWithCmd allows setting custom `go test` command.`
func UnitTestWithCmd(ctx context.Context, coverDir string, cmd Cmd) error {
	err := os.MkdirAll(coverDir, 0o755)
	if err != nil {
		return err
	}

	dir, err := filepath.Abs(coverDir)
	if err != nil {
		return err
	}

	args := []string{"-tags=unit", "-race", "-cover", "-covermode", "atomic", "./...", "-test.gocoverdir=" + dir}
	env := map[string]string{"CGO_ENABLED": "1"}
	return cmd(ctx, env, args...)
}

// RunIntegrationTests runs tests inside given package with integration tag.
// To prevent caching -count=1 argument is also provided.
func RunIntegrationTests(ctx context.Context, integrationTestPkg string) error {
	return RunIntegrationTestsWithCmd(ctx, integrationTestPkg, DefaultTestCmd)
}

// RunIntegrationTestsWithCmd allows setting custom `go test` command.
func RunIntegrationTestsWithCmd(ctx context.Context, integrationTestPkg string, cmd Cmd) error {
	_, err := os.Stat(IntegrationTestPkg)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("No integration tests to run")
		return nil
	}

	args := []string{"-tags=integration", "-count=1", integrationTestPkg}
	env := map[string]string{"CGO_ENABLED": "1"}
	return cmd(ctx, env, args...)
}

// StartAppForIntegrationTests starts application for integration testing in background.
func StartAppForIntegrationTests(ctx context.Context, bin, coverDir string, args ...string) (stop func() error, err error) {
	err = os.MkdirAll(coverDir, 0o755)
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
	err := os.MkdirAll(filepath.Dir(output), 0o755)
	if err != nil {
		return err
	}
	return Go(ctx, "tool", "covdata", "textfmt", "-i="+strings.Join(inputDirs, ","), "-o", output)
}
