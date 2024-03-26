package target

import (
	"context"

	"github.com/elisasre/mageutil/cyclonedx"
	"github.com/magefile/mage/mg"
)

type Go mg.Namespace

var (
	SbomTargets []string
	BinRoot     = "target/bin/"
)

// SBOM generates a sbom for all binaries under target/bin
func (Go) SBOM(ctx context.Context) error {
	if len(SbomTargets) == 0 {
		targets, err := cyclonedx.FindTargets(BinRoot)
		if err != nil {
			return err
		}
		SbomTargets = targets
	}

	fns := make([]any, 0, len(SbomTargets))
	for _, target := range SbomTargets {
		fns = append(fns, mg.F(cyclonedx.AnalyzeBinToFile, target, target+".bom.json"))
	}

	mg.CtxDeps(ctx, fns...)
	return nil
}
