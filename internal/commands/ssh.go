package commands

import (
	"errors"
	"fmt"
	"os"

	"github.com/letscloud-community/letscloud-cli/internal/pkg/writer"
	"github.com/urfave/cli/v2"
)

// Manage your SSH Keys
//
// Mostly getting SSH key, Deleting, Creating etc.
func (c Commands) sshKeyCmd() *cli.Command {
	return &cli.Command{
		Name:  "ssh-key",
		Usage: "Manage your SSH Keys",
		Subcommands: []*cli.Command{
			{
				Name: "list",
				Action: func(ctx *cli.Context) error {
					// Get all the SSH Keys
					sshKeys, err := c.sdk.SSHKeys()
					if err != nil {
						return err
					}

					wr := writer.New(os.Stdout)
					wr.WriteHeader("SLUG", "TITLE", "PUBLIC_KEY")
					for _, v := range sshKeys {
						wr.WriteData(v.Slug, v.Title, v.PublicKey)
					}
					return wr.Flush()
				},
			},
			{
				Name: "create",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name: "title",
					},
					&cli.StringFlag{
						Name: "key",
					},
				},
				Action: func(ctx *cli.Context) error {
					title := ctx.String("title")
					// Validate SSH Key title
					if title == "" {
						return errors.New("Please Provide a title for your SSH Key `ssh-key create --title>`")
					}

					key := ctx.String("key")

					// Create New SSH Key
					sk, err := c.sdk.NewSSHKey(title, key)
					if err != nil {
						return err
					}

					// Return friendly message
					fmt.Printf("SSH Key %s successfully created!\n", sk.Title)

					if key == "" {
						fmt.Printf("Here's your Private Key, Please Store it in a Safe Place\n%s\n", sk.PrivateKey)
					}

					return nil
				},
			},
			{
				Name: "delete",
				Action: func(ctx *cli.Context) error {
					slug := ctx.Args().First()
					// Validate SSH Key title
					if slug == "" {
						return errors.New("Please Provide a slug of your SSH Key `ssh-key delete <slug>`")
					}
					// Delete existing SSH Key
					err := c.sdk.DeleteSSHKey(slug)
					if err != nil {
						return err
					}

					// Return friendly message
					fmt.Printf("SSH Key %s successfully destroyed\n", ctx.Args().First())

					return nil
				},
			},
		},
	}
}
