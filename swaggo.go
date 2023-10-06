package mageutil

import (
	"context"

	"github.com/swaggo/swag"
	"github.com/swaggo/swag/gen"
)

// SwaggerDocs generates the swagger docs with sane default config.
func SwaggerDocs(ctx context.Context, searchDir, apiFile, outputDir string) error {
	return SwaggerDocsWithConf(ctx, gen.Config{
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

// SwaggerDocsWithConf generates the swagger docs with given config.
func SwaggerDocsWithConf(ctx context.Context, conf gen.Config) error {
	return gen.New().Build(&conf)
}
