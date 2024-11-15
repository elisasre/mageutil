// Package target exposes git targets that can be imported in magefile using [import syntax].
//
// [import syntax]: https://magefile.org/importing/

package target

import (
	"context"

	"github.com/elisasre/mageutil/git"
	"github.com/magefile/mage/mg"
)

type Git mg.Namespace

// Clean removes all untracked files from workspace
func (Git) Clean(ctx context.Context) error { return CleanFn.Run(ctx) }

var CleanFn mg.Fn = mg.F(func(ctx context.Context) error {
	return git.Clean(ctx, "-Xdf")
})
