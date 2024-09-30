package docker

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
	DefaultAuthors    = "DiSe/SRE"
	DefaultBuildCtx   = "."
	DefaultExtraCtx   = "./target/bin/linux/amd64/"
)

// Docker executes systems docker cmd with given args.
func Docker(ctx context.Context, args ...string) error {
	return sh.RunV("docker", args...)
}

// Run executes docker run cmd with given args.
func Run(ctx context.Context, args ...string) error {
	return Docker(ctx, append([]string{"run"}, args...)...)
}

// PushAllTags pushes all tags for given image.
func PushAllTags(ctx context.Context, imageName string) error {
	return Docker(ctx, "push", "--all-tags", imageName)
}

// BuildDefault builds image with sane defaults.
func BuildDefault(ctx context.Context, imageName, url string) error {
	return BuildDefaultWithDockerfile(ctx, imageName, url, DefaultAuthors, DefaultDockerfile)
}

// BuildDefaultWithDockerfile builds image from custom Dockerfile location
func BuildDefaultWithDockerfile(ctx context.Context, imageName, url, authors string, dockerfile string) error {
	fullTags := Tags(imageName)
	extraCtx := map[string]string{"bin": DefaultExtraCtx}
	labels := DefaultLabels(imageName, url, "", authors)
	return Build(ctx, DefaultPlatform, dockerfile, DefaultBuildCtx, fullTags, extraCtx, labels)
}

// Build is a short hand for docker buildx build with sane default flags.
func Build(ctx context.Context, platform, dockerfile, buildCtx string, tags []string, extraCtx, labels map[string]string) error {
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

// Tags creates slice of tags using `tags` variable and DOCKER_IMAGE_TAGS env variable.
func Tags(imageName string, tags ...string) []string {
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

// DefaultLabels provides labels for Elisa organization.
func DefaultLabels(imageName, url, desc, authors string) map[string]string {
	return map[string]string{
		OCILabelTitle:       path.Base(imageName),
		OCILabelURL:         url,
		OCILabelVersion:     "",
		OCILabelDescription: desc,
		OCILabelCreated:     time.Now().String(),
		OCILabelSource:      url,
		OCILabelLicenses:    "",
		OCILabelAuthors:     authors,
		OCILabelVendor:      "Elisa",
		OCILabelRevision:    "",
	}
}
