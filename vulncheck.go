package mageutil

import (
	"context"

	"github.com/elisasre/mageutil/govulncheck"
	"github.com/elisasre/mageutil/govulncheck/target"
)

// Deprecated: use sub package.
func VulnCheck(ctx context.Context, args ...string) error {
	deprecated()
	return govulncheck.Run(ctx, args...)
}

// Deprecated: use sub package.
func VulnCheckAll(ctx context.Context) error {
	deprecated()
	return target.Go{}.VulnCheck(ctx)
}
