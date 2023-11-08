// Package sops provides sops decrypting functionality.
package sops

import (
	"context"

	"github.com/getsops/sops/v3/decrypt"
)

// DecryptFile decrypts sops file.
func DecryptFile(_ context.Context, file string) ([]byte, error) {
	return decrypt.File(file, "")
}
