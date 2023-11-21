package mageutil

import (
	"context"
	"io"

	"github.com/elisasre/mageutil/golicenses"
)

const (
	UNKNOWN = "Unknown"
)

// Deprecated: use sub package.
// LicenseCheck runs github.com/google/go-licenses/licenses for given targets
// and writes toe output into w.
func LicenseCheck(ctx context.Context, w io.Writer, targets ...string) error {
	deprecated()
	return golicenses.Run(ctx, w, targets...)
}
