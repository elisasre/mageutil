package mageutil

import (
	"context"

	"github.com/elisasre/mageutil/yamlfmt"
	"github.com/google/yamlfmt/command"
)

// Deprecated: use sub package.
func YamlFmt(ctx context.Context, paths ...string) error {
	deprecated()
	return yamlfmt.Fmt(ctx, paths...)
}

// Deprecated: use sub package.
func YamlLint(ctx context.Context, paths ...string) error {
	deprecated()
	return yamlfmt.Lint(ctx, paths...)
}

// Deprecated: use sub package.
func YamlOperationWithDefaultConfig(ctx context.Context, op command.Operation, paths ...string) error {
	deprecated()
	return yamlfmt.Run(ctx, op, paths...)
}
