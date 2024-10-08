// Package swaggo exposes openapi tools as library with sane defaults.
package swaggo

import (
	"context"
	"fmt"
	"os"

	"github.com/elisasre/mageutil/git"
	"github.com/magefile/mage/sh"
	"github.com/swaggo/swag"
	"github.com/swaggo/swag/gen"
)

// GenerateDocs generates the OpenAPI docs with sane default config.
func GenerateDocs(ctx context.Context, searchDir, apiFile, outputDir string) error {
	return GenerateDocsWithConf(ctx, gen.Config{
		SearchDir:          searchDir,
		PropNamingStrategy: swag.CamelCase,
		MainAPIFile:        apiFile,
		OutputDir:          outputDir,
		ParseInternal:      true,
		ParseVendor:        true,
		ParseDependency:    1,
		OutputTypes:        []string{"go", "json", "yaml"},
		ParseDepth:         100,
		OverridesFile:      gen.DefaultOverridesFile,
		ParseGoList:        true,
		LeftTemplateDelim:  "{{",
		RightTemplateDelim: "}}",
		CollectionFormat:   "csv",
	})
}

// GenerateDocsWithConf generates the OpenAPI docs with given config.
func GenerateDocsWithConf(ctx context.Context, conf gen.Config) error {
	return gen.New().Build(&conf)
}

// GenerateAndVerify generates the OpenAPI docs with sane default config and and verifies that there are no changes to output files.
// This is useful in CI/CD pipelines to validate that OpenAPI docs are up to date.
func GenerateDocsAndVerify(ctx context.Context, searchDir, apiFile, outputDir string) error {
	if err := GenerateDocs(ctx, searchDir, apiFile, outputDir); err != nil {
		return err
	}
	if err := git.DiffFilesWithExit(ctx, outputDir); err != nil {
		return fmt.Errorf("%s is not in sync with the version control", outputDir)
	}
	return nil
}

// LintDocs lints the OpenAPI docs against Elisa's OpenAPI specification ruleset.
func LintDocs(ctx context.Context, severity, ruleset, outputDir string) error {
	return sh.RunV(
		"docker",
		"run",
		"--rm",
		"--volume", fmt.Sprintf("%s:/data", os.Getenv("PWD")),
		"stoplight/spectral:latest",
		"lint",
		"--fail-on-unmatched-globs",
		"--ruleset", ruleset,
		"--format", "text",
		"--fail-severity", severity,
		"--verbose",
		fmt.Sprintf("/data/%s/swagger.yaml", outputDir),
	)
}
