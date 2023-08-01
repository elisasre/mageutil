package mageutil

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"sort"
	"strings"

	"golang.org/x/tools/cover"
)

const (
	UnitCoverProfile        = ReportsDir + "unit-test-coverage.out"
	IntegrationCoverProfile = ReportsDir + "integration-test-coverage.out"
	MergedCoverProfile      = ReportsDir + "merged-test-coverage.out"
)

// UnitTest executes all unit tests with default flags.
func UnitTest(ctx context.Context) error {
	err := os.MkdirAll(path.Dir(UnitCoverProfile), 0755)
	if err != nil {
		return err
	}

	env := map[string]string{"CGO_ENABLED": "1"}
	return GoWith(ctx, env, "test", "-race", "-covermode", "atomic", "-coverprofile="+UnitCoverProfile, "./...")
}

// IntegrationTest executes all tests in given pkg with default flags.
func IntegrationTest(ctx context.Context, pkg string) error {
	pkgs, err := GoList(ctx, "./...")
	if err != nil {
		return err
	}

	err = os.MkdirAll(path.Dir(IntegrationCoverProfile), 0755)
	if err != nil {
		return err
	}

	env := map[string]string{"CGO_ENABLED": "1"}
	return GoWith(ctx, env, "test", "-tags=integration", "-race", "-covermode", "atomic", "-coverpkg="+strings.Join(pkgs, ","), "-coverprofile="+IntegrationCoverProfile, pkg)
}

// MergeCover merges multiple go test -cover profiles and writes the combined coverage into out.
//
// coverage merging is adapted from https://github.com/wadey/gocovmerge
func MergeCover(ctx context.Context, coverFiles []string, w io.Writer) error {
	var merged []*cover.Profile
	for _, file := range coverFiles {
		profiles, err := cover.ParseProfiles(file)
		if err != nil {
			return fmt.Errorf("failed to parse profiles: %w", err)
		}

		for _, p := range profiles {
			merged, err = addProfile(merged, p)
			if err != nil {
				return err
			}
		}
	}

	return dumpProfiles(merged, w)
}

// MergeCoverProfiles merges default unit and integration cover profile.
func MergeCoverProfiles(ctx context.Context) error {
	err := os.RemoveAll(MergedCoverProfile)
	if err != nil {
		return err
	}

	f, err := os.Create(MergedCoverProfile)
	if err != nil {
		return err
	}
	defer f.Close()

	return MergeCover(ctx, []string{UnitCoverProfile, IntegrationCoverProfile}, f)
}

// CoverInfo prints function level cover stats from given profile.
func CoverInfo(ctx context.Context, profile string) error {
	return Go(ctx, "tool", "cover", "-func", profile)
}

func addProfile(profiles []*cover.Profile, p *cover.Profile) ([]*cover.Profile, error) {
	i := sort.Search(len(profiles), func(i int) bool { return profiles[i].FileName >= p.FileName })
	if i < len(profiles) && profiles[i].FileName == p.FileName {
		err := mergeProfiles(profiles[i], p)
		return profiles, err
	}

	profiles = append(profiles, nil)
	copy(profiles[i+1:], profiles[i:])
	profiles[i] = p
	return profiles, nil
}

func mergeProfiles(p *cover.Profile, merge *cover.Profile) (err error) {
	if p.Mode != merge.Mode {
		return errors.New("cannot merge profiles with different modes")
	}
	// Since the blocks are sorted, we can keep track of where the last block
	// was inserted and only look at the blocks after that as targets for merge
	startIndex := 0
	for _, b := range merge.Blocks {
		startIndex, err = mergeProfileBlock(p, b, startIndex)
		if err != nil {
			return err
		}
	}
	return nil
}

func mergeProfileBlock(p *cover.Profile, pb cover.ProfileBlock, startIndex int) (int, error) {
	sortFunc := func(i int) bool {
		pi := p.Blocks[i+startIndex]
		return pi.StartLine >= pb.StartLine && (pi.StartLine != pb.StartLine || pi.StartCol >= pb.StartCol)
	}

	i := 0
	if !sortFunc(i) {
		i = sort.Search(len(p.Blocks)-startIndex, sortFunc)
	}

	i += startIndex
	if i < len(p.Blocks) && p.Blocks[i].StartLine == pb.StartLine && p.Blocks[i].StartCol == pb.StartCol {
		if p.Blocks[i].EndLine != pb.EndLine || p.Blocks[i].EndCol != pb.EndCol {
			return 0, fmt.Errorf("OVERLAP MERGE: %v %v %v", p.FileName, p.Blocks[i], pb)
		}

		switch p.Mode {
		case "set":
			p.Blocks[i].Count |= pb.Count
			return i + 1, nil
		case "count", "atomic":
			p.Blocks[i].Count += pb.Count
			return i + 1, nil
		default:
			return 0, fmt.Errorf("unsupported covermode: '%s'", p.Mode)
		}
	}

	if i > 0 {
		pa := p.Blocks[i-1]
		if pa.EndLine >= pb.EndLine && (pa.EndLine != pb.EndLine || pa.EndCol > pb.EndCol) {
			return 0, fmt.Errorf("OVERLAP BEFORE: %v %v %v", p.FileName, pa, pb)
		}
	}

	if i < len(p.Blocks)-1 {
		pa := p.Blocks[i+1]
		if pa.StartLine <= pb.StartLine && (pa.StartLine != pb.StartLine || pa.StartCol < pb.StartCol) {
			return 0, fmt.Errorf("OVERLAP AFTER: %v %v %v", p.FileName, pa, pb)
		}
	}

	p.Blocks = append(p.Blocks, cover.ProfileBlock{})
	copy(p.Blocks[i+1:], p.Blocks[i:])
	p.Blocks[i] = pb
	return i + 1, nil
}

func dumpProfiles(profiles []*cover.Profile, out io.Writer) error {
	if len(profiles) == 0 {
		return nil
	}

	fmt.Fprintf(out, "mode: %s\n", profiles[0].Mode)
	for _, p := range profiles {
		for _, b := range p.Blocks {
			_, err := fmt.Fprintf(out, "%s:%d.%d,%d.%d %d %d\n", p.FileName, b.StartLine, b.StartCol, b.EndLine, b.EndCol, b.NumStmt, b.Count)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
