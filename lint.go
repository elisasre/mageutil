package mageutil

import (
	"context"

	"github.com/elisasre/mageutil/golangcilint"
	"github.com/elisasre/mageutil/golangcilint/target"
)

// LintAll uses golangci-lint library to lint all go files.
// Deprecated: use sub package.
func LintAll(ctx context.Context) error {
	deprecated()
	return target.Go{}.Lint(ctx)
}

// LintNative imports golangci-lint and runs it as a helper library.
// Deprecated: use sub package.
func GolangCILint(ctx context.Context, args ...string) error {
	deprecated()
	return golangcilint.Run(ctx, args...)
}
