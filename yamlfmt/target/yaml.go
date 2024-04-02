// Package target exposes yaml targets that can be imported in magefile using [import syntax].
// When using this package the user has to set target.YamlFiles.
// For more low level usage the yamlfmt package should be preferred.
//
// [import syntax]: https://magefile.org/importing/
package target

import (
	"context"

	"github.com/elisasre/mageutil/yamlfmt"
	"github.com/magefile/mage/mg"
)

type Yaml mg.Namespace

// Fmt formats yaml files
func (Yaml) Fmt(ctx context.Context) error { return FmtFn.Run(ctx) }

// Lint lints yaml files
func (Yaml) Lint(ctx context.Context) error { return LintFn.Run(ctx) }

var YamlFiles = []string{}

var (
	FmtFn mg.Fn = mg.F(func(ctx context.Context) error {
		return yamlfmt.Fmt(ctx, YamlFiles...)
	})

	LintFn mg.Fn = mg.F(func(ctx context.Context) error {
		return yamlfmt.Lint(ctx, YamlFiles...)
	})
)
