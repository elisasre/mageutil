package mageutil

import (
	"context"
	"io"

	"github.com/elisasre/mageutil/golicenses"
)

const (
	UNKNOWN = "Unknown"
)

// LicenseCheck runs github.com/google/go-licenses/licenses for given targets
// and writes toe output into w.
// Deprecated: use sub package.
func LicenseCheck(ctx context.Context, w io.Writer, targets ...string) error {
	return golicenses.Run(ctx, w, targets...)
}
