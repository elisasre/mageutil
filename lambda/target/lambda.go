// Package target exposes Lambda targets that can be imported in magefile using [import syntax].
// When using this package the user has to set target.BuildTargets.
//
// [import syntax]: https://magefile.org/importing/
package target

import (
	"context"

	"github.com/elisasre/mageutil/lambda"
	"github.com/magefile/mage/mg"
)

var (
	BuildTargets = []string{}
	GOARCH       = "amd64"
)

type Lambda mg.Namespace

// BuildAll builds lambda bootstrap binaries and calculates sha sums for them
func (Lambda) BuildAll(ctx context.Context) error { return BuildAllFn.Run(ctx) }

var BuildAllFn mg.Fn = mg.F(func(ctx context.Context) error {
	return lambda.BuildAll(ctx, BuildTargets, GOARCH)
})
