package mageutil

import (
	"context"

	"github.com/elisasre/mageutil/git"
	"github.com/elisasre/mageutil/git/target"
)

// Deprecated: use sub package.
func Git(ctx context.Context, args ...string) error {
	deprecated()
	return git.Git(ctx, args...)
}

// Deprecated: use sub package.
func Clean(ctx context.Context) error {
	deprecated()
	return target.Git{}.Clean(ctx)
}
