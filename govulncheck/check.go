// Package govulncheck exposes golang.org/x/vuln/scan as library.
package govulncheck

import (
	"context"

	"golang.org/x/vuln/scan"
)

// Run executes golang.org/x/vuln/scan with given args.
func Run(ctx context.Context, args ...string) error {
	cmd := scan.Command(ctx, args...)
	err := cmd.Start()
	if err == nil {
		err = cmd.Wait()
	}
	return err
}
