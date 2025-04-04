// Package provides gotestsum command which can be used to moneky-patch golang.DefaultTestCmd
package gotestsum

import (
	"github.com/elisasre/mageutil/tool"
)

const ToolName = "gotest.tools/gotestsum"

var Tool = tool.New(ToolName)
