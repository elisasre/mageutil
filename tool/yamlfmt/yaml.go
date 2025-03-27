// Package yamlfmt exposes yamlfmt commands as targets that can be imported in magefile using [import syntax].
//
// [import syntax]: https://magefile.org/importing/
package yamlfmt

import (
	"context"

	"github.com/elisasre/mageutil/tool"
	"github.com/magefile/mage/mg"
)

const ToolName = "github.com/google/yamlfmt/cmd/yamlfmt"

var YamlFiles = []string{}

// Fmt yaml files
func Fmt(ctx context.Context) error { return FmtFn.Run(ctx) }

// Lint yaml files
func Lint(ctx context.Context) error { return LintFn.Run(ctx) }

var (
	FmtFn mg.Fn = mg.F(func(ctx context.Context) error {
		return tool.Exec(ctx, ToolName, YamlFiles...)
	})

	LintFn mg.Fn = mg.F(func(ctx context.Context) error {
		args := append([]string{"-lint"}, YamlFiles...)
		return tool.Exec(ctx, ToolName, args...)
	})
)
