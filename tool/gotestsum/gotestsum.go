// Package provides gotestsum command which can be used to moneky-patch golang.DefaultTestCmd
package gotestsum

import (
	"context"

	"github.com/elisasre/mageutil/tool"
)

const ToolName = "gotest.tools/gotestsum"

var Exec = func(ctx context.Context, args ...string) error {
	return tool.Exec(ctx, ToolName, args...)
}
