// Package target exposes govulncheck targets that can be imported in magefile using [import syntax].
//
// [import syntax]: https://magefile.org/importing/
package target

import (
	"context"

	"github.com/elisasre/mageutil/govulncheck"
	"github.com/magefile/mage/mg"
)

type Go mg.Namespace

// VulnCheck runs golang.org/x/vuln/scan for all packages
func (Go) VulnCheck(ctx context.Context) error {
	return govulncheck.Run(ctx, "./...")
}
