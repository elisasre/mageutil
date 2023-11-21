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

// Deprecated: use sub package.
func Go(ctx context.Context, args ...string) error {
	deprecated()
	return golang.Go(ctx, args...)
}

// Deprecated: use sub package.
func GoWith(ctx context.Context, env map[string]string, args ...string) error {
	deprecated()
	return golang.GoWith(ctx, env, args...)
}

// Deprecated: use sub package.
func Targets(ctx context.Context) ([]string, error) {
	deprecated()
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

// Deprecated: use sub package.
func BuildAll(ctx context.Context) error {
	deprecated()
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

// Deprecated: use sub package.
func Build(ctx context.Context, name string) error {
	deprecated()
	_, err := BuildWithInfo(ctx, name)
	return err
}

// Deprecated: use sub package.
func BuildWithInfo(ctx context.Context, name string) (BuildInfo, error) {
	deprecated()
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

// Deprecated: use sub package.
func BuildWithSHA(ctx context.Context, goos, goarch, name string) {
	deprecated()
	mg.CtxDeps(ctx, func() error {
		_, err := BuildWithSHAWithInfo(ctx, goos, goarch, name)
		return err
	})
}

// Deprecated: use sub package.
func BuildWithSHAWithInfo(ctx context.Context, goos, goarch, name string) (BuildInfo, error) {
	deprecated()
	info, err := BuildForWithInfo(ctx, goos, goarch, name)
	if err != nil {
		return BuildInfo{}, err
	}

	return info, SHA256Sum(ctx, info.BinPath)
}

// Deprecated: use sub package.
func BuildFor(ctx context.Context, goos, goarch, name string) error {
	deprecated()
	_, err := BuildForWithInfo(ctx, goos, goarch, name)
	return err
}

// Deprecated: use sub package.
func BuildForWithInfo(ctx context.Context, goos, goarch, name string) (BuildInfo, error) {
	deprecated()
	cmdPath := CmdDir + name
	binaryPath := path.Join(TargetDir, "bin", goos, goarch, name)
	env := map[string]string{
		"GOOS":   goos,
		"GOARCH": goarch,
	}

	res := BuildInfo{BinPath: binaryPath, GOOS: goos, GOARCH: goarch}
	return res, GoWith(ctx, env, "build", "-o", binaryPath, cmdPath)
}

// Deprecated: use sub package.
func BuildForLinux(ctx context.Context, name string) {
	deprecated()
	BuildWithSHA(ctx, "linux", "amd64", name)
}

// Deprecated: use sub package.
func BuildForLinuxWithInfo(ctx context.Context, name string) (BuildInfo, error) {
	deprecated()
	return BuildWithSHAWithInfo(ctx, "linux", "amd64", name)
}

// Deprecated: use sub package.
func BuildForMac(ctx context.Context, name string) {
	BuildWithSHA(ctx, "darwin", "amd64", name)
}

// Deprecated: use sub package.
func BuildForMacWithInfo(ctx context.Context, name string) (BuildInfo, error) {
	deprecated()
	return BuildWithSHAWithInfo(ctx, "darwin", "amd64", name)
}

// Deprecated: use sub package.
func BuildForArmMac(ctx context.Context, name string) {
	BuildWithSHA(ctx, "darwin", "arm64", name)
}

// Deprecated: use sub package.
func BuildForArmMacWithInfo(ctx context.Context, name string) (BuildInfo, error) {
	deprecated()
	return BuildWithSHAWithInfo(ctx, "darwin", "arm64", name)
}

// Deprecated: use sub package.
func BuildForWindows(ctx context.Context, name string) {
	BuildWithSHA(ctx, "windows", "amd64", name)
}

// Deprecated: use sub package.
func BuildForWindowsWithInfo(ctx context.Context, name string) (BuildInfo, error) {
	deprecated()
	return BuildWithSHAWithInfo(ctx, "windows", "amd64", name)
}

// Deprecated: use sub package.
func Run(ctx context.Context, name string, args ...string) error {
	deprecated()
	bd, err := BinDir()
	if err != nil {
		return err
	}

	binaryPath := path.Join(bd, name)
	return sh.RunV(binaryPath, args...)
}

// Deprecated: use sub package.
func GoList(ctx context.Context, target string) ([]string, error) {
	deprecated()
	return golang.ListPackages(ctx, target)
}

// Deprecated: use sub package.
func BinDir() (string, error) {
	deprecated()
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

// Deprecated: use sub package.
func Ensure(ctx context.Context) error {
	deprecated()
	return golang.Tidy(ctx)
}

// Deprecated: use sub package.
func EnsureInSync(ctx context.Context) error {
	deprecated()
	return golang.TidyAndVerify(ctx)
}

// Deprecated: use sub package.
func Tidy(ctx context.Context) error {
	deprecated()
	return golang.Tidy(ctx)
}

// Deprecated: use sub package.
func TidyAndVerifyNoChanges(ctx context.Context) error {
	deprecated()
	return golang.TidyAndVerify(ctx)
}
