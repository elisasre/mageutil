package lambda

import (
	"context"
	"os"
	"path"
	"path/filepath"

	"github.com/elisasre/mageutil/golang"
	"github.com/magefile/mage/sh"
)

// BuildAll builds lambda bootstrap binaries and calculates sha sums for them.
func BuildAll(ctx context.Context, buildTargets []string, goarch string) error {
	for _, target := range buildTargets {
		err := Build(ctx, target, goarch)
		if err != nil {
			return err
		}
	}
	return nil
}

// Build builds lambda bootstrap binary and calculates sha sum for it.
func Build(ctx context.Context, target string, goarch string) error {
	info, err := golang.WithSHA(golang.BuildForPlatform(ctx, "linux", goarch, target))
	if err != nil {
		return err
	}

	src := filepath.Dir(info.BinPath)
	dst := path.Join(src, "lambda", filepath.Base(target))
	if err := os.MkdirAll(dst, 0o755); err != nil {
		return err
	}

	dst = path.Join(dst, "bootstrap")
	err = sh.Copy(dst, info.BinPath)
	if err != nil {
		return err
	}
	return nil
}
