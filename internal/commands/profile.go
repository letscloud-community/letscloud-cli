package commands

import (
	"os"

	"github.com/letscloud-community/letscloud-cli/internal/pkg/writer"
	"github.com/urfave/cli/v2"
)

func (c *Commands) profileCmd() *cli.Command {
	return &cli.Command{
		Name:  "profile",
		Usage: "Show your Profile Info",
		Action: func(ctx *cli.Context) error {
			prof, err := c.sdk.Profile()
			if err != nil {
				return err
			}

			wr := writer.New(os.Stdout)
			wr.WriteHeader("NAME", "COMPANY", "E-MAIL", "BALANCE")
			wr.WriteData(prof.Name, prof.CompanyName, prof.Email, prof.Currency+prof.Balance)
			return wr.Flush()
		},
	}
}
