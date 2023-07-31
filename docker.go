package mageutil

import (
	"context"
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
	DefaultImageTagPrefix = "quay.io/elisaoyj/sre-"
	DefaultPlatform       = "linux/amd64"
	DefaultDockerfile     = "Dockerfile"
	DefaultBuildCtx       = "."
	DefaultExtraCtx       = BinDir + "linux/amd64/"
)

// Docker runs systems docker cmd with given args.
func Docker(ctx context.Context, args ...string) error {
	return sh.RunV("docker", args...)
}

// DockerBuildDefault build image with sane defaults.
func DockerBuildDefault(ctx context.Context, name, url string, tags []string) error {
	imageName := DefaultImageTagPrefix + name
	fullTags := make([]string, 0, len(tags))
	for _, tag := range tags {
		fullTags = append(fullTags, imageName+tag)
	}

	extraCtx := map[string]string{"bin": DefaultExtraCtx}
	labels := DefaultLabels(name, url, "")

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

// DefaultLabels provides labels for Elisa SoSe/SRE organization.
func DefaultLabels(title, url, desc string) map[string]string {
	return map[string]string{
		OCILabelTitle:       title,
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
