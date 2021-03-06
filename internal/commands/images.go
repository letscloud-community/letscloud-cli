package commands

import (
	"errors"
	"os"

	"github.com/letscloud-community/letscloud-cli/internal/helpers"
	"github.com/letscloud-community/letscloud-cli/internal/pkg/writer"
	"github.com/urfave/cli/v2"
)

// listing out all the images
func (c Commands) imagesCmd() *cli.Command {
	return &cli.Command{
		Name:  "images",
		Usage: "Show All the Images by Location",
		Action: func(ctx *cli.Context) error {
			// validating data
			slug := ctx.Args().First()

			if slug == "" {
				return errors.New("Please provide a slug for the image `images <location_slug>`")
			}
			// get all the locations
			locs, err := c.sdk.LocationImages(slug)
			if err != nil {
				return err
			}

			// init new tab writer
			wr := writer.New(os.Stdout)
			// write headers
			wr.WriteHeader(helpers.GetStructHeaders(locs[0])...)

			for _, v := range locs {
				wr.WriteData(v.Slug, v.Distro, v.OS)
			}

			return wr.Flush()
		},
	}
}
