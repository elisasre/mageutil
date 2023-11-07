package mageutil

import (
	"context"

	"github.com/elisasre/mageutil/git"
	"github.com/elisasre/mageutil/git/target"
)

// Git is shorthand for git executable provided by system.
// Deprecated: use sub package.
func Git(ctx context.Context, args ...string) error {
	return git.Git(ctx, args...)
}

// Clean removes all files ignored by git.
// Deprecated: use sub package.
func Clean(ctx context.Context) error {
	return target.Git{}.Clean(ctx)
}
