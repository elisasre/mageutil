// Package swaggo exposes openapi tools as library with sane defaults.
package swaggo

import (
	"context"

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
