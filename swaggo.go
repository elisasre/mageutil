package mageutil

import (
	"context"

	"github.com/elisasre/mageutil/swaggo"
	"github.com/swaggo/swag/gen"
)

// Deprecated: use sub package.
func SwaggerDocs(ctx context.Context, searchDir, apiFile, outputDir string) error {
	deprecated()
	return swaggo.GenerateDocs(ctx, searchDir, apiFile, outputDir)
}

// Deprecated: use sub package.
func SwaggerDocsWithConf(ctx context.Context, conf gen.Config) error {
	deprecated()
	return swaggo.GenerateDocsWithConf(ctx, conf)
}
