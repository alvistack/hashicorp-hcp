// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package secrets

import (
	"context"
	"fmt"
	"os"

	preview_secret_service "github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/preview/2023-11-28/client/secret_service"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/stable/2023-06-13/client/secret_service"
	appname "github.com/hashicorp/hcp/internal/commands/vaultsecrets/secrets/helper"
	"github.com/hashicorp/hcp/internal/pkg/cmd"
	"github.com/hashicorp/hcp/internal/pkg/flagvalue"
	"github.com/hashicorp/hcp/internal/pkg/format"
	"github.com/hashicorp/hcp/internal/pkg/heredoc"
	"github.com/hashicorp/hcp/internal/pkg/iostreams"
	"github.com/hashicorp/hcp/internal/pkg/profile"
)

func NewCmdOpen(ctx *cmd.Context, runF func(*OpenOpts) error) *cmd.Command {
	opts := &OpenOpts{
		Ctx:           ctx.ShutdownCtx,
		Profile:       ctx.Profile,
		IO:            ctx.IO,
		Output:        ctx.Output,
		PreviewClient: preview_secret_service.New(ctx.HCP, nil),
		Client:        secret_service.New(ctx.HCP, nil),
	}

	cmd := &cmd.Command{
		Name:      "open",
		ShortHelp: "Open a static secret.",
		LongHelp: heredoc.New(ctx.IO).Must(`
		The {{ template "mdCodeOrBold" "hcp vault-secrets secrets open" }} command reads the plaintext value of a static secret from the Vault Secrets application.
		`),
		Examples: []cmd.Example{
			{
				Preamble: `Open plaintext secret:`,
				Command: heredoc.New(ctx.IO, heredoc.WithPreserveNewlines()).Must(`
				$ hcp vault-secrets secret open "test_secret"
				`),
			},
		},
		Args: cmd.PositionalArguments{
			Args: []cmd.PositionalArgument{
				{
					Name:          "NAME",
					Documentation: "The name of the secret to open.",
				},
			},
		},
		Flags: cmd.Flags{
			Local: []*cmd.Flag{
				{
					Name:         "out-file",
					DisplayValue: "OUTPUT_FILE_PATH",
					Shorthand:    "o",
					Description:  "File path where the secret value should be written.",
					Value:        flagvalue.Simple("", &opts.OutputFilePath),
				},
			},
		},
		RunF: func(c *cmd.Command, args []string) error {
			opts.AppName = appname.Get()
			opts.SecretName = args[0]

			if runF != nil {
				return runF(opts)
			}
			return openRun(opts)
		},
	}

	return cmd
}

type OpenOpts struct {
	Ctx     context.Context
	Profile *profile.Profile
	IO      iostreams.IOStreams
	Output  *format.Outputter

	AppName        string
	SecretName     string
	OutputFilePath string
	PreviewClient  preview_secret_service.ClientService
	Client         secret_service.ClientService
}

func openRun(opts *OpenOpts) error {
	var fd *os.File
	var err error
	if opts.OutputFilePath != "" {
		fd, err = os.OpenFile(opts.OutputFilePath, os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			return fmt.Errorf("failed to open the outout file %q: %w", opts.OutputFilePath, err)
		}
	}
	defer func() {
		if opts.OutputFilePath != "" {
			fd.Close()
		}
	}()

	req := secret_service.NewOpenAppSecretParamsWithContext(opts.Ctx)
	req.LocationOrganizationID = opts.Profile.OrganizationID
	req.LocationProjectID = opts.Profile.ProjectID
	req.AppName = opts.AppName
	req.SecretName = opts.SecretName

	resp, err := opts.Client.OpenAppSecret(req, nil)
	if err != nil {
		return fmt.Errorf("failed to read the secret %q: %w", opts.SecretName, err)
	}

	if opts.OutputFilePath != "" {
		_, err = fd.WriteString(resp.Payload.Secret.Version.Value)
		if err != nil {
			return fmt.Errorf("failed to write the secret value to the output file %q: %w", opts.OutputFilePath, err)
		}
		fmt.Fprintf(opts.IO.Err(), "%s Successfully wrote plaintext secret with name %q to path %q\n", opts.IO.ColorScheme().SuccessIcon(), opts.SecretName, opts.OutputFilePath)
		return nil
	}
	d := newDisplayer(true).OpenAppSecrets(resp.Payload.Secret).SetDefaultFormat(format.Pretty)
	return opts.Output.Display(d.OpenAppSecrets(resp.Payload.Secret))
}
