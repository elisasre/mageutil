// Package target exposes docker targets that can be imported in magefile using [import syntax].
// To use these targets target.ImageName and target.ImageUrl has to be set
// For projects that build more than one image using the util package should be preferred.
// DOCKER_IMAGE_TAGS environment variable can be used to set wanted tags for commands.
//
// [import syntax]: https://magefile.org/importing/
package target

import (
	"context"

	"github.com/elisasre/mageutil/docker"
	"github.com/magefile/mage/mg"
)

type Docker mg.Namespace

// Push pushes all tags for image
func (Docker) Push(ctx context.Context) error { return PushFn.Run(ctx) }

// Build builds docker image
func (Docker) Build(ctx context.Context) error { return BuildFn.Run(ctx) }

// Up start containers in daemon mode
func (Docker) Up(ctx context.Context) error { return UpFn.Run(ctx) }

// Down stops containers in daemon mode
func (Docker) Down(ctx context.Context) error { return DownFn.Run(ctx) }

// Recreate teardowns and recreates containers in daemon mode
func (Docker) Recreate(ctx context.Context) error { return RecreateFn.Run(ctx) }

var (
	ImageName      = ""
	ProjectUrl     = "" // Used for OCI label.
	ProjectAuthors = docker.DefaultAuthors
	ProjectVendor  = docker.DefaultVendor
	Dockerfile     = docker.DefaultDockerfile
)

var (
	PushFn mg.Fn = mg.F(func(ctx context.Context) error {
		return docker.PushAllTags(ctx, ImageName)
	})

	BuildFn mg.Fn = mg.F(func(ctx context.Context) error {
		return docker.BuildDefaultWithDockerfile(ctx, ImageName, &docker.Labels{
			URL:     ProjectUrl,
			Authors: ProjectAuthors,
			Vendor:  docker.DefaultVendor,
		}, Dockerfile)
	})

	UpFn mg.Fn = mg.F(func(ctx context.Context) error {
		return docker.Docker(ctx, "compose", "up", "-d", "--wait")
	})

	DownFn mg.Fn = mg.F(func(ctx context.Context) error {
		return docker.Docker(ctx, "compose", "down", "-v", "--remove-orphans")
	})

	RecreateFn mg.Fn = mg.F(func(ctx context.Context) {
		mg.SerialCtxDeps(ctx, Docker.Down, Docker.Up)
	})
)
