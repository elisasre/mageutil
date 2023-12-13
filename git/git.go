// Package git exposes git command.
package git

import (
	"bytes"
	"context"
	"errors"

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

// Untracked runs git ls-files --exclude-standard --others arg1 arg2...
// If untracked file is detected error is returned.
func Untracked(ctx context.Context, args ...string) error {
	out, err := sh.Output("git", append([]string{"ls-files", "--exclude-standard", "--others"}, args...)...)
	if err != nil {
		return err
	}
	if len(bytes.TrimSpace([]byte(out))) > 0 {
		return errors.New("untracked files found")
	}
	return nil
}
