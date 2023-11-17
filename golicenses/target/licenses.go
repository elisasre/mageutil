// Package target exposes golicenses targets that can be imported in magefile using [import syntax].
//
// [import syntax]: https://magefile.org/importing/
package target

import (
	"context"
	"os"

	"github.com/elisasre/mageutil/golicenses"
	"github.com/magefile/mage/mg"
)

type Go mg.Namespace

// Licenses reports licenses used by dependencies
func (Go) Licenses(ctx context.Context) error {
	return golicenses.Run(ctx, os.Stdout, "./...")
}
