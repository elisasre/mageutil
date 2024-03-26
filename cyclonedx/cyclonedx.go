package cyclonedx

import (
	"context"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	cdx "github.com/CycloneDX/cyclonedx-go"
	"github.com/CycloneDX/cyclonedx-gomod/pkg/generate/bin"
	"github.com/CycloneDX/cyclonedx-gomod/pkg/licensedetect/local"
	"github.com/rs/zerolog"
)

// AnalyzeBinToFile generates a BOM for the given binary and, creates new file and writes bom data to it.
func AnalyzeBinToFile(ctx context.Context, binPath, outputPath string) error {
	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return AnalyzeBin(ctx, binPath, f)
}

// AnalyzeBinToStdout generates a BOM for the given binary and writes it to stdout.
func AnalyzeBinToStdout(ctx context.Context, binPath string) error {
	return AnalyzeBin(ctx, binPath, os.Stdout)
}

// AnalyzeBin generates a BOM for the given binary and writes it to the given writer.
func AnalyzeBin(ctx context.Context, binPath string, w io.Writer) error {
	logger := zerolog.New(os.Stderr).Level(zerolog.WarnLevel)
	generator, err := bin.NewGenerator(binPath,
		bin.WithLogger(logger),
		bin.WithLicenseDetector(local.NewDetector(logger)),
		bin.WithIncludeStdlib(true),
	)
	if err != nil {
		return err
	}

	bom, err := generator.Generate()
	if err != nil {
		return err
	}

	return WriteBOM(bom, w)
}

// WriteBOM writes the given bom in json format into writer.
func WriteBOM(bom *cdx.BOM, w io.Writer) error {
	outputVersion := cdx.SpecVersion1_4
	encoder := cdx.NewBOMEncoder(w, cdx.BOMFileFormatJSON)
	encoder.SetPretty(true)

	if err := encoder.EncodeVersion(bom, outputVersion); err != nil {
		return fmt.Errorf("failed to encode sbom: %w", err)
	}

	return nil
}

func FindTargets(basePath string) ([]string, error) {
	targets := []string{}
	err := filepath.WalkDir(basePath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		info, err := d.Info()
		if err != nil {
			return err
		}

		mode := info.Mode()
		if mode.IsRegular() && (mode.Perm()&0111) > 0 {
			targets = append(targets, path)
		}

		return nil
	})

	return targets, err
}
