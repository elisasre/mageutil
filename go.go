// package mageutil provides util functions for [Magefile].
// For usage please refer to [documentation] provided by Magefile.
//
// [Magefile]: https://magefile.org/
// [documentation]: https://magefile.org/importing/
package mageutil

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

const (
	GoVersion = "go1.20.5"
	CmdDir    = "./cmd/"
	TargetDir = "./target/"
	BinDir    = TargetDir + "bin/"
)

// Go is shorthand for go executable provided by system.
func Go(ctx context.Context, args ...string) error {
	return GoWith(ctx, nil, args...)
}

// GoV prints wanted go version.
// Currently this is hardcoded in this library but in in go1.21.0 this could be defined
// by toolchain directive in go.mod file. So once go1.21 is out this has to be revised.
func GoV() {
	fmt.Println(GoVersion)
}

// GoWith is shorthand for go executable provided by system.
func GoWith(ctx context.Context, env map[string]string, args ...string) error {
	fmt.Printf("env: %s\n", env)
	return sh.RunWithV(env, "go", args...)
}

// Targets returns list of main pkgs under utils.CmdDir.
func Targets(ctx context.Context) ([]string, error) {
	entries, err := os.ReadDir(CmdDir)
	if err != nil {
		return nil, err
	}

	targets := make([]string, 0, len(entries))
	for _, e := range entries {
		if e.IsDir() {
			targets = append(targets, e.Name())
		}
	}
	return targets, nil
}

// BuildAll binaries for targets returned by utils.Targets using utils.Build.
func BuildAll(ctx context.Context) error {
	targets, err := Targets(ctx)
	if err != nil {
		return fmt.Errorf("failed to get list of targets: %w", err)
	}

	for _, target := range targets {
		// Building is CPU intensive and already scales to all available cores hence use of the mg.Serial.
		mg.SerialCtxDeps(ctx, mg.F(Build, target))
	}
	return nil
}

// Build binary using settings from system env.
func Build(ctx context.Context, name string) error {
	goos, err := sh.Output("go", "env", "GOOS")
	if err != nil {
		return err
	}

	goarch, err := sh.Output("go", "env", "GOARCH")
	if err != nil {
		return err
	}

	return BuildFor(ctx, goos, goarch, name)
}

// BuildDefault binary using settings from system env.
func BuildFor(ctx context.Context, goos, goarch, name string) error {
	cmdPath := CmdDir + name
	binaryPath := path.Join(BinDir, goos, goarch, name)
	env := map[string]string{
		"GOOS":   goos,
		"GOARCH": goarch,
	}
	return GoWith(ctx, env, "build", "-o", binaryPath, cmdPath)
}

// BuildForLinux builds binary for amd64 based linux systems.
func BuildForLinux(ctx context.Context, name string) error {
	return BuildFor(ctx, "linux", "amd64", name)
}

// BuildForMac builds binary for amd64 based mac systems.
func BuildForMac(ctx context.Context, name string) error {
	return BuildFor(ctx, "darwin", "amd64", name)
}

// BuildForArmMac builds binary for arm64 based mac systems.
func BuildForArmMac(ctx context.Context, name string) error {
	return BuildFor(ctx, "darwin", "arm64", name)
}

// Run builds and executes app binary from default path.
func Run(ctx context.Context, name string, args ...string) error {
	binaryPath := BinDir + name
	return sh.RunV(binaryPath, args...)
}

// UnitTest executes all unit tests with default flags.
func UnitTest(ctx context.Context) error {
	env := map[string]string{"CGO_ENABLED": "1"}
	return GoWith(ctx, env, "test", "-race", "-covermode", "atomic", "-coverprofile=target/reports/unit-test-coverage.out", "./...")
}
