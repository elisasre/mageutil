package mageutil

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/elisasre/mageutil/awsvault"
)

// Deprecated: use sub package.
func AwsVault(ctx context.Context, args ...string) error {
	deprecated()
	return awsvault.AwsVault(ctx, args...)
}

// Deprecated: use sub package.
func AwsVaultExec(ctx context.Context, awsProfile string, args ...string) error {
	deprecated()
	return awsvault.AwsVaultExec(ctx, awsProfile, args...)
}

// Deprecated: use sub package.
func AwsWithEnvCredentials(ctx context.Context, awsProfile string, fn func() error) error {
	deprecated()
	return awsvault.AwsWithEnvCredentials(ctx, awsProfile, fn)
}

// Deprecated: use sub package.
func AwsVaultCredentials(ctx context.Context, awsProfile string) (aws.Credentials, error) {
	deprecated()
	return awsvault.AwsVaultCredentials(ctx, awsProfile)
}

// Deprecated: use sub package.
func AwsVaultEnv(ctx context.Context, awsProfile string) (map[string]string, error) {
	deprecated()
	return awsvault.AwsVaultEnv(ctx, awsProfile)
}
