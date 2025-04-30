package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	VERSION = "v1.2.0"
)

func main() {
	// Initializes LetsCloud SDK
	cmdDict, err := initLetscloud()
	if err != nil {
		log.Fatal(err)
	}

	// Create new CLI App
	app := cli.NewApp()
	app.Name = "Official LetsCloud CLI"
	app.Version = VERSION
	app.Usage = "This cli helps you to manage your LetsCloud infrastructure from your terminal"
	app.Commands = cmdDict.Commands()
	app.Before = func(ctx *cli.Context) error {
		if ctx.Args().Len() > 1 {
			if ctx.Args().First() == "api-key" && ctx.Args().Get(1) == "set" {
				// allow
				return allow()
			}
		}

		if cmdDict.Sdk().APIKey() != "" {
			return allow()
		}

		return notAllow()
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func allow() error {
	return nil
}

func notAllow() error {
	os.Exit(0)

	return nil
}
