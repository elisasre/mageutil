// Package yamlfmt exposes yaml formatting functionality as library with sane defaults.
package yamlfmt

import (
	"context"

	"github.com/google/yamlfmt"
	"github.com/google/yamlfmt/command"
	"github.com/google/yamlfmt/formatters/basic"
)

// Fmt formats yaml files using default config.
func Fmt(ctx context.Context, paths ...string) error {
	return Run(ctx, command.OperationFormat, paths...)
}

// Lint lints yaml files using default config.
func Lint(ctx context.Context, paths ...string) error {
	return Run(ctx, command.OperationLint, paths...)
}

func Run(_ context.Context, op command.Operation, paths ...string) error {
	conf := command.NewConfig()
	conf.LineEnding = yamlfmt.LineBreakStyleLF
	conf.Extensions = []string{"yaml", "yml"}
	conf.Include = paths

	c := &command.Command{
		Operation: op,
		Registry:  yamlfmt.NewFormatterRegistry(&basic.BasicFormatterFactory{}),
		Quiet:     false,
		Config:    &conf,
	}

	return c.Run()
}
