package mageutil

import (
	"context"

	"github.com/elisasre/mageutil/docker"
)

const (
	OCILabelTitle       = docker.OCILabelTitle
	OCILabelURL         = docker.OCILabelURL
	OCILabelVersion     = docker.OCILabelVersion
	OCILabelDescription = docker.OCILabelDescription
	OCILabelCreated     = docker.OCILabelCreated
	OCILabelSource      = docker.OCILabelSource
	OCILabelLicenses    = docker.OCILabelLicenses
	OCILabelAuthors     = docker.OCILabelAuthors
	OCILabelVendor      = docker.OCILabelVendor
	OCILabelRevision    = docker.OCILabelRevision
)

const (
	DefaultPlatform   = docker.DefaultPlatform
	DefaultDockerfile = docker.DefaultDockerfile
	DefaultBuildCtx   = docker.DefaultBuildCtx
	DefaultExtraCtx   = docker.DefaultExtraCtx
)

// Docker runs systems docker cmd with given args.
// Deprecated: use sub package.
func Docker(ctx context.Context, args ...string) error {
	deprecated()
	return docker.Docker(ctx, args...)
}

// DockerPushAllTags push all tags for given image.
// Deprecated: use sub package.
func DockerPushAllTags(ctx context.Context, imageName string) error {
	deprecated()
	return docker.PushAllTags(ctx, imageName)
}

// DockerBuildDefault build image with sane defaults.
// Deprecated: use sub package.
func DockerBuildDefault(ctx context.Context, imageName, url string) error {
	deprecated()
	return docker.BuildDefault(ctx, imageName, url)
}

// DockerBuild is a short hand for docker buildx build with saine default flags
// Deprecated: use sub package.
func DockerBuild(ctx context.Context, platform, dockerfile, buildCtx string, tags []string, extraCtx, labels map[string]string) error {
	deprecated()
	return docker.Build(ctx, platform, dockerfile, buildCtx, tags, extraCtx, labels)
}

// DockerTags creates slice of tags usign `tags` variable and DOCKER_IMAGE_TAGS env var.
// Deprecated: use sub package.
func DockerTags(imageName string, tags ...string) []string {
	deprecated()
	return docker.Tags(imageName, tags...)
}

// DefaultLabels provides labels for Elisa SoSe/SRE organization.
// Deprecated: use sub package.
func DefaultLabels(imageName, url, desc string) map[string]string {
	deprecated()
	return docker.DefaultLabels(imageName, url, desc)
}
