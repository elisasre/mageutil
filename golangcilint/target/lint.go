// Package target exposes golanci-lint targets that can be imported in magefile using [import syntax].
//
// [import syntax]: https://magefile.org/importing/
package target

import (
	"context"

	"github.com/elisasre/mageutil/golangcilint"
	"github.com/magefile/mage/mg"
)

type Go mg.Namespace

// Lint runs golangci-lint for all go files.
func (Go) Lint(ctx context.Context) error {
	return golangcilint.Run(ctx, "run", "./...")
}
