// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package secrets

import (
	preview_secret_models "github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/preview/2023-11-28/models"
	models "github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/stable/2023-06-13/models"
	"github.com/hashicorp/hcp/internal/pkg/format"
)

type displayer struct {
	secrets        []*models.Secrets20230613Secret
	previewSecrets []*preview_secret_models.Secrets20231128Secret
	openAppSecrets []*models.Secrets20230613OpenSecret
	single         bool
	fields         []format.Field
}

func newDisplayer(single bool) *displayer {
	return &displayer{
		single: single,
	}
}

func (d *displayer) Secrets(secrets ...*models.Secrets20230613Secret) *displayer {
	d.secrets = secrets
	return d
}

func (d *displayer) PreviewSecrets(secrets ...*preview_secret_models.Secrets20231128Secret) *displayer {
	d.previewSecrets = secrets
	return d
}

func (d *displayer) OpenAppSecrets(secrets ...*models.Secrets20230613OpenSecret) *displayer {
	d.openAppSecrets = secrets
	return d
}

func (d *displayer) AddFields(fields []format.Field) []format.Field {
	d.fields = append(d.fields, fields...)
	return d.fields
}

func (d *displayer) DefaultFormat() format.Format {
	return format.Table
}

func (d *displayer) Payload() any {
	if d.previewSecrets != nil {
		return d.previewSecretsPayload()
	}

	if d.openAppSecrets != nil {
		return d.openAppSecretsPayload()
	}

	if d.secrets == nil {
		return nil
	}
	return d.secretsPayload()
}

func (d *displayer) FieldTemplates() []format.Field {
	fields := []format.Field{
		{
			Name:        "Secret Name",
			ValueFormat: "{{ .Name }}",
		},
		{
			Name:        "Latest Version",
			ValueFormat: "{{ .LatestVersion }}",
		},
		{
			Name:        "Created At",
			ValueFormat: "{{ .CreatedAt }}",
		},
	}
	fields = append(fields, d.fields...)
	return fields
}

func (d *displayer) secretsPayload() any {
	if d.single {
		if len(d.secrets) != 1 {
			return nil
		}
		return d.secrets[0]
	}
	return d.secrets
}

func (d *displayer) previewSecretsPayload() any {
	if d.single {
		if len(d.previewSecrets) != 1 {
			return nil
		}
		return d.previewSecrets[0]
	}
	return d.previewSecrets
}

func (d *displayer) openAppSecretsPayload() any {
	if d.single {
		if len(d.openAppSecrets) != 1 {
			return nil
		}
		return d.openAppSecrets[0]
	}
	return d.openAppSecrets
}
