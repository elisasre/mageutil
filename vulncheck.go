package mageutil

import (
	"context"

	"golang.org/x/vuln/scan"
)

// VulnChek runs golang.org/x/vuln/scan with given args.
func VulnCheck(ctx context.Context, args ...string) error {
	cmd := scan.Command(ctx, args...)
	err := cmd.Start()
	if err == nil {
		err = cmd.Wait()
	}
	return err
}

// VulnChek runs golang.org/x/vuln/scan for all packages.
func VulnCheckAll(ctx context.Context) error {
	return VulnCheck(ctx, "./...")
}
