package mageutil

import (
	"context"

	"github.com/magefile/mage/sh"
)

type NpmCmd func(ctx context.Context, args ...string) error

var defaultNpmCmd = NewNpmCmd()

// Deprecated: use sub package.
func NewNpmCmd(args ...string) NpmCmd {
	deprecated()
	return func(ctx context.Context, newArgs ...string) error {
		return Npm(ctx, append(args, newArgs...)...)
	}
}

// Deprecated: use sub package.
func (npm NpmCmd) Npm(ctx context.Context, args ...string) error {
	deprecated()
	return npm(ctx, args...)
}

// Deprecated: use sub package.
func (npm NpmCmd) Audit(ctx context.Context) error {
	deprecated()
	return npm(ctx, "audit")
}

// Deprecated: use sub package.
func (npm NpmCmd) Install(ctx context.Context) error {
	deprecated()
	return npm(ctx, "install")
}

// Deprecated: use sub package.
func (npm NpmCmd) CleanInstall(ctx context.Context) error {
	deprecated()
	return npm(ctx, "clean-install")
}

// Deprecated: use sub package.
func (npm NpmCmd) Update(ctx context.Context) error {
	deprecated()
	return npm(ctx, "update")
}

// Deprecated: use sub package.
func (npm NpmCmd) Run(ctx context.Context, args ...string) error {
	deprecated()
	return npm(ctx, append([]string{"run"}, args...)...)
}

// Deprecated: use sub package.
func Npm(ctx context.Context, args ...string) error {
	deprecated()
	return NpmWith(ctx, nil, args...)
}

// Deprecated: use sub package.
func NpmWith(_ context.Context, env map[string]string, args ...string) error {
	deprecated()
	return sh.RunWithV(env, "npm", args...)
}

// Deprecated: use sub package.
func NpmInstall(ctx context.Context) error {
	deprecated()
	return defaultNpmCmd(ctx, "install")
}

// Deprecated: use sub package.
func NpmCleanInstall(ctx context.Context) error {
	deprecated()
	return defaultNpmCmd(ctx, "clean-install")
}

// Deprecated: use sub package.
func NpmUpdate(ctx context.Context) error {
	deprecated()
	return defaultNpmCmd(ctx, "update")
}

// Deprecated: use sub package.
func NpmRun(ctx context.Context, args ...string) error {
	deprecated()
	return defaultNpmCmd(ctx, append([]string{"run"}, args...)...)
}
