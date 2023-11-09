// package mageutil provides util functions for [Magefile].
// For usage please refer to [documentation] provided by Magefile.
// For autocompletions see [completions].
//
// [Magefile]: https://magefile.org/
// [documentation]: https://magefile.org/importing/
// [completions]: https://github.com/elisasre/mageutil/tree/main/completions
package mageutil

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/elisasre/mageutil/golang"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

const (
	CmdDir     = "./cmd/"
	TargetDir  = "./target/"
	ReportsDir = TargetDir + "reports/"
)

// BuildInfo contains relevant information about produced binary.
type BuildInfo struct {
	BinPath string
	GOOS    string
	GOARCH  string
}

// Go is shorthand for go executable provided by system.
// Deprecated: use sub package.
func Go(ctx context.Context, args ...string) error {
	return golang.Go(ctx, args...)
}

// GoWith is shorthand for go executable provided by system.
// Deprecated: use sub package.
func GoWith(ctx context.Context, env map[string]string, args ...string) error {
	return golang.GoWith(ctx, env, args...)
}

// Targets returns list of main pkgs under CmdDir.
// Deprecated: use sub package.
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
// Deprecated: use sub package.
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
// Deprecated: use sub package.
func Build(ctx context.Context, name string) error {
	_, err := BuildWithInfo(ctx, name)
	return err
}

// BuildWithInfo builds binary using settings from system env and returns additional build information.
// Deprecated: use sub package.
func BuildWithInfo(ctx context.Context, name string) (BuildInfo, error) {
	goos, err := sh.Output("go", "env", "GOOS")
	if err != nil {
		return BuildInfo{}, err
	}

	goarch, err := sh.Output("go", "env", "GOARCH")
	if err != nil {
		return BuildInfo{}, err
	}

	return BuildForWithInfo(ctx, goos, goarch, name)
}

// BuildDefault binary and SHA256 sum using settings from system env.
// Deprecated: use sub package.
func BuildWithSHA(ctx context.Context, goos, goarch, name string) {
	mg.CtxDeps(ctx, func() error {
		_, err := BuildWithSHAWithInfo(ctx, goos, goarch, name)
		return err
	})
}

// BuildDefault binary and SHA256 sum using settings from system env
// Deprecated: use sub package.
func BuildWithSHAWithInfo(ctx context.Context, goos, goarch, name string) (BuildInfo, error) {
	info, err := BuildForWithInfo(ctx, goos, goarch, name)
	if err != nil {
		return BuildInfo{}, err
	}

	return info, SHA256Sum(ctx, info.BinPath)
}

// BuildFor builds binary for wanted architecture using default directory schema for output path.
// Deprecated: use sub package.
func BuildFor(ctx context.Context, goos, goarch, name string) error {
	_, err := BuildForWithInfo(ctx, goos, goarch, name)
	return err
}

// BuildForWithInfo builds binary for wanted architecture using default directory schema for output path
// and returns additional build information.
// Deprecated: use sub package.
func BuildForWithInfo(ctx context.Context, goos, goarch, name string) (BuildInfo, error) {
	cmdPath := CmdDir + name
	binaryPath := path.Join(TargetDir, "bin", goos, goarch, name)
	env := map[string]string{
		"GOOS":   goos,
		"GOARCH": goarch,
	}

	res := BuildInfo{BinPath: binaryPath, GOOS: goos, GOARCH: goarch}
	return res, GoWith(ctx, env, "build", "-o", binaryPath, cmdPath)
}

// BuildForLinux builds binary for amd64 based linux systems.
// Deprecated: use sub package.
func BuildForLinux(ctx context.Context, name string) {
	BuildWithSHA(ctx, "linux", "amd64", name)
}

// BuildForLinuxWithInfo builds binary for amd64 based linux systems and returns additional build information.
// Deprecated: use sub package.
func BuildForLinuxWithInfo(ctx context.Context, name string) (BuildInfo, error) {
	return BuildWithSHAWithInfo(ctx, "linux", "amd64", name)
}

// BuildForMac builds binary for amd64 based mac systems.
func BuildForMac(ctx context.Context, name string) {
	BuildWithSHA(ctx, "darwin", "amd64", name)
}

// BuildForMacWithInfo builds binary for amd64 based Mac systems and returns additional build information.
// Deprecated: use sub package.
func BuildForMacWithInfo(ctx context.Context, name string) (BuildInfo, error) {
	return BuildWithSHAWithInfo(ctx, "darwin", "amd64", name)
}

// BuildForArmMac builds binary for arm64 based mac systems.
func BuildForArmMac(ctx context.Context, name string) {
	BuildWithSHA(ctx, "darwin", "arm64", name)
}

// BuildForArmMacWithInfo builds binary for amd64 based Mac systems and returns additional build information.
// Deprecated: use sub package.
func BuildForArmMacWithInfo(ctx context.Context, name string) (BuildInfo, error) {
	return BuildWithSHAWithInfo(ctx, "darwin", "arm64", name)
}

// BuildForWindows builds binary for amd64 based windows systems.
func BuildForWindows(ctx context.Context, name string) {
	BuildWithSHA(ctx, "windows", "amd64", name)
}

// BuildForWindowsWithInfo builds binary for amd64 based Windows systems and returns additional build information.
// Deprecated: use sub package.
func BuildForWindowsWithInfo(ctx context.Context, name string) (BuildInfo, error) {
	return BuildWithSHAWithInfo(ctx, "windows", "amd64", name)
}

// Run executes app binary from default path.
// Deprecated: use sub package.
func Run(ctx context.Context, name string, args ...string) error {
	bd, err := BinDir()
	if err != nil {
		return err
	}

	binaryPath := path.Join(bd, name)
	return sh.RunV(binaryPath, args...)
}

// GoList lists all packages in given target.
// Deprecated: use sub package.
func GoList(ctx context.Context, target string) ([]string, error) {
	return golang.ListPackages(ctx, target)
}

// BinDir returns path in format of target/bin/{GOOS}/{GOARCH}
// Deprecated: use sub package.
func BinDir() (string, error) {
	goos, err := sh.Output("go", "env", "GOOS")
	if err != nil {
		return "", err
	}

	goarch, err := sh.Output("go", "env", "GOARCH")
	if err != nil {
		return "", err
	}

	return path.Join(TargetDir, "bin", goos, goarch), nil
}

// Ensure runs Tidy checks that all dependencies are up to date
// Deprecated: use Tidy instead
func Ensure(ctx context.Context) error {
	log.Println("WARNING: Ensure is deprecated, use Tidy instead")
	return golang.Tidy(ctx)
}

// EnsureInSync checks that all dependencies are up to date
// useful in CI/CD pipelines to validate that dependencies match go.mod
// Deprecated: use TidyAndVerifyNoChanges instead
func EnsureInSync(ctx context.Context) error {
	log.Println("WARNING: EnsureInSync is deprecated, use TidyAndVerifyNoChanges instead")
	return golang.TidyAndVerify(ctx)
}

// Tidy runs go mod tidy
// Deprecated: use sub package.
func Tidy(ctx context.Context) error {
	return golang.Tidy(ctx)
}

// TidyAndVerifyNoChanges runs go mod tidy and verifies that there are no changes to go.mod or go.sum
// useful in CI/CD pipelines to validate that dependencies match go.mod
// Deprecated: use sub package.
func TidyAndVerifyNoChanges(ctx context.Context) error {
	return golang.TidyAndVerify(ctx)
}
