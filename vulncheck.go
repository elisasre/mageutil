package mageutil

import (
	"context"

	"github.com/elisasre/mageutil/govulncheck"
	"github.com/elisasre/mageutil/govulncheck/target"
)

// VulnChek runs golang.org/x/vuln/scan with given args.
// Deprecated: use sub package.
func VulnCheck(ctx context.Context, args ...string) error {
	deprecated()
	return govulncheck.Run(ctx, args...)
}

// VulnChek runs golang.org/x/vuln/scan for all packages.
// Deprecated: use sub package.
func VulnCheckAll(ctx context.Context) error {
	deprecated()
	return target.Go{}.VulnCheck(ctx)
}
