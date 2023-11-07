// Package awsvault exposes github.com/99designs/aws-vault as library.
package awsvault

import (
	"context"
	"fmt"
	"os"

	"github.com/99designs/aws-vault/v7/cli"
	"github.com/99designs/aws-vault/v7/iso8601"
	"github.com/99designs/aws-vault/v7/vault"
	"github.com/99designs/keyring"
	"github.com/alecthomas/kingpin/v2"
	"github.com/aws/aws-sdk-go-v2/aws"
)

// AwsVault runs embedded aws-vault command with given args.
func AwsVault(ctx context.Context, args ...string) error {
	app := kingpin.New("aws-vault", "A vault for securely storing and accessing AWS credentials in development environments.")
	app.Version("Embedded")

	a := cli.ConfigureGlobals(app)
	cli.ConfigureAddCommand(app, a)
	cli.ConfigureRemoveCommand(app, a)
	cli.ConfigureListCommand(app, a)
	cli.ConfigureRotateCommand(app, a)
	cli.ConfigureExecCommand(app, a)
	cli.ConfigureExportCommand(app, a)
	cli.ConfigureClearCommand(app, a)
	cli.ConfigureLoginCommand(app, a)
	cli.ConfigureProxyCommand(app)

	_, err := app.Parse(args)
	return err
}

// AwsVaultExec executes command inside subshell with aws credentials.
func AwsVaultExec(ctx context.Context, awsProfile string, args ...string) error {
	return AwsVault(ctx, append([]string{"exec", awsProfile, "--"}, args...)...)
}

// AwsWithEnvCredentials sets aws credentials in environment variables, executes given functions and unsets env.
func AwsWithEnvCredentials(ctx context.Context, awsProfile string, fn func() error) error {
	env, err := AwsVaultEnv(ctx, awsProfile)
	if err != nil {
		return err
	}

	defer func() {
		for k := range env {
			_ = os.Unsetenv(k)
		}
	}()

	for k, v := range env {
		_ = os.Setenv(k, v)
	}

	return fn()
}

// AwsVaultCredentials fetches aws vault credentials.
func AwsVaultCredentials(ctx context.Context, awsProfile string) (aws.Credentials, error) {
	app := kingpin.New("aws-vault", "A vault for securely storing and accessing AWS credentials in development environments.")
	app.Version("Embedded")
	a := cli.ConfigureGlobals(app)
	input := cli.ExportCommandInput{
		ProfileName: awsProfile,
		Config: vault.ProfileConfig{
			MfaPromptMethod: a.PromptDriver(false),
		},
	}

	f, err := a.AwsConfigFile()
	if err != nil {
		return aws.Credentials{}, err
	}
	keyring, err := a.Keyring()
	if err != nil {
		return aws.Credentials{}, err
	}

	return fetchCredentials(input, f, keyring)
}

// AwsVaultEnv fetches aws vault credentials in env format.
func AwsVaultEnv(ctx context.Context, awsProfile string) (map[string]string, error) {
	creds, err := AwsVaultCredentials(ctx, awsProfile)
	if err != nil {
		return nil, err
	}

	return CredentialToEnvMap(creds), nil
}

func CredentialToEnvMap(creds aws.Credentials) map[string]string {
	env := make(map[string]string)
	env["AWS_ACCESS_KEY_ID"] = creds.AccessKeyID
	env["AWS_SECRET_ACCESS_KEY"] = creds.SecretAccessKey

	if creds.SessionToken != "" {
		env["AWS_SESSION_TOKEN"] = creds.SessionToken
	}
	if creds.CanExpire {
		env["AWS_CREDENTIAL_EXPIRATION"] = iso8601.Format(creds.Expires)
	}
	return env
}

func fetchCredentials(input cli.ExportCommandInput, f *vault.ConfigFile, keyring keyring.Keyring) (aws.Credentials, error) {
	if os.Getenv("AWS_VAULT") != "" {
		return aws.Credentials{}, fmt.Errorf("in an existing aws-vault subshell; 'exit' from the subshell or unset AWS_VAULT to force")
	}

	config, err := vault.NewConfigLoader(input.Config, f, input.ProfileName).GetProfileConfig(input.ProfileName)
	if err != nil {
		return aws.Credentials{}, fmt.Errorf("error loading config: %w", err)
	}

	ckr := &vault.CredentialKeyring{Keyring: keyring}
	credsProvider, err := vault.NewTempCredentialsProvider(config, ckr, input.NoSession, false)
	if err != nil {
		return aws.Credentials{}, fmt.Errorf("error getting temporary credentials: %w", err)
	}

	creds, err := credsProvider.Retrieve(context.TODO())
	if err != nil {
		return aws.Credentials{}, fmt.Errorf("failed to get credentials for %s: %w", input.ProfileName, err)
	}

	return creds, nil
}
