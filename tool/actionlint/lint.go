// Package actionlint exposes actionlint command as target that can be imported in magefile using [import syntax].
//
// [import syntax]: https://magefile.org/importing/
package actionlint

import (
	"context"

	"github.com/elisasre/mageutil/tool"
	"github.com/magefile/mage/mg"
)

const ToolName = "github.com/rhysd/actionlint/cmd/actionlint"

var Tool = tool.New(ToolName)

// Lint runs actionlint
func Lint(ctx context.Context) error { return LintFn.Run(ctx) }

var LintFn mg.Fn = mg.F(func(ctx context.Context) error {
	return Tool.Exec(ctx)
})
