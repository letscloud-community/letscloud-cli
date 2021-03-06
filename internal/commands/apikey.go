package commands

import (
	"errors"
	"fmt"

	"github.com/letscloud-community/letscloud-cli/internal/helpers"
	"github.com/urfave/cli/v2"
)

// Show or Set your API Key
//
// Mostly getting API Key, seting etc.
func (c *Commands) apiKeyCmd() *cli.Command {
	return &cli.Command{
		Name:  "api-key",
		Usage: "Show or Set your API Key",
		Subcommands: []*cli.Command{
			{
				Name: "set",
				Action: func(ctx *cli.Context) error {
					akValue := ctx.Args().First()
					// Validate API Key value
					if akValue == "" {
						return errors.New("Please provide a API Key value `letscloud api-key set <value>`")
					}

					// Save API Key from args passed to the CLI
					err := helpers.SaveAPIKey(akValue)
					if err != nil {
						return err
					}

					err = c.sdk.SetAPIKey(akValue)
					if err != nil {
						return err
					}

					// Return friendly message once done
					fmt.Println("Your API Key has been saved. Thanks!")

					return nil
				},
			},
			{
				Name: "show",
				Action: func(ctx *cli.Context) error {
					// Get API Key
					tok, err := helpers.GetAPIKey()
					if err != nil {
						return err
					}

					fmt.Printf("Your API Key: %s\n", tok)

					return nil
				},
			},
		},
	}
}
