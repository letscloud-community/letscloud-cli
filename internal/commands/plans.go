package commands

import (
	"errors"
	"os"
	"strconv"

	"github.com/letscloud-community/letscloud-cli/internal/pkg/writer"
	"github.com/urfave/cli/v2"
)

// interact with plans
//
// mostly getting all the plans in a specific location
func (c Commands) plansCmd() *cli.Command {
	return &cli.Command{
		Name:  "plans",
		Usage: "Show All Plans by Location",
		Action: func(ctx *cli.Context) error {
			// checking if location slug is valid
			slug := ctx.Args().First()

			if slug == "" {
				return errors.New("please provide a slug for the location `plans <location_slug>`")
			}

			// get all the location plans
			locs, err := c.sdk.LocationPlans(slug)
			if err != nil {
				return err
			}

			// init new tab writer
			wr := writer.New(os.Stdout)
			// write headers
			wr.WriteHeader("SLUG", "CORE", "MEMORY", "SSD", "BANDWIDTH", "MONTHLY VALUE")

			// write data
			for _, v := range locs {
				wr.WriteData(v.Slug, v.Core, strconv.Itoa(v.Memory)+" MB", strconv.Itoa(v.Disk)+" GB", strconv.Itoa(v.Bandwidth)+" GB", v.Shortcode+v.MonthlyValue)
			}

			return wr.Flush()
		},
	}
}
