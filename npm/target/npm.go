// Package target exposes npm targets that can be imported in magefile using [import syntax].
// For more low level usage the npm package should be preferred.
//
// [import syntax]: https://magefile.org/importing/
package target

import (
	"context"

	"github.com/elisasre/mageutil/npm"
	"github.com/magefile/mage/mg"
)

type Npm mg.Namespace

var NpmCmd = npm.NewCmd()

// Audit runs npm's security audit tool
func (Npm) Audit(ctx context.Context) error {
	return NpmCmd.Audit(ctx)
}

// Install installs npm dependencies
func (Npm) Install(ctx context.Context) error {
	return NpmCmd.Install(ctx)
}

// CleanInstall does clean install on npm dependencies
func (Npm) CleanInstall(ctx context.Context) error {
	return NpmCmd.CleanInstall(ctx)
}

// Update updates npm dependencies
func (Npm) Update(ctx context.Context) error {
	return NpmCmd.Update(ctx)
}

// CleanBuild executes clean-install before build script
func (Npm) CleanBuild(ctx context.Context) {
	mg.SerialCtxDeps(ctx, Npm.CleanInstall, Npm.Build)
}

// Build executes build script
func (Npm) Build(ctx context.Context) error {
	return NpmCmd.Run(ctx, "build")
}

// Start executes start script
func (Npm) Start(ctx context.Context) error {
	return NpmCmd.Run(ctx, "start")
}

// TypeCheck executes typecheck script
func (Npm) TypeCheck(ctx context.Context) error {
	return NpmCmd.Run(ctx, "typecheck")
}

// Lint executes lint script
func (Npm) Lint(ctx context.Context) error {
	return NpmCmd.Run(ctx, "lint")
}

// Lint executes test script
func (Npm) Test(ctx context.Context) error {
	return NpmCmd.Run(ctx, "test")
}
