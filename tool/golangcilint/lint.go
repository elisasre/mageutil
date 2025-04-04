// Package golangcilint exposes golanci-lint commands as targets that can be imported in magefile using [import syntax].
//
// [import syntax]: https://magefile.org/importing/
package golangcilint

import (
	"context"

	"github.com/elisasre/mageutil/tool"
	"github.com/magefile/mage/mg"
)

const ToolName = "github.com/golangci/golangci-lint/v2/cmd/golangci-lint"

var Tool = tool.New(ToolName)

// Lint runs golangci-lint for all go files
func Lint(ctx context.Context) error { return LintFn.Run(ctx) }

// LintAndFix runs golangci-lint for all go files with --fix flag
func LintAndFix(ctx context.Context) error { return LintAndFixFn.Run(ctx) }

var (
	LintFn mg.Fn = mg.F(func(ctx context.Context) error {
		return Tool.Exec(ctx, "run", "./...")
	})

	LintAndFixFn mg.Fn = mg.F(func(ctx context.Context) error {
		return Tool.Exec(ctx, "run", "--fix", "./...")
	})
)
