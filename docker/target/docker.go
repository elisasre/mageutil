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

var (
	ImageName      = ""
	ProjectUrl     = "" // Used for OCI label.
	ProjectAuthors = docker.DefaultAuthors
	Dockerfile     = docker.DefaultDockerfile
)

// Push pushes all tags for image
func (Docker) Push(ctx context.Context) error {
	return docker.PushAllTags(ctx, ImageName)
}

// Build builds docker image
func (Docker) Build(ctx context.Context) error {
	return docker.BuildDefaultWithDockerfile(ctx, ImageName, ProjectUrl, ProjectAuthors, Dockerfile)
}

// Up start containers in daemon mode
func (Docker) Up(ctx context.Context) error {
	return docker.Docker(ctx, "compose", "up", "-d", "--wait")
}

// Down stops containers in daemon mode
func (Docker) Down(ctx context.Context) error {
	return docker.Docker(ctx, "compose", "down", "-v", "--remove-orphans")
}
