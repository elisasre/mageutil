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

var (
	SearchDir = ""
	ApiFile   = ""
	OutputDir = ""
)

// OpenAPI generates OpenAPI files using swaggo.
func (Docs) OpenAPI(ctx context.Context) error {
	return swaggo.GenerateDocs(ctx, SearchDir, ApiFile, OutputDir)
}
