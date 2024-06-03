// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package run

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/hashicorp/go-hclog"
	preview_secret_service "github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/preview/2023-11-28/client/secret_service"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/stable/2023-06-13/client/secret_service"
	"github.com/hashicorp/hcp/internal/commands/vaultsecrets/apps/helper"
	"github.com/hashicorp/hcp/internal/commands/vaultsecrets/secrets/appname"
	"github.com/hashicorp/hcp/internal/pkg/cmd"
	"github.com/hashicorp/hcp/internal/pkg/format"
	"github.com/hashicorp/hcp/internal/pkg/heredoc"
	"github.com/hashicorp/hcp/internal/pkg/iostreams"
	"github.com/hashicorp/hcp/internal/pkg/profile"
)

type RunOpts struct {
	Ctx     context.Context
	Profile *profile.Profile
	IO      iostreams.IOStreams
	Output  *format.Outputter
	Logger  hclog.Logger

	AppName       string
	Command       []string
	PreviewClient preview_secret_service.ClientService
	Client        secret_service.ClientService
}

func NewCmdRun(ctx *cmd.Context, runF func(*RunOpts) error) *cmd.Command {
	opts := &RunOpts{
		Ctx:           ctx.ShutdownCtx,
		Profile:       ctx.Profile,
		IO:            ctx.IO,
		Output:        ctx.Output,
		PreviewClient: preview_secret_service.New(ctx.HCP, nil),
		Client:        secret_service.New(ctx.HCP, nil),
	}

	cmd := &cmd.Command{
		Name:      "run",
		ShortHelp: "Run a process with secrets from a Vault Secrets app.",
		LongHelp: heredoc.New(ctx.IO).Must(`
		The {{ template "mdCodeOrBold" "hcp vault-secrets run" }} command lets you run 
		the provided command as a child process while injecting all of the app's secrets
		as environment variables, with all secret names converted to upper-case. The stdout
		and stderr from the child process are forwarded to the top level {{ template "mdCodeOrBold" "hcp vault-secrets run" }} command.
		`),
		Examples: []cmd.Example{
			{
				Preamble: `Inject secrets as environment variables:`,
				Command: heredoc.New(ctx.IO, heredoc.WithPreserveNewlines()).Must(`
				$ hcp vault-secrets run "env"
				`),
			},
		},
		Args: cmd.PositionalArguments{
			Args: []cmd.PositionalArgument{
				{
					Name:          "COMMAND",
					Documentation: "Defines the invocation of the child process to inject secrets to.",
				},
			},
			Validate: cmd.ArbitraryArgs,
		},
		Flags: cmd.Flags{
			Local: []*cmd.Flag{
				{
					Name:         "app",
					DisplayValue: "NAME",
					Description:  "The application you want to pull all secrets from.",
					Value:        appname.Flag(),
				},
			},
		},
		PersistentPreRun: func(c *cmd.Command, args []string) error {
			return appname.Require(ctx)
		},
		RunF: func(c *cmd.Command, args []string) error {
			opts.Command = args[0:]
			opts.AppName = appname.Get()

			if runF != nil {
				return runF(opts)
			}
			return runRun(opts)
		},
	}
	for _, f := range cmd.Flags.Local {
		if f.Name == "app" {
			f.Autocomplete = helper.PredictAppName(ctx, cmd, opts.PreviewClient)
		}
	}

	return cmd
}

func runRun(opts *RunOpts) (err error) {
	if len(opts.Command) == 0 {
		return fmt.Errorf("failed to run app secrets from name %q - no command provided", opts.AppName)
	}

	envSecrets, err := getAllSecretsForEnv(opts)
	if err != nil {
		return fmt.Errorf("failed to run with secrets in app %q: %w", opts.AppName, err)
	}

	childProcess := setupChildProcess(opts.Ctx, opts.Command, envSecrets)

	err = childProcess.Run()
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			exitCode := exitErr.ExitCode()
			return fmt.Errorf("failed to run with secrets in app %q; received code %v : %w", opts.AppName, exitCode, err)
		} else {
			return fmt.Errorf("failed to run with secrets in app %q: %w", opts.AppName, err)
		}
	}

	return nil
}

func getAllSecretsForEnv(opts *RunOpts) ([]string, error) {
	params := preview_secret_service.NewOpenAppSecretsParamsWithContext(opts.Ctx)
	params.OrganizationID = opts.Profile.OrganizationID
	params.ProjectID = opts.Profile.ProjectID
	params.AppName = opts.AppName

	res, err := opts.PreviewClient.OpenAppSecrets(params, nil)
	if err != nil {
		return nil, err
	}

	result := os.Environ()

	for _, secret := range res.Payload.Secrets {
		// we need to append results in case of duplicates we want secrets to override
		switch {
		case secret.RotatingVersion != nil:
			for name, value := range secret.RotatingVersion.Values {
				result = append(result, fmt.Sprintf("%v_%v=%v", strings.ToUpper(secret.Name), name, value))
			}
		case secret.StaticVersion != nil:
			result = append(result, fmt.Sprintf("%v=%v", strings.ToUpper(secret.Name), secret.StaticVersion.Value))
		}
	}

	return result, nil
}

func setupChildProcess(ctx context.Context, command []string, envVars []string) *exec.Cmd {
	var (
		args   []string
		cmd    string
		cmdCtx *exec.Cmd
	)

	if len(command) < 2 {
		pieces := strings.Split(command[0], " ")
		cmd = pieces[0]
		if len(pieces) > 1 {
			args = pieces[1:]
		}
	} else {
		cmd = command[0]
		args = command[1:]
	}

	cmdCtx = exec.CommandContext(ctx, cmd, args...)
	cmdCtx.Stdout = os.Stdout
	cmdCtx.Stdin = os.Stdin
	cmdCtx.Stderr = os.Stderr
	cmdCtx.Env = envVars

	return cmdCtx
}
