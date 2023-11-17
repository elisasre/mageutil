package golang

import (
	"context"
	"fmt"

	"github.com/elisasre/mageutil/git"
	"github.com/magefile/mage/sh"
)

// Go is shorthand for go executable provided by system.
func Go(ctx context.Context, args ...string) error {
	return GoWith(ctx, nil, args...)
}

// GoWith is shorthand for go executable provided by system.
func GoWith(ctx context.Context, env map[string]string, args ...string) error {
	fmt.Printf("env: %s\n", env)
	return sh.RunWithV(env, "go", args...)
}

// Tidy runs go mod tidy.
func Tidy(ctx context.Context) error {
	return Go(ctx, "mod", "tidy")
}

// TidyAndVerify runs go mod tidy and verifies that there are no changes to go.mod or go.sum.
// This is useful in CI/CD pipelines to validate that dependencies match go.mod.
func TidyAndVerify(ctx context.Context) error {
	if err := Tidy(ctx); err != nil {
		return err
	}
	if err := git.DiffFilesWithExit(ctx, "go.mod", "go.sum"); err != nil {
		return fmt.Errorf("go.mod and go.sum are not in sync. run `go mod tidy` and commit changes")
	}
	return nil
}
