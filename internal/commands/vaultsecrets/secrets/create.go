// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package secrets

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	preview_secret_service "github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/preview/2023-11-28/client/secret_service"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/stable/2023-06-13/client/secret_service"

	"github.com/hashicorp/hcp/internal/pkg/cmd"
	"github.com/hashicorp/hcp/internal/pkg/flagvalue"
	"github.com/hashicorp/hcp/internal/pkg/format"
	"github.com/hashicorp/hcp/internal/pkg/heredoc"
	"github.com/hashicorp/hcp/internal/pkg/iostreams"
	"github.com/hashicorp/hcp/internal/pkg/profile"
)

func NewCmdCreate(ctx *cmd.Context, runF func(*CreateOpts) error) *cmd.Command {
	opts := &CreateOpts{
		Ctx:           ctx.ShutdownCtx,
		Profile:       ctx.Profile,
		IO:            ctx.IO,
		Output:        ctx.Output,
		PreviewClient: preview_secret_service.New(ctx.HCP, nil),
		Client:        secret_service.New(ctx.HCP, nil),
	}

	cmd := &cmd.Command{
		Name:      "create",
		ShortHelp: "Create a new static secret.",
		LongHelp: heredoc.New(ctx.IO).Must(`
		The {{ template "mdCodeOrBold" "hcp vault-secrets secrets create" }} command creates a new static secret under an Vault Secrets application.

		Once the secret is created, it can be read using
		{{ template "mdCodeOrBold" "hcp vault-secrets secrets read" }} subcommand.
		`),
		Examples: []cmd.Example{
			{
				Preamble: `Create new secret in Vault Secrets application on active profile:`,
				Command: heredoc.New(ctx.IO, heredoc.WithPreserveNewlines()).Must(`
				$ hcp vault-secrets secrets create secret_1 --data-file=tmp/secrets1.txt
				`),
			},
			{
				Preamble: `Create new secret in Vault Secrets application on active profile with pipe:`,
				Command: heredoc.New(ctx.IO, heredoc.WithNoWrap()).Must(`
				$ echo "my super secret" | hcp vault-secrets secrets create secret_2 --data-file=-
				`),
			},
			{
				Preamble: `Create secret in different Vault Secrets application, not active profile:`,
				Command: heredoc.New(ctx.IO, heredoc.WithNoWrap()).Must(`
				$ hcp vault-secrets secrets create secret_3 --app-name test-app --secret_file=/tmp/secrets2.txt
				`),
			},
		},
		Args: cmd.PositionalArguments{
			Args: []cmd.PositionalArgument{
				{
					Name:          "SECRET_NAME",
					Documentation: "The name of the secret to create.",
				},
			},
		},
		Flags: cmd.Flags{
			Local: []*cmd.Flag{
				{
					Name:         "data-file",
					DisplayValue: "DATA_FILE_PATH",
					Description:  "File path to read secret data from. Set this to '-' to read the secret data from stdin.",
					Value:        flagvalue.Simple("", &opts.SecretFilePath),
				},
			},
		},
		RunF: func(c *cmd.Command, args []string) error {
			opts.AppName = appName
			opts.SecretName = args[0]

			if runF != nil {
				return runF(opts)
			}
			return createRun(opts)
		},
	}

	return cmd
}

type CreateOpts struct {
	Ctx     context.Context
	Profile *profile.Profile
	IO      iostreams.IOStreams
	Output  *format.Outputter

	AppName              string
	SecretName           string
	SecretValuePlaintext string
	SecretFilePath       string
	PreviewClient        preview_secret_service.ClientService
	Client               secret_service.ClientService
}

func createRun(opts *CreateOpts) error {
	if err := readPlainTextSecret(opts); err != nil {
		return err
	}

	req := secret_service.NewCreateAppKVSecretParamsWithContext(opts.Ctx)
	req.LocationOrganizationID = opts.Profile.OrganizationID
	req.LocationProjectID = opts.Profile.ProjectID
	req.AppName = opts.AppName

	req.Body = secret_service.CreateAppKVSecretBody{
		Name:  opts.SecretName,
		Value: opts.SecretValuePlaintext,
	}

	resp, err := opts.Client.CreateAppKVSecret(req, nil)
	if err != nil {
		return fmt.Errorf("failed to create secret with name: %q - %w", opts.SecretName, err)
	}

	if opts.Output.GetFormat() == format.Unset {
		fmt.Fprintf(opts.IO.Err(), "%s Successfully created secret with name: %q\n", opts.IO.ColorScheme().SuccessIcon(), opts.SecretName)
		return nil
	}

	return opts.Output.Display(newDisplayer(true, resp.Payload.Secret))
}

func readPlainTextSecret(opts *CreateOpts) error {
	// If the secret value is provided, then we don't need to read it from the file
	// this is used for making testing easier without needing to create a file
	if opts.SecretValuePlaintext != "" {
		return nil
	}

	if opts.SecretFilePath == "" {
		return errors.New("data file path is required")
	}

	if opts.SecretFilePath == "-" {
		plaintextSecretBytes, err := io.ReadAll(opts.IO.In())
		if err != nil {
			return fmt.Errorf("failed to read the plaintext secret: %w", err)
		}

		if len(plaintextSecretBytes) == 0 {
			return errors.New("secret value cannot be empty")
		}
		opts.SecretValuePlaintext = string(plaintextSecretBytes)
		return nil
	}

	fileInfo, err := os.Stat(opts.SecretFilePath)
	if err != nil {
		return fmt.Errorf("failed to get data file info: %w", err)
	}

	if fileInfo.Size() == 0 {
		return errors.New("data file cannot be empty")
	}

	data, err := os.ReadFile(opts.SecretFilePath)
	if err != nil {
		return fmt.Errorf("unable to read the data file: %w", err)
	}
	opts.SecretValuePlaintext = string(data)
	return nil
}
