package version

import (
	"fmt"

	"github.com/hashicorp/hcp/internal/pkg/cmd"
	"github.com/hashicorp/hcp/internal/pkg/heredoc"
	"github.com/hashicorp/hcp/version"
)

func NewCmdVersion(ctx *cmd.Context) *cmd.Command {
	cmd := &cmd.Command{
		Name:      "version",
		ShortHelp: "Display the HCP CLI version.",
		LongHelp: heredoc.New(ctx.IO).Must(`
		Display the HCP CLI version.
		`),
		RunF: func(c *cmd.Command, args []string) error {
			fmt.Fprintln(ctx.IO.Out(), version.GetHumanVersion())
			return nil
		},
		NoAuthRequired: true,
	}

	return cmd
}
