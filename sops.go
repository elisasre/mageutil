package mageutil

import (
	"context"

	"github.com/elisasre/mageutil/sops"
)

// Deprecated: use sub package.
func SopsDecryptFile(ctx context.Context, file string) ([]byte, error) {
	deprecated()
	return sops.DecryptFile(ctx, file)
}

// Deprecated: use sub packages.
func SopsDecryptWithAwsEnv(ctx context.Context, file, profile string) ([]byte, error) {
	deprecated()
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
