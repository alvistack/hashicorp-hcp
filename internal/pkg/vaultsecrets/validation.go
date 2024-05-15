// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vaultsecrets

import (
	"errors"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/hashicorp/hcp/internal/pkg/cmd"
)

// RequireVaultSecretsAppName requires that the profile has a set organization and project ID along with
// the Vault Secrets application name.
func RequireVaultSecretsAppName(ctx *cmd.Context, appName string) error {
	err := cmd.RequireOrgAndProject(ctx)
	if err != nil {
		return err
	}

	if appName != "" || ctx.Profile.VaultSecrets != nil && ctx.Profile.VaultSecrets.AppName != "" {
		return nil
	}

	cs := ctx.IO.ColorScheme()
	help := heredoc.Docf(`%v

	Set the app name using the --app flag or set the app name on your active profile one of the following commands:

	%v
	%v
	`,
		cs.String("Vault Secrets application name must set.").Color(cs.Orange()),
		cs.String("$ hcp profile set vault-secrets/app <app_name>").Bold(),
		cs.String("$ hcp profile init --vault-secrets").Bold(),
	)

	return errors.New(help)
}
