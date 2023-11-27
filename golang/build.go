// package golang provides util functions for managing Go project with Go.
package golang

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/magefile/mage/sh"
)

// BinDir is base directory for build outputs.
const BinDir = "./target/bin/"

// BuildInfo contains relevant information about produced binary.
type BuildInfo struct {
	BinPath string
	GOOS    string
	GOARCH  string
}

type (
	BuildPlatform struct{ OS, Arch string }
	BuildMatrix   []BuildPlatform
)

// DefaultBuildMatrix defines subset of cross-compile targets supported by Go.
var DefaultBuildMatrix = BuildMatrix{
	{OS: "linux", Arch: "amd64"},
	{OS: "linux", Arch: "arm64"},
	{OS: "darwin", Arch: "amd64"},
	{OS: "darwin", Arch: "arm64"},
}

// Build builds binary which is created under ./target/bin/{GOOS}/{GOARCH}/.
func Build(ctx context.Context, target string, buildArgs ...string) (BuildInfo, error) {
	return BuildWith(ctx, nil, target, buildArgs...)
}

// BuildWith injects given env and builds binary which is created under ./target/bin/{GOOS}/{GOARCH}/.
func BuildWith(ctx context.Context, env map[string]string, target string, buildArgs ...string) (BuildInfo, error) {
	return build(ctx, env, BinDir, target, buildArgs...)
}

// BuildForPlatform builds binary for wanted architecture and os.
func BuildForPlatform(ctx context.Context, goos, goarch, target string, buildArgs ...string) (BuildInfo, error) {
	return BuildForPlatformWith(ctx, nil, goos, goarch, target, buildArgs...)
}

// BuildForPlatform injects env, builds binary for wanted architecture and os.
func BuildForPlatformWith(ctx context.Context, env map[string]string, goos, goarch, target string, buildArgs ...string) (BuildInfo, error) {
	if env == nil {
		env = map[string]string{}
	}
	env["GOOS"] = goos
	env["GOARCH"] = goarch
	return BuildWith(ctx, env, target, buildArgs...)
}

// WithSHA is a wrapper for build functions that adds SHA256Sum calculation.
func WithSHA(info BuildInfo, err error) (BuildInfo, error) {
	if err != nil {
		return BuildInfo{}, err
	}

	return info, SHA256Sum(info.BinPath, info.BinPath+".sha256")
}

// BuildForTesting builds binary that is instrumented for coverage collection and race detection.
func BuildForTesting(ctx context.Context, target string, e2e bool, binDir string) (BuildInfo, error) {
	var env map[string]string
	if !e2e {
		env = map[string]string{
			"CGO_ENABLED": "1",
		}
	}

	pkgs, err := ListPackages(ctx, "./...")
	if err != nil {
		return BuildInfo{}, err
	}

	args := []string{"-cover", "-covermode", "atomic", "-coverpkg=" + strings.Join(pkgs, ",")}
	if !e2e {
		args = append([]string{"-race"}, args...)
	}
	return build(ctx, env, binDir, target, args...)
}

// BuildFromMatrixWithSHA is a higher level build utility function doing cross compilation with sha calculation.
func BuildFromMatrixWithSHA(ctx context.Context, env map[string]string, matrix BuildMatrix, target string, buildArgs ...string) ([]BuildInfo, error) {
	info := make([]BuildInfo, 0, len(matrix))
	for _, m := range matrix {
		i, err := WithSHA(BuildForPlatformWith(ctx, env, m.OS, m.Arch, target, buildArgs...))
		if err != nil {
			return nil, err
		}
		info = append(info, i)
	}

	return info, nil
}

// ListPackages lists all packages in given target.
func ListPackages(ctx context.Context, target string) ([]string, error) {
	pkgsRaw, err := sh.Output("go", "list", target)
	if err != nil {
		return nil, err
	}
	pkgs := strings.Split(strings.ReplaceAll(pkgsRaw, "\r\n", ","), "\n")
	return pkgs, nil
}

// SHA256Sum calculates sum for input file and stores it in output file.
// Output should be compatible with sha256sum program.
func SHA256Sum(input, output string) error {
	f, err := os.Open(input)
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	sum := sha256.Sum256(data)
	hexSum := hex.EncodeToString(sum[:])

	sumFile, err := os.Create(output)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = fmt.Fprintf(sumFile, "%s *%s\n", hexSum, input)
	if err != nil {
		return err
	}

	return nil
}

func build(ctx context.Context, env map[string]string, base, target string, buildArgs ...string) (BuildInfo, error) {
	info, err := prepareBuildInfo(env, base, target)
	if err != nil {
		return BuildInfo{}, err
	}

	args := append([]string{"build", "-o", info.BinPath}, append(buildArgs, target)...)
	if err = GoWith(ctx, env, args...); err != nil {
		return BuildInfo{}, err
	}

	return info, nil
}

func prepareBuildInfo(env map[string]string, base, target string) (BuildInfo, error) {
	goos, err := sh.Output("go", "env", "GOOS")
	if err != nil {
		return BuildInfo{}, err
	}

	goarch, err := sh.Output("go", "env", "GOARCH")
	if err != nil {
		return BuildInfo{}, err
	}

	if envOS, ok := env["GOOS"]; ok {
		goos = envOS
	}

	if envArch, ok := env["GOARCH"]; ok {
		goarch = envArch
	}

	parts := strings.Split(strings.TrimSuffix(target, "/"), "/")
	name := parts[len(parts)-1]
	binaryPath := path.Join(base, goos, goarch, name)
	return BuildInfo{BinPath: binaryPath, GOOS: goos, GOARCH: goarch}, nil
}
