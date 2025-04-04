// Package provides gotestsum command which can be used to moneky-patch golang.DefaultTestCmd
package gotestsum

import (
	"context"

	"github.com/elisasre/mageutil/golang"
	"github.com/elisasre/mageutil/tool"
)

const ToolName = "gotest.tools/gotestsum"

var Tool = tool.New(ToolName)

var DefaultTestCmd golang.Cmd = func(ctx context.Context, env map[string]string, args ...string) error {
	return Tool.ExecWith(ctx, env, append([]string{"--"}, args...)...)
}
