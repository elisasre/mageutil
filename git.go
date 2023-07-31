package mageutil

import (
	"context"

	"github.com/magefile/mage/sh"
)

// Git is shorthand for git executable provided by system.
func Git(ctx context.Context, args ...string) error {
	return sh.RunV("git", args...)
}

// Clean removes all files ignored by git.
func Clean(ctx context.Context) error {
	return Git(ctx, "clean", "-Xdf")
}
