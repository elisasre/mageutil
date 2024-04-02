// Package target exposes govulncheck targets that can be imported in magefile using [import syntax].
//
// [import syntax]: https://magefile.org/importing/
package target

import (
	"context"
	"fmt"
	"os"

	"github.com/elisasre/mageutil/govulncheck"
	"github.com/magefile/mage/mg"
)

type Go mg.Namespace

// VulnCheck runs golang.org/x/vuln/scan for all packages
func (Go) VulnCheck(ctx context.Context) error {
	fmt.Fprintf(os.Stderr, "Warning: vuln check integration in mageutil package will be deprecated in future major release!\n")
	fmt.Fprintf(os.Stderr, "Warning: Please migrate to use alternate solutions e.g. GitHub Dependabot Alerts for your projects.\n")
	return govulncheck.Run(ctx, "./...")
}
