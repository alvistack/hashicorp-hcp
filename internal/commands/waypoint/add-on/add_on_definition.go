package addon

import (
	"github.com/hashicorp/hcp/internal/commands/waypoint/opts"
	"github.com/hashicorp/hcp/internal/pkg/cmd"
)

type AddOnDefinitionOpts struct {
	opts.WaypointOpts

	Name                       string
	Summary                    string
	Description                string
	ReadmeMarkdownTemplateFile string
	Labels                     []string

	TerraformNoCodeModuleSource  string
	TerraformNoCodeModuleVersion string
	TerraformCloudProjectName    string
	TerraformCloudProjectID      string

	// testFunc is used for testing, so that the command can be tested without
	// using the real API.
	testFunc func(c *cmd.Command, args []string) error
}

func NewCmdAddOnDefinition(ctx *cmd.Context) *cmd.Command {
	opts := &AddOnDefinitionOpts{
		WaypointOpts: opts.New(ctx),
	}

	cmd := &cmd.Command{
		Name:      "definitions",
		ShortHelp: "Manage HCP Waypoint add-on definitions.",
		LongHelp: "Manage HCP Waypoint add-on definitions. Add-on definitions " +
			"are reusable configurations for creating add-ons.",
	}

	cmd.AddChild(NewCmdAddOnDefinitionCreate(ctx, opts))

	return cmd
}
