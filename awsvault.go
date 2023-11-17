package mageutil

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/elisasre/mageutil/awsvault"
)

// AwsVault runs embedded aws-vault command with given args.
// Deprecated: use sub package.
func AwsVault(ctx context.Context, args ...string) error {
	deprecated()
	return awsvault.AwsVault(ctx, args...)
}

// AwsVaultExec executes command inside subshell with aws credentials.
// Deprecated: use sub package.
func AwsVaultExec(ctx context.Context, awsProfile string, args ...string) error {
	deprecated()
	return awsvault.AwsVaultExec(ctx, awsProfile, args...)
}

// AwsWithEnvCredentials sets aws credentials in environment variables, executes given functions and unsets env.
// Deprecated: use sub package.
func AwsWithEnvCredentials(ctx context.Context, awsProfile string, fn func() error) error {
	deprecated()
	return awsvault.AwsWithEnvCredentials(ctx, awsProfile, fn)
}

// AwsVaultCredentials fetches aws vault credentials.
// Deprecated: use sub package.
func AwsVaultCredentials(ctx context.Context, awsProfile string) (aws.Credentials, error) {
	deprecated()
	return awsvault.AwsVaultCredentials(ctx, awsProfile)
}

// AwsVaultEnv fetches aws vault credentials in env format.
// Deprecated: use sub package.
func AwsVaultEnv(ctx context.Context, awsProfile string) (map[string]string, error) {
	deprecated()
	return awsvault.AwsVaultEnv(ctx, awsProfile)
}
