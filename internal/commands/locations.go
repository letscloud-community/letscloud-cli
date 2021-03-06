package commands

import (
	"os"
	"strconv"

	"github.com/letscloud-community/letscloud-cli/internal/helpers"
	"github.com/letscloud-community/letscloud-cli/internal/pkg/writer"
	"github.com/urfave/cli/v2"
)

// get all the locations
func (c Commands) locationsCmd() *cli.Command {
	return &cli.Command{
		Name:  "locations",
		Usage: "Show All Locations",
		Action: func(ctx *cli.Context) error {
			locs, err := c.sdk.Locations()
			if err != nil {
				return err
			}

			// init new tab writer
			wr := writer.New(os.Stdout)
			// write headers
			wr.WriteHeader(helpers.GetStructHeaders(locs[0])...)

			for _, v := range locs {
				wr.WriteData(v.Slug, v.Country, v.City, strconv.FormatBool(v.Available))
			}

			return wr.Flush()
		},
	}
}
