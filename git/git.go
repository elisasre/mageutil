// Package git exposes git command.
package git

import (
	"context"

	"github.com/magefile/mage/sh"
)

// Git is shorthand for git executable provided by system.
func Git(ctx context.Context, args ...string) error {
	return sh.RunV("git", args...)
}

// Clean runs git clean with given args.
func Clean(ctx context.Context, args ...string) error {
	return Git(ctx, append([]string{"clean"}, args...)...)
}

// Diff runs git diff with given args.
func Diff(ctx context.Context, args ...string) error {
	return Git(ctx, append([]string{"diff"}, args...)...)
}

// DiffFiles runs git diff --exit-code -- arg1 arg2...
// If change is detected error is returned.
func DiffFilesWithExit(ctx context.Context, args ...string) error {
	return Diff(ctx, append([]string{"--exit-code", "--"}, args...)...)
}
