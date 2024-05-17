// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package secrets

import (
	"github.com/hashicorp/hcp/internal/commands/vaultsecrets/secrets/helper"
	"github.com/hashicorp/hcp/internal/commands/vaultsecrets/secrets/versions"
	"github.com/hashicorp/hcp/internal/pkg/cmd"
	"github.com/hashicorp/hcp/internal/pkg/heredoc"
)

func NewCmdSecrets(ctx *cmd.Context) *cmd.Command {
	cmd := &cmd.Command{
		Name:      "secrets",
		ShortHelp: "Manage Vault Secrets application secrets.",
		LongHelp: heredoc.New(ctx.IO).Must(`
		The {{ template "mdCodeOrBold" "hcp vault-secrets secrets" }} command group lets you
		manage Vault Secrets application secrets.
		`),
		Flags: cmd.Flags{
			Persistent: []*cmd.Flag{
				{
					Name:         "app",
					DisplayValue: "NAME",
					Description:  "The name of the Vault Secrets application. If not specified, the value from the active profile will be used.",
					Shorthand:    "a",
					Value:        helper.AppNameFlag(),
				},
			},
		},
		PersistentPreRun: func(c *cmd.Command, args []string) error {

			return helper.RequireVaultSecretsAppName(ctx)
		},
	}

	cmd.AddChild(NewCmdCreate(ctx, nil))
	cmd.AddChild(NewCmdRead(ctx, nil))
	cmd.AddChild(NewCmdDelete(ctx, nil))
	cmd.AddChild(NewCmdList(ctx, nil))

	cmd.AddChild(versions.NewCmdVersions(ctx))
	return cmd
}
