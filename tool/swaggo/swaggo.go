// Package swaggo exposes swaggo commands as targets that can be imported in magefile using [import syntax].
//
// [import syntax]: https://magefile.org/importing/
package swaggo

import (
	"context"
	"fmt"

	"github.com/elisasre/mageutil/git"
	"github.com/elisasre/mageutil/tool"
	"github.com/magefile/mage/mg"
)

const ToolName = "github.com/swaggo/swag/cmd/swag"

// OpenAPI generates OpenAPI files using swaggo
func OpenAPI(ctx context.Context) error { return OpenAPIFn.Run(ctx) }

// OpenAPIAndVerify generates OpenAPI files using swaggo and verifies output against the version control
func OpenAPIAndVerify(ctx context.Context) error { return OpenAPIAndVerifyFn.Run(ctx) }

var (
	SearchDir = ""
	ApiFile   = ""
	OutputDir = ""
)

var (
	Exec = func(ctx context.Context, args ...string) error {
		return tool.Exec(ctx, ToolName, args...)
	}
	OpenAPIFn mg.Fn = mg.F(func(ctx context.Context) error {
		return Exec(ctx, "init",
			"--parseVendor", "--parseInternal", "--parseDependency",
			"-d", SearchDir,
			"-g", ApiFile,
			"-o", OutputDir,
		)
	})
	OpenAPIAndVerifyFn mg.Fn = mg.F(func(ctx context.Context) error {
		if err := OpenAPI(ctx); err != nil {
			return fmt.Errorf("generate docs: %w", err)
		}
		if err := git.DiffFilesWithExit(ctx, OutputDir); err != nil {
			return fmt.Errorf("%s is not in sync with the version control", OutputDir)
		}
		return nil
	})
)
