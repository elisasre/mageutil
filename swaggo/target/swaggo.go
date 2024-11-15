// Package target exposes swaggo targets that can be imported in magefile using [import syntax].
// When using this package the user has to set target.SearchDir, target.ApiFile, target.OutputDir.
// For more low level usage the swaggo package should be preferred.
//
// [import syntax]: https://magefile.org/importing/
package target

import (
	"context"

	"github.com/elisasre/mageutil/swaggo"
	"github.com/magefile/mage/mg"
)

type Docs mg.Namespace

// OpenAPI generates OpenAPI files using swaggo
func (Docs) OpenAPI(ctx context.Context) error { return OpenAPIFn.Run(ctx) }

// OpenAPIAndVerify generates OpenAPI files using swaggo and verifies output against the version control
func (Docs) OpenAPIAndVerify(ctx context.Context) error { return OpenAPIAndVerifyFn.Run(ctx) }

// OpenAPIAndLint generates OpenAPI files using swaggo and lints output against OpenAPI specification ruleset
func (Docs) OpenAPIAndLint(ctx context.Context) error { return OpenAPIAndLintFn.Run(ctx) }

var (
	SearchDir    = ""
	ApiFile      = ""
	OutputDir    = ""
	LintRuleset  = "https://sre-media.csf.elisa.fi/spectral/elisa-oas.yaml"
	LintSeverity = "error"
)

var (
	OpenAPIFn mg.Fn = mg.F(func(ctx context.Context) error {
		return swaggo.GenerateDocs(ctx, SearchDir, ApiFile, OutputDir)
	})

	OpenAPIAndVerifyFn mg.Fn = mg.F(func(ctx context.Context) error {
		return swaggo.GenerateDocsAndVerify(ctx, SearchDir, ApiFile, OutputDir)
	})

	OpenAPIAndLintFn mg.Fn = mg.F(func(ctx context.Context) {
		mg.SerialCtxDeps(ctx, Docs.OpenAPI, mg.F(swaggo.LintDocs, LintSeverity, LintRuleset, OutputDir))
	})
)
