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

// Audit runs npm's security audit tool
func (Npm) Audit(ctx context.Context) error { return AuditFn.Run(ctx) }

// Install installs npm dependencies
func (Npm) Install(ctx context.Context) error { return InstallFn.Run(ctx) }

// CleanInstall does clean install on npm dependencies
func (Npm) CleanInstall(ctx context.Context) error { return CleanInstallFn.Run(ctx) }

// Update updates npm dependencies
func (Npm) Update(ctx context.Context) error { return UpdateFn.Run(ctx) }

// CleanBuild executes clean-install before build script
func (Npm) CleanBuild(ctx context.Context) error { return CleanBuildFn.Run(ctx) }

// Build executes build script
func (Npm) Build(ctx context.Context) error { return BuildFn.Run(ctx) }

// Start executes start script
func (Npm) Start(ctx context.Context) error { return StartFn.Run(ctx) }

// TypeCheck executes typecheck script
func (Npm) TypeCheck(ctx context.Context) error { return TypeCheckFn.Run(ctx) }

// Lint executes lint script
func (Npm) Lint(ctx context.Context) error { return LintFn.Run(ctx) }

// Lint executes test script
func (Npm) Test(ctx context.Context) error { return TestFn.Run(ctx) }

var NpmCmd = npm.NewCmd()

var (
	AuditFn mg.Fn = mg.F(func(ctx context.Context) error {
		return NpmCmd.Audit(ctx)
	})

	InstallFn mg.Fn = mg.F(func(ctx context.Context) error {
		return NpmCmd.Install(ctx)
	})

	CleanInstallFn mg.Fn = mg.F(func(ctx context.Context) error {
		return NpmCmd.CleanInstall(ctx)
	})

	UpdateFn mg.Fn = mg.F(func(ctx context.Context) error {
		return NpmCmd.Update(ctx)
	})

	CleanBuildFn mg.Fn = mg.F(func(ctx context.Context) {
		mg.SerialCtxDeps(ctx, Npm.CleanInstall, Npm.Build)
	})

	BuildFn mg.Fn = mg.F(func(ctx context.Context) error {
		return NpmCmd.Run(ctx, "build")
	})

	StartFn mg.Fn = mg.F(func(ctx context.Context) error {
		return NpmCmd.Run(ctx, "start")
	})

	TypeCheckFn mg.Fn = mg.F(func(ctx context.Context) error {
		return NpmCmd.Run(ctx, "typecheck")
	})

	LintFn mg.Fn = mg.F(func(ctx context.Context) error {
		return NpmCmd.Run(ctx, "lint")
	})

	TestFn mg.Fn = mg.F(func(ctx context.Context) error {
		return NpmCmd.Run(ctx, "test")
	})
)
