package mageutil

import (
	"context"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/magefile/mage/sh"
)

const (
	OCILabelTitle       = "org.opencontainers.image.title"
	OCILabelURL         = "org.opencontainers.image.url"
	OCILabelVersion     = "org.opencontainers.image.version"
	OCILabelDescription = "org.opencontainers.image.description"
	OCILabelCreated     = "org.opencontainers.image.created"
	OCILabelSource      = "org.opencontainers.image.source"
	OCILabelLicenses    = "org.opencontainers.image.licenses"
	OCILabelAuthors     = "org.opencontainers.image.authors"
	OCILabelVendor      = "org.opencontainers.image.vendor"
	OCILabelRevision    = "org.opencontainers.image.revision"
)

const (
	DefaultPlatform   = "linux/amd64"
	DefaultDockerfile = "Dockerfile"
	DefaultBuildCtx   = "."
	DefaultExtraCtx   = TargetDir + "bin/" + "linux/amd64/"
)

// Docker runs systems docker cmd with given args.
func Docker(ctx context.Context, args ...string) error {
	return sh.RunV("docker", args...)
}

// DockerPushAllTags push all tags for given image.
func DockerPushAllTags(ctx context.Context, imageName string) error {
	return Docker(ctx, "push", "--all-tags", imageName)
}

// DockerBuildDefault build image with sane defaults.
func DockerBuildDefault(ctx context.Context, imageName, url string) error {
	fullTags := DockerTags(imageName)
	extraCtx := map[string]string{"bin": DefaultExtraCtx}
	labels := DefaultLabels(imageName, url, "")
	return DockerBuild(ctx, DefaultPlatform, DefaultDockerfile, DefaultBuildCtx, fullTags, extraCtx, labels)
}

// DockerBuild is a short hand for docker buildx build with saine default flags
func DockerBuild(ctx context.Context, platform, dockerfile, buildCtx string, tags []string, extraCtx, labels map[string]string) error {
	args := []string{"buildx", "build", "--platform", platform, "-f", dockerfile, "--progress", "plain", "--load"}
	for _, tag := range tags {
		args = append(args, "--tag", tag)
	}
	for k, v := range extraCtx {
		args = append(args, "--build-context", k+"="+v)
	}
	for k, v := range labels {
		args = append(args, "--label", k+"="+v)
	}
	args = append(args, buildCtx)

	return Docker(ctx, args...)
}

// DockerTags creates slice of tags usign `tags` variable and DOCKER_IMAGE_TAGS env var.
func DockerTags(imageName string, tags ...string) []string {
	envTag := os.Getenv("DOCKER_IMAGE_TAGS")
	if envTag != "" {
		tags = append(tags, strings.Split(envTag, " ")...)
	}

	// If no tags were provided `snapshot` is used.
	if len(tags) == 0 {
		tags = append(tags, "snapshot")
	}

	fullTags := make([]string, 0, len(tags))
	for _, tag := range tags {
		fullTags = append(fullTags, fmt.Sprintf("%s:%s", imageName, tag))
	}
	return fullTags
}

// DefaultLabels provides labels for Elisa SoSe/SRE organization.
func DefaultLabels(imageName, url, desc string) map[string]string {
	return map[string]string{
		OCILabelTitle:       path.Base(imageName),
		OCILabelURL:         url,
		OCILabelVersion:     "",
		OCILabelDescription: desc,
		OCILabelCreated:     time.Now().String(),
		OCILabelSource:      url,
		OCILabelLicenses:    "",
		OCILabelAuthors:     "SoSe/SRE",
		OCILabelVendor:      "Elisa",
		OCILabelRevision:    "",
	}
}
