// Package npm exposes npm functionality as library.
package npm

import (
	"context"

	"github.com/magefile/mage/sh"
)

type Cmd func(ctx context.Context, args ...string) error

// Npm executes npm with args.
func Npm(ctx context.Context, args ...string) error {
	return NpmWith(ctx, nil, args...)
}

// Npm commands with env.
func NpmWith(_ context.Context, env map[string]string, args ...string) error {
	return sh.RunWithV(env, "npm", args...)
}

// NewCmd creates npm command that runs always with given args.
// This can be useful for example creating command that runs always within given context.
func NewCmd(args ...string) Cmd {
	return func(ctx context.Context, newArgs ...string) error {
		return Npm(ctx, append(args, newArgs...)...)
	}
}

// Npm runs any nmp command.
func (c Cmd) Npm(ctx context.Context, args ...string) error {
	return c(ctx, args...)
}

// Audit runs npm's security audit tool.
func (c Cmd) Audit(ctx context.Context) error {
	return c(ctx, "audit")
}

// Install installs npm dependencies.
func (c Cmd) Install(ctx context.Context) error {
	return c(ctx, "install")
}

// CleanInstall does clean install on npm dependencies.
func (c Cmd) CleanInstall(ctx context.Context) error {
	return c(ctx, "clean-install")
}

// Update updates npm dependencies.
func (c Cmd) Update(ctx context.Context) error {
	return c(ctx, "update")
}

// Run runs command from script object.
func (c Cmd) Run(ctx context.Context, args ...string) error {
	return c(ctx, append([]string{"run"}, args...)...)
}
