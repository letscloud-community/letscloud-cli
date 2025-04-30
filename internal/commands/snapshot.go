package commands

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/letscloud-community/letscloud-cli/internal/pkg/writer"
	"github.com/urfave/cli/v2"
)

func (c Commands) snapshotsCmd() *cli.Command {
	return &cli.Command{
		Name:  "snapshot",
		Usage: "Manage your snapshots",
		Subcommands: []*cli.Command{
			{
				Name: "list",
				Action: func(ctx *cli.Context) error {
					// Get all the Snapshots
					snapshots, err := c.sdk.Snapshots()
					if err != nil {
						return err
					}
					wr := writer.New(os.Stdout)
					wr.WriteHeader("SLUG", "LABEL", "SIZE", "OS REFERENCE", "REFERENCE", "BUILD", "LOCATIONS")

					for _, v := range snapshots {
						locationsStr := strings.Join(v.Locations, ",")
						wr.WriteData(v.Slug, v.Label, v.Size, v.OsReference, v.Reference, v.Build, locationsStr)
					}
					return wr.Flush()
				},
			},
			{
				Name:      "create",
				ArgsUsage: "<slug>",
				Usage:     "Create a new snapshot",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "instance-identifier",
						Usage: "Identifier of the instance",
					},
					&cli.StringFlag{
						Name:  "label",
						Usage: "Label for the snapshot",
					},
				},
				Action: func(ctx *cli.Context) error {
					label := ctx.String("label")
					instanceIdentifier := ctx.String("instance-identifier")
					// Validate Snapshot label
					if label == "" || instanceIdentifier == "" {
						return errors.New("please provide a label for your Snapshot `snapshot create --label <label> --instance-identifier <instance-identifier>`")
					}
					// Create New Snapshot
					snapshot, err := c.sdk.NewSnapshot(label, instanceIdentifier)
					if err != nil {
						return err
					}
					// Return friendly message
					fmt.Printf("Snapshot %s creation has been queued. This is a background process and may take several minutes to complete.", snapshot.Data.Slug)

					return nil
				},
			},
			{
				Name:      "delete",
				Usage:     "Delete an existing snapshot",
				ArgsUsage: "<slug>",
				Action: func(ctx *cli.Context) error {
					slug := ctx.Args().First()
					// Validate Snapshot slug
					if slug == "" {
						return errors.New("please provide a slug of your Snapshot `snapshot delete <slug>`")
					}
					// Delete existing Snapshot
					err := c.sdk.DeleteSnapshot(slug)
					if err != nil {
						return err
					}

					fmt.Printf("Snapshot %s successfully destroyed\n", slug)

					return nil
				},
			},
			{
				Name:      "update",
				Usage:     "Update an existing snapshot",
				ArgsUsage: "<slug>",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "label",
						Usage: "Label for the snapshot",
					},
				},
				Action: func(ctx *cli.Context) error {
					label := ctx.Value("label")
					labelStr := label.(string)
					var slug string
					if ctx.NArg() > 0 {
						slug = ctx.Args().Get(0)
					}

					if label == "" {
						return errors.New("please provide a label for your Snapshot `snapshot update <slug> --label <label>`")
					}
					// Validate Snapshot slug
					if slug == "" {
						return errors.New("please provide a slug of your Snapshot `snapshot update <slug> --label <label>`")
					}
					// Update existing Snapshot
					err := c.sdk.UpdateSnapshot(slug, labelStr)
					if err != nil {
						return err
					}

					fmt.Printf("Snapshot %s successfully updated\n", slug)

					return nil
				},
			},
		},
	}
}
