package mageutil

import (
	"context"

	"github.com/getsops/sops/v3/decrypt"
)

// SopsDecryptFile decrypts sops file.
func SopsDecryptFile(_ context.Context, file string) ([]byte, error) {
	return decrypt.File(file, "")
}

// SopsDecryptWithAwsEnv uses aws-vault to temporarily set credentials in env for decryption functionality.
func SopsDecryptWithAwsEnv(ctx context.Context, file, profile string) ([]byte, error) {
	var data []byte
	if err := AwsWithEnvCredentials(ctx, profile, func() error {
		v, err := SopsDecryptFile(ctx, file)
		if err != nil {
			return err
		}
		data = v
		return nil
	}); err != nil {
		return nil, err
	}

	return data, nil
}
