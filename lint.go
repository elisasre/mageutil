package mageutil

import (
	"context"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/golangci/golangci-lint/pkg/commands"
)

// LintAll uses golangci-lint library to lint all go files.
func LintAll(ctx context.Context) error {
	return GolangCILint(ctx, "run", "./...")
}

// LintNative imports golangci-lint and runs it as a helper library.
func GolangCILint(ctx context.Context, args ...string) error {
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

	return commands.NewExecutor(info).Execute()
}
