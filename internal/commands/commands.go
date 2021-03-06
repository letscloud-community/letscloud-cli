package commands

import (
	"github.com/letscloud-community/letscloud-go"
	"github.com/urfave/cli/v2"
)

//Commands represents all the available commands in the CLI
type Commands struct {
	cmds []*cli.Command
	sdk  *letscloud.LetsCloud
}

//Sdk returns the encapsulated sdk
func (c Commands) Sdk() *letscloud.LetsCloud {
	return c.sdk
}

//New instantiates new commands object
func New(sdk *letscloud.LetsCloud) *Commands {
	return &Commands{sdk: sdk}
}

//Commands returns all the available commands in the CLI
func (c Commands) Commands() []*cli.Command {
	return []*cli.Command{
		c.apiKeyCmd(),
		c.locationsCmd(),
		c.plansCmd(),
		c.imagesCmd(),
		c.sshKeyCmd(),
		c.instanceCmd(),
		c.profileCmd(),
	}
}
