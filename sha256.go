package mageutil

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

// SHA256Sum calculates sum for single file and stores it in file.
// Output should be compatible with sha256sum program.
func SHA256Sum(ctx context.Context, name string) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	sum := sha256.Sum256(data)
	hexSum := hex.EncodeToString(sum[:])

	sumFile, err := os.Create(name + ".sha256")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = fmt.Fprintf(sumFile, "%s *%s\n", hexSum, name)
	if err != nil {
		return err
	}

	return nil
}
