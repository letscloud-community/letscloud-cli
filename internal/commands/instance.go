package commands

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jedib0t/go-pretty/v6/list"
	"github.com/letscloud-community/letscloud-cli/internal/helpers"
	"github.com/letscloud-community/letscloud-cli/internal/pkg/writer"
	"github.com/letscloud-community/letscloud-go/domains"
	"github.com/urfave/cli/v2"
)

// Manage all your Instances
//
// Mostly getting list, details, start, shutdown, reboot, reset-password
func (c Commands) instanceCmd() *cli.Command {
	return &cli.Command{
		Name:  "instance",
		Usage: "Manage your instances",
		Subcommands: []*cli.Command{
			{
				Name: "list",
				Action: func(ctx *cli.Context) error {
					// List out All the Instances
					srvs, err := c.sdk.Instances()
					if err != nil {
						return err
					}

					wr := writer.New(os.Stdout)
					wr.WriteHeader("IDENTIFIER", "LABEL", "IPv4", "DC", "OS", "STATUS")
					for _, v := range srvs {
						wr.WriteData(v.Identifier, v.Label, v.IPAddresses[0].Address, v.Location.Slug, v.TemplateLabel, helpers.GetInstanceStatus(v.Booted, v.Locked, v.Built, v.Suspended))
					}

					return wr.Flush()
				},
			},
			{
				Name: "details",
				Action: func(ctx *cli.Context) error {
					// Data Validation
					identifier := ctx.Args().First()
					if identifier == "" {
						return errors.New("Identifier can not be empty. Please use `instance details <identifier>`")
					}
					srv, err := c.sdk.Instance(identifier)
					if err != nil {
						return err
					}

					//fmt.Println(srv.location)

					l := list.NewWriter()
					lTemp := list.List{}
					lTemp.Render()

					l.AppendItem(srv.Label)
					l.Indent()
					l.AppendItem("CPU " + strconv.Itoa(srv.CPUS))
					l.AppendItem("RAM " + strconv.Itoa(srv.Memory) + " MB")
					l.AppendItem("SSD " + strconv.Itoa(srv.TotalDiskSize) + " GB")
					l.UnIndent()

					l.AppendItem("Location")
					l.Indent()
					l.AppendItem(srv.Location.Slug)
					l.AppendItem(srv.Location.Country)
					l.AppendItem(srv.Location.City)
					l.UnIndent()

					l.AppendItem("Hostname")
					l.Indent()
					l.AppendItem(srv.Hostname)
					l.UnIndent()

					l.AppendItem("IPv4")
					l.Indent()
					for _, IPAddress := range srv.IPAddresses {
						if helpers.IsIpv4Regex(IPAddress.Address) {
							l.AppendItem(IPAddress.Address)
						}
					}
					l.UnIndent()

					l.AppendItem("IPv6")
					l.Indent()
					for _, IPAddress := range srv.IPAddresses {
						if !helpers.IsIpv4Regex(IPAddress.Address) {
							l.AppendItem(IPAddress.Address)
						}
					}
					l.UnIndent()

					l.AppendItem("Distro")
					l.Indent()
					l.AppendItem(srv.TemplateLabel)
					l.UnIndent()

					l.AppendItem("Root Password")
					l.Indent()
					l.AppendItem(srv.RootPassword)
					l.UnIndent()

					l.AppendItem("Status")
					l.Indent()
					l.AppendItem(helpers.GetInstanceStatus(srv.Booted, srv.Locked, srv.Built, srv.Suspended))
					l.UnIndent()

					l.SetStyle(list.StyleConnectedRounded)

					prefix := ""
					title := "Instance " + identifier
					fmt.Printf("%s\n", title)
					fmt.Println(strings.Repeat("-", len(title)+1))
					for _, line := range strings.Split(l.Render(), "\n") {
						fmt.Printf("%s%s\n", prefix, line)
					}
					fmt.Println()

					return nil
				},
			},
			{
				Name: "create",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name: "location",
					},
					&cli.StringFlag{
						Name: "plan",
					},
					&cli.StringFlag{
						Name: "hostname",
					},
					&cli.StringFlag{
						Name: "label",
					},
					&cli.StringFlag{
						Name: "image",
					},
					&cli.StringFlag{
						Name: "ssh",
					},
				},
				Action: func(ctx *cli.Context) error {
					// Data Validation
					locSlug := ctx.String("location")
					if locSlug == "" {
						return errors.New("Location slug can not be empty. Please use --location")
					}

					planSlug := ctx.String("plan")
					if planSlug == "" {
						return errors.New("Plan slug can not be empty. Please use --plan")
					}

					hostName := ctx.String("hostname")
					if hostName == "" {
						return errors.New("Hostname can not be empty. Please use --hostname")
					}

					label := ctx.String("label")
					if label == "" {
						return errors.New("Label can not be empty. Please use --label")
					}

					imgSlug := ctx.String("image")
					if imgSlug == "" {
						return errors.New("Image slug can not be empty. Please use --image")
					}

					sshSlug := ctx.String("ssh")

					req := domains.CreateInstanceRequest{
						LocationSlug: locSlug,
						PlanSlug:     planSlug,
						Hostname:     hostName,
						Label:        label,
						ImageSlug:    imgSlug,
						SSHSlug:      sshSlug,
					}

					if sshSlug != "" {
						req.SSHSlug = sshSlug
					}

					// Create Instance
					err := c.sdk.CreateInstance(&req)
					if err != nil {
						return err
					}

					fmt.Println("Instance successfully created!")

					return nil
				},
			},
			{
				Name: "destroy",
				Action: func(ctx *cli.Context) error {
					// Data Validation
					identifier := ctx.Args().First()
					if identifier == "" {
						return errors.New("Identifier can not be empty. Please use `instance destroy <identifier>`")
					}

					// Destroy Instance
					err := c.sdk.DeleteInstance(identifier)
					if err != nil {
						return err
					}

					fmt.Printf("Instance %s successfully destroyed!\n", identifier)

					return nil
				},
			},
			{
				Name: "start",
				Action: func(ctx *cli.Context) error {
					// Data Validation
					identifier := ctx.Args().First()
					if identifier == "" {
						return errors.New("Identifier can not be empty. Please use `instance start <identifier>`")
					}

					// Start Instance
					err := c.sdk.PowerOnInstance(identifier)
					if err != nil {
						return err
					}

					fmt.Printf("Instance %s startup has been queued.\n", identifier)

					return nil
				},
			},
			{
				Name: "shutdown",
				Action: func(ctx *cli.Context) error {
					// Data Validation
					identifier := ctx.Args().First()
					if identifier == "" {
						return errors.New("Identifier can not be empty. Please use `instance shutdown <identifier>`")
					}

					// Shutdown Instance
					err := c.sdk.PowerOffInstance(identifier)
					if err != nil {
						return err
					}
					fmt.Printf("Instance %s will be shut down shortly.\n", identifier)

					return nil
				},
			},
			{
				Name: "reboot",
				Action: func(ctx *cli.Context) error {
					// Data Validation
					identifier := ctx.Args().First()
					if identifier == "" {
						return errors.New("Identifier can not be empty. Please use `instance reboot <identifier>`")
					}

					// Reboot Instance
					err := c.sdk.RebootInstance(identifier)
					if err != nil {
						return err
					}

					fmt.Printf("Instance %s successfully Rebooted!\n", identifier)

					return nil
				},
			},
			{
				Name: "reset-password",
				Action: func(ctx *cli.Context) error {
					// Data Validation
					identifier := ctx.Args().First()
					if identifier == "" {
						return errors.New("Identifier can not be empty. Please use `instance reset-password <identifier>`")
					}

					// Reset Root Password Instance
					newPass := ctx.Args().Get(2)

					err := c.sdk.ResetPasswordInstance(identifier, newPass)
					if err != nil {
						return err
					}

					fmt.Printf("Instance %s Password Successfully Reset!\n", identifier)

					return nil
				},
			},
		},
	}
}
