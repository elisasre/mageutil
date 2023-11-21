package mageutil

import (
	"context"

	"github.com/elisasre/mageutil/golangcilint"
	"github.com/elisasre/mageutil/golangcilint/target"
)

// Deprecated: use sub package.
func LintAll(ctx context.Context) error {
	deprecated()
	return target.Go{}.Lint(ctx)
}

// Deprecated: use sub package.
func GolangCILint(ctx context.Context, args ...string) error {
	deprecated()
	return golangcilint.Run(ctx, args...)
}
