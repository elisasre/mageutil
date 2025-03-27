// Package target exposes govulncheck command as target that can be imported in magefile using [import syntax].
//
// [import syntax]: https://magefile.org/importing/
package target

import (
	"context"

	"github.com/elisasre/mageutil/tool"
	"github.com/magefile/mage/mg"
)

const ToolName = "golang.org/x/vuln/cmd/govulncheck"

// VulnCheck runs golang.org/x/vuln/scan for all packages
func VulnCheck(ctx context.Context) error { return VulnCheckFn.Run(ctx) }

var VulnCheckFn mg.Fn = mg.F(func(ctx context.Context) error {
	return tool.Exec(ctx, ToolName, "./...")
})
