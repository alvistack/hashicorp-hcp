package application

import (
	"fmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-waypoint-service/preview/2023-08-18/client/waypoint_service"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-waypoint-service/preview/2023-08-18/models"
	"github.com/hashicorp/hcp/internal/pkg/cmd"
	"github.com/hashicorp/hcp/internal/pkg/flagvalue"
	"github.com/hashicorp/hcp/internal/pkg/heredoc"
	"github.com/pkg/errors"
)

func NewCmdCreateApplication(ctx *cmd.Context, opts *ApplicationOpts) *cmd.Command {
	cmd := &cmd.Command{
		Name:      "create",
		ShortHelp: "Create a new HCP Waypoint application.",
		LongHelp: heredoc.New(ctx.IO).Must(`
The {{ Bold "hcp waypoint applications create" }} command lets you create
a new HCP Waypoint application.
		`),
		Examples: []cmd.Example{
			{
				Preamble: "Create a new HCP Waypoint application:",
				Command: heredoc.New(ctx.IO, heredoc.WithPreserveNewlines()).Must(`
$ hcp waypoint applications create -n my-application -t my-template
`),
			},
		},
		RunF: func(c *cmd.Command, args []string) error {
			if opts.testFunc != nil {
				return opts.testFunc(c, args)
			}
			return applicationCreate(opts)
		},
		PersistentPreRun: func(c *cmd.Command, args []string) error {
			return cmd.RequireOrgAndProject(ctx)
		},
		Flags: cmd.Flags{
			Local: []*cmd.Flag{
				{
					Name:         "name",
					Shorthand:    "n",
					DisplayValue: "NAME",
					Description:  "The name of the application.",
					Value:        flagvalue.Simple("", &opts.Name),
					Required:     true,
				},
				{
					Name:         "template-name",
					Shorthand:    "t",
					DisplayValue: "TEMPLATE_NAME",
					Description:  "The name of the template to use for the application.",
					Value:        flagvalue.Simple("", &opts.TemplateName),
					Required:     true,
				},
				{
					Name:         "action-config-name",
					DisplayValue: "ACTION_CONFIG_NAME",
					Description:  "The name of the action configuration to be added to the application.",
					Value:        flagvalue.SimpleSlice(nil, &opts.ActionConfigNames),
					Required:     false,
					Repeatable:   true,
				},
			},
		},
	}

	return cmd
}

func applicationCreate(opts *ApplicationOpts) error {
	ns, err := opts.Namespace()
	if err != nil {
		return errors.Wrap(err, "unable to access HCP project")
	}

	actionConfigs := make([]*models.HashicorpCloudWaypointActionCfgRef, len(opts.ActionConfigNames))
	for i, name := range opts.ActionConfigNames {
		actionConfigs[i] = &models.HashicorpCloudWaypointActionCfgRef{
			Name: name,
		}
	}

	_, err = opts.WS.WaypointServiceCreateApplicationFromTemplate(
		&waypoint_service.WaypointServiceCreateApplicationFromTemplateParams{
			NamespaceID: ns.ID,
			Context:     opts.Ctx,
			Body: &models.HashicorpCloudWaypointWaypointServiceCreateApplicationFromTemplateBody{
				Name: opts.Name,
				ApplicationTemplate: &models.HashicorpCloudWaypointRefApplicationTemplate{
					Name: opts.TemplateName,
				},
				ActionCfgRefs: actionConfigs,
			},
		}, nil)
	if err != nil {
		return errors.Wrapf(err, "failed to create application %q", opts.Name)
	}

	fmt.Fprintf(opts.IO.Err(), "Application %q created.", opts.Name)

	return nil
}
