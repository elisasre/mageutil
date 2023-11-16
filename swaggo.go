package mageutil

import (
	"context"

	"github.com/elisasre/mageutil/swaggo"
	"github.com/swaggo/swag/gen"
)

// SwaggerDocs generates the swagger docs with sane default config.
// Deprecated: use sub package.
func SwaggerDocs(ctx context.Context, searchDir, apiFile, outputDir string) error {
	deprecated()
	return swaggo.GenerateDocs(ctx, searchDir, apiFile, outputDir)
}

// SwaggerDocsWithConf generates the swagger docs with given config.
// Deprecated: use sub package.
func SwaggerDocsWithConf(ctx context.Context, conf gen.Config) error {
	deprecated()
	return swaggo.GenerateDocsWithConf(ctx, conf)
}
