package mageutil

import (
	"context"

	"github.com/magefile/mage/sh"
)

type NpmCmd func(ctx context.Context, args ...string) error

var defaultNpmCmd = NewNpmCmd()

// NewNpmCmd creates npm command that runs always with given args.
// This can be useful for example creating command that runs always within given context.
func NewNpmCmd(args ...string) NpmCmd {
	return func(ctx context.Context, newArgs ...string) error {
		return Npm(ctx, append(args, newArgs...)...)
	}
}

// Npm runs any nmp command.
func (npm NpmCmd) Npm(ctx context.Context, args ...string) error {
	return npm(ctx, args...)
}

// Audit runs npm's security audit tool.
func (npm NpmCmd) Audit(ctx context.Context) error {
	return npm(ctx, "audit")
}

// Install installs npm dependencies.
func (npm NpmCmd) Install(ctx context.Context) error {
	return npm(ctx, "install")
}

// CleanInstall does clean install on npm dependencies.
func (npm NpmCmd) CleanInstall(ctx context.Context) error {
	return npm(ctx, "clean-install")
}

// Update updates npm dependencies.
func (npm NpmCmd) Update(ctx context.Context) error {
	return npm(ctx, "update")
}

// Run runs command from script object.
func (npm NpmCmd) Run(ctx context.Context, args ...string) error {
	return npm(ctx, append([]string{"run"}, args...)...)
}

// Npm executes npm with args.
func Npm(ctx context.Context, args ...string) error {
	return NpmWith(ctx, nil, args...)
}

// Npm commands with env.
func NpmWith(_ context.Context, env map[string]string, args ...string) error {
	return sh.RunWithV(env, "npm", args...)
}

// Install installs npm dependencies.
func NpmInstall(ctx context.Context) error {
	return defaultNpmCmd(ctx, "install")
}

// CleanInstall does clean install on npm dependencies.
func NpmCleanInstall(ctx context.Context) error {
	return defaultNpmCmd(ctx, "clean-install")
}

// Update updates npm dependencies.
func NpmUpdate(ctx context.Context) error {
	return defaultNpmCmd(ctx, "update")
}

// Run runs command from script object.
func NpmRun(ctx context.Context, args ...string) error {
	return defaultNpmCmd(ctx, append([]string{"run"}, args...)...)
}
