package mageutil

import (
	"context"

	"github.com/magefile/mage/sh"
)

type NpmCmd func(ctx context.Context, args ...string) error

var defaultNpmCmd = NewNpmCmd()

// NewNpmCmd creates npm command that runs always with given args.
// This can be useful for example creating command that runs always within given context.
// Deprecated: use sub package.
func NewNpmCmd(args ...string) NpmCmd {
	deprecated()
	return func(ctx context.Context, newArgs ...string) error {
		return Npm(ctx, append(args, newArgs...)...)
	}
}

// Npm runs any nmp command.
// Deprecated: use sub package.
func (npm NpmCmd) Npm(ctx context.Context, args ...string) error {
	deprecated()
	return npm(ctx, args...)
}

// Audit runs npm's security audit tool.
// Deprecated: use sub package.
func (npm NpmCmd) Audit(ctx context.Context) error {
	deprecated()
	return npm(ctx, "audit")
}

// Install installs npm dependencies.
// Deprecated: use sub package.
func (npm NpmCmd) Install(ctx context.Context) error {
	deprecated()
	return npm(ctx, "install")
}

// CleanInstall does clean install on npm dependencies.
// Deprecated: use sub package.
func (npm NpmCmd) CleanInstall(ctx context.Context) error {
	deprecated()
	return npm(ctx, "clean-install")
}

// Update updates npm dependencies.
// Deprecated: use sub package.
func (npm NpmCmd) Update(ctx context.Context) error {
	deprecated()
	return npm(ctx, "update")
}

// Run runs command from script object.
// Deprecated: use sub package.
func (npm NpmCmd) Run(ctx context.Context, args ...string) error {
	deprecated()
	return npm(ctx, append([]string{"run"}, args...)...)
}

// Npm executes npm with args.
// Deprecated: use sub package.
func Npm(ctx context.Context, args ...string) error {
	deprecated()
	return NpmWith(ctx, nil, args...)
}

// Npm commands with env.
// Deprecated: use sub package.
func NpmWith(_ context.Context, env map[string]string, args ...string) error {
	deprecated()
	return sh.RunWithV(env, "npm", args...)
}

// Install installs npm dependencies.
// Deprecated: use sub package.
func NpmInstall(ctx context.Context) error {
	deprecated()
	return defaultNpmCmd(ctx, "install")
}

// CleanInstall does clean install on npm dependencies.
// Deprecated: use sub package.
func NpmCleanInstall(ctx context.Context) error {
	deprecated()
	return defaultNpmCmd(ctx, "clean-install")
}

// Update updates npm dependencies.
// Deprecated: use sub package.
func NpmUpdate(ctx context.Context) error {
	deprecated()
	return defaultNpmCmd(ctx, "update")
}

// Run runs command from script object.
// Deprecated: use sub package.
func NpmRun(ctx context.Context, args ...string) error {
	deprecated()
	return defaultNpmCmd(ctx, append([]string{"run"}, args...)...)
}
