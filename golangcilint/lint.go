// Package golangcilint exposes github.com/golangci/golangci-lint as library.

package golangcilint

import (
	"context"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/golangci/golangci-lint/pkg/commands"
)

// Run executes golangci-lint with given args.
func Run(ctx context.Context, args ...string) error {
	oldArgs := make([]string, len(os.Args))
	copy(oldArgs, os.Args)
	os.Args = append([]string{"golangci-lint"}, args...)
	defer func() { os.Args = oldArgs }()

	info := commands.BuildInfo{}
	if buildInfo, available := debug.ReadBuildInfo(); available {
		info.GoVersion = buildInfo.GoVersion
		info.Version = buildInfo.Main.Version
		info.Commit = fmt.Sprintf("(unknown, mod sum: %q)", buildInfo.Main.Sum)
		info.Date = "(unknown)"
	}
	return commands.Execute(info)
}
