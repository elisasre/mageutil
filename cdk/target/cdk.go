// Package target exposes cdk targets that can be imported in magefile using [import syntax].
// For more low level usage the npm package should be preferred.
//
// [import syntax]: https://magefile.org/importing/
package target

import (
	"context"

	"github.com/elisasre/mageutil/npm"
	"github.com/magefile/mage/mg"
)

type CDK mg.Namespace

// Install installs npm dependencies for CDK project
func (CDK) Install(ctx context.Context) error { return InstallFn.Run(ctx) }

// CleanInstall performs clean install for CDK project dependencies
func (CDK) CleanInstall(ctx context.Context) error { return CleanInstallFn.Run(ctx) }

// Deploy runs deploy on CDK project to production environment
func (CDK) Deploy(ctx context.Context) error { return DeployFn.Run(ctx) }

// Diff runs projen diff on CDK project
func (CDK) Diff(ctx context.Context) error { return DiffFn.Run(ctx) }

// Build runs build for CDK project
func (CDK) Build(ctx context.Context) error { return BuildFn.Run(ctx) }

// Synth runs projen synth on CDK project
func (CDK) Synth(ctx context.Context) error { return SynthFn.Run(ctx) }

// Update runs package upgrades on CDK project
func (CDK) Update(ctx context.Context) error { return UpdateFn.Run(ctx) }

var NpmCmd = npm.NewCmd("manifests/cdk")

var (
	InstallFn mg.Fn = mg.F(func(ctx context.Context) error {
		return NpmCmd.Install(ctx)
	})

	CleanInstallFn mg.Fn = mg.F(func(ctx context.Context) error {
		return NpmCmd.CleanInstall(ctx)
	})

	DeployFn mg.Fn = mg.F(func(ctx context.Context) error {
		return NpmCmd.Run(ctx, "deploy-all-production")
	})

	DiffFn mg.Fn = mg.F(func(ctx context.Context) error {
		return NpmCmd.Run(ctx, "diff-all-production")
	})

	BuildFn mg.Fn = mg.F(func(ctx context.Context) error {
		return NpmCmd.Run(ctx, "build")
	})

	SynthFn mg.Fn = mg.F(func(ctx context.Context) error {
		return NpmCmd.Run(ctx, "synth")
	})

	UpdateFn mg.Fn = mg.F(func(ctx context.Context) error {
		return NpmCmd.Run(ctx, "projen", "upgrade")
	})
)
