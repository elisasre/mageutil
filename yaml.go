package mageutil

import (
	"context"

	"github.com/google/yamlfmt"
	"github.com/google/yamlfmt/command"
	"github.com/google/yamlfmt/formatters/basic"
)

// YamlFmt formats yaml files using default config.
func YamlFmt(ctx context.Context, paths ...string) error {
	return YamlOperationWithDefaultConfig(ctx, command.OperationFormat, paths...)
}

// YamlLint lints yaml files using default config.
func YamlLint(ctx context.Context, paths ...string) error {
	return YamlOperationWithDefaultConfig(ctx, command.OperationLint, paths...)
}

func YamlOperationWithDefaultConfig(_ context.Context, op command.Operation, paths ...string) error {
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
