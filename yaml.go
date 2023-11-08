package mageutil

import (
	"context"

	"github.com/elisasre/mageutil/yamlfmt"
	"github.com/google/yamlfmt/command"
)

// YamlFmt formats yaml files using default config.
// Deprecated: use sub package.
func YamlFmt(ctx context.Context, paths ...string) error {
	return yamlfmt.Fmt(ctx, paths...)
}

// YamlLint lints yaml files using default config.
// Deprecated: use sub package.
func YamlLint(ctx context.Context, paths ...string) error {
	return yamlfmt.Lint(ctx, paths...)
}

// Deprecated: use sub package.
func YamlOperationWithDefaultConfig(ctx context.Context, op command.Operation, paths ...string) error {
	return yamlfmt.Run(ctx, op, paths...)
}
