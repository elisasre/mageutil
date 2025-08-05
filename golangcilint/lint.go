// Package golangcilint exposes github.com/golangci/golangci-lint as library.

package golangcilint

import (
	"context"
	"errors"
)

// Run executes golangci-lint with given args.
func Run(ctx context.Context, args ...string) error {
	return errors.New("deprecated: use tool/golangcilint")
}
