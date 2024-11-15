// Package target exposes golanci-lint targets that can be imported in magefile using [import syntax].
//
// [import syntax]: https://magefile.org/importing/
package target

import (
	"context"

	"github.com/elisasre/mageutil/golangcilint"
	"github.com/magefile/mage/mg"
)

type Go mg.Namespace

// Lint runs golangci-lint for all go files
func (Go) Lint(ctx context.Context) error { return LintFn.Run(ctx) }

// LintAndFix runs golangci-lint for all go files with --fix flag
func (Go) LintAndFix(ctx context.Context) error { return LintAndFixFn.Run(ctx) }

var (
	LintFn mg.Fn = mg.F(func(ctx context.Context) error {
		return golangcilint.Run(ctx, "run", "./...")
	})

	LintAndFixFn mg.Fn = mg.F(func(ctx context.Context) error {
		return golangcilint.Run(ctx, "run", "--fix", "./...")
	})
)
