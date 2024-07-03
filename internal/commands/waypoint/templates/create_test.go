// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package templates

import (
	"context"
	"testing"

	"github.com/go-openapi/runtime/client"
	"github.com/hashicorp/hcp/internal/pkg/cmd"
	"github.com/hashicorp/hcp/internal/pkg/format"
	"github.com/hashicorp/hcp/internal/pkg/iostreams"
	"github.com/hashicorp/hcp/internal/pkg/profile"
	"github.com/stretchr/testify/require"
)

func TestCmdTemplateCreate(t *testing.T) {
	t.Parallel()

	cases := []struct {
		Name    string
		Args    []string
		Profile func(t *testing.T) *profile.Profile
		Error   string
		Expect  *TemplateOpts
	}{
		{
			Name:    "No Org",
			Profile: profile.TestProfile,
			Args:    []string{},
			Error:   "Organization ID must be configured",
		},
		{
			Name: "no args",
			Profile: func(t *testing.T) *profile.Profile {
				return profile.TestProfile(t).SetOrgID("123")
			},
			Args:  []string{},
			Error: "accepts 1 arg(s), received 0",
		},
		{
			Name: "happy",
			Profile: func(t *testing.T) *profile.Profile {
				return profile.TestProfile(t).SetOrgID("123")
			},
			Args: []string{
				"-n=cli-test",
				"-s", "A template created using the CLI.",
				"--tfc-project-id", "prj-abcdefghij",
				"--tfc-project-name", "test",
				"--tfc-no-code-module-source", "private/waypoint/waypoint-nocode-module/null",
				"-l", "cli",
				"-d", "A template created with the CLI.",
				"-t", "cli=true",
				"--readme-markdown-template-file", "readme_test.txt",
			},
			Expect: &TemplateOpts{
				Name:                        "cli-test",
				Summary:                     "A template created using the CLI.",
				Description:                 "A template created with the CLI.",
				TerraformCloudProjectID:     "prj-abcdefghij",
				TerraformCloudProjectName:   "test",
				TerraformNoCodeModuleSource: "private/waypoint/waypoint-nocode-module/null",
				ReadmeMarkdownTemplateFile:  "readme_test.txt",
				Labels:                      []string{"cli"},
				Tags:                        map[string]string{"cli": "true"},
			},
		},
		{
			Name: "accepts valid variable file",
			Profile: func(t *testing.T) *profile.Profile {
				return profile.TestProfile(t).SetOrgID("123")
			},
			Args: []string{
				"-n=cli-test",
				"-s", "A template created using the CLI.",
				"--tfc-project-id", "prj-abcdefghij",
				"--tfc-project-name", "test",
				"--tfc-no-code-module-source", "private/waypoint/waypoint-nocode-module/null",
				"--tfc-no-code-module-version", "0.0.1",
				"-l", "cli",
				"-d", "A template created with the CLI.",
				"-t", "cli=true",
				"--readme-markdown-template-file", "readme_test.txt",
				"--variable-options-file", "variable_options.hcl",
			},
			Expect: &TemplateOpts{
				Name:                         "cli-test",
				Summary:                      "A template created using the CLI.",
				Description:                  "A template created with the CLI.",
				TerraformCloudProjectID:      "prj-abcdefghij",
				TerraformCloudProjectName:    "test",
				TerraformNoCodeModuleSource:  "private/waypoint/waypoint-nocode-module/null",
				TerraformNoCodeModuleVersion: "0.0.1",
				ReadmeMarkdownTemplateFile:   "readme_test.txt",
				VariableOptionsFile:          "variable_options.hcl",
				// VariableOptions:   "variable_options.hcl",
				Labels: []string{"cli"},
				Tags:   map[string]string{"cli": "true"},
			},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()

			r := require.New(t)

			// Create a context.
			io := iostreams.Test()
			ctx := &cmd.Context{
				IO:          io,
				Profile:     c.Profile(t),
				Output:      format.New(io),
				HCP:         &client.Runtime{},
				ShutdownCtx: context.Background(),
			}

			var tplOpts TemplateOpts
			tplOpts.testFunc = func(c *cmd.Command, args []string) error {
				return nil
			}
			cmd := NewCmdCreate(ctx, &tplOpts)
			cmd.SetIO(io)

			cmd.Run(c.Args)

			if c.Expect != nil {
				r.NotNil(c.Expect)

				r.Equal(c.Expect.Name, tplOpts.Name)
				r.Equal(c.Expect.Description, tplOpts.Description)
				r.Equal(c.Expect.Summary, tplOpts.Summary)
				r.Equal(c.Expect.TerraformCloudProjectID, tplOpts.TerraformCloudProjectID)
				r.Equal(c.Expect.TerraformCloudProjectName, tplOpts.TerraformCloudProjectName)
				r.Equal(c.Expect.TerraformNoCodeModuleSource, tplOpts.TerraformNoCodeModuleSource)
				r.Equal(c.Expect.ReadmeMarkdownTemplateFile, tplOpts.ReadmeMarkdownTemplateFile)
				r.Equal(c.Expect.VariableOptionsFile, tplOpts.VariableOptionsFile)
				r.Equal(c.Expect.Labels, tplOpts.Labels)
				r.Equal(c.Expect.Tags, tplOpts.Tags)
			}
		})
	}
}

// Test_VariableInputFileParse tests the parsing of variable options from a
// file.
func Test_VariableInputFileParse(t *testing.T) {
	t.Parallel()

	t.Run("can parse variables", func(t *testing.T) {
		t.Parallel()

		r := require.New(t)

		hcl := `
          variable_option "string_variable" {
            type = "string"
            options = [
              "a string value",
            ]
            user_editable = false
          }
`

		variableInputs, err := parseVariableInputs("blah.hcl", []byte(hcl))
		r.NoError(err)
		r.Equal(1, len(variableInputs))
		r.Equal("string_variable", variableInputs[0].Name)
		r.Equal("string", variableInputs[0].VariableType)
		r.Equal(false, variableInputs[0].UserEditable)
	})

	t.Run("handles multiple options", func(t *testing.T) {
		t.Parallel()

		r := require.New(t)

		hcl := `
          variable_option "string_variable" {
            type = "string"
            options = [
              "a string value",
        		"another string value",
            ]
            user_editable =true
          }
`

		variableInputs, err := parseVariableInputs("blah.hcl", []byte(hcl))
		r.NoError(err)
		r.Equal(1, len(variableInputs))
		r.Equal("string_variable", variableInputs[0].Name)
		r.Equal(2, len(variableInputs[0].Options))
		r.Equal(true, variableInputs[0].UserEditable)
	})

	t.Run("handles multiple variables", func(t *testing.T) {
		t.Parallel()

		r := require.New(t)

		hcl := `
variable_option "string_variable" {
  type = "string"
  options = [
    "a string value",
		"another",
  ]
  user_editable = false
}

variable_option "misc_variable" {
  type = "int"
  options = [
    8,
		2,
  ]
  user_editable = false
}
`

		variableInputs, err := parseVariableInputs("blah.hcl", []byte(hcl))
		r.NoError(err)
		r.Equal(2, len(variableInputs))
		r.Equal("string_variable", variableInputs[0].Name)
		r.Equal(2, len(variableInputs[0].Options))
		r.Equal(false, variableInputs[0].UserEditable)

		r.Equal("misc_variable", variableInputs[1].Name)
		r.Equal(2, len(variableInputs[1].Options))
		r.Equal("int", variableInputs[1].VariableType)
		r.Equal(false, variableInputs[1].UserEditable)
	})

	t.Run("errors on missing attributes", func(t *testing.T) {
		t.Parallel()

		r := require.New(t)

		hcl := `
variable_option "" {
  options = [
    8,
		2,
  ]
}
`

		_, err := parseVariableInputs("blah.hcl", []byte(hcl))
		r.Error(err)
	})

	t.Run("returns nil empty", func(t *testing.T) {
		t.Parallel()

		r := require.New(t)

		hcl := ``

		variableInputs, err := parseVariableInputs("blah.hcl", []byte(hcl))
		r.NoError(err)
		r.Equal(0, len(variableInputs))
	})
}
