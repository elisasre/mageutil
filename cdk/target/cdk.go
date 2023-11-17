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

var NpmCmd = npm.NewCmd("manifests/cdk")

// Install installs npm dependencies for CDK project
func (CDK) Install(ctx context.Context) error {
	return NpmCmd.Install(ctx)
}

// CleanInstall performs clean install for CDK project dependencies
func (CDK) CleanInstall(ctx context.Context) error {
	return NpmCmd.CleanInstall(ctx)
}

// Deploy runs deploy on CDK project to production environment
func (CDK) Deploy(ctx context.Context) error {
	return NpmCmd.Run(ctx, "deploy-all-production")
}

// Diff runs projen diff on CDK project
func (CDK) Diff(ctx context.Context) error {
	return NpmCmd.Run(ctx, "diff-all-production")
}

// Build runs build for CDK project
func (CDK) Build(ctx context.Context) error {
	return NpmCmd.Run(ctx, "build")
}

// Synth runs projen synth on CDK project
func (CDK) Synth(ctx context.Context) error {
	return NpmCmd.Run(ctx, "synth")
}

// Update runs package upgrades on CDK project
func (CDK) Update(ctx context.Context) error {
	return NpmCmd.Run(ctx, "projen", "upgrade")
}
