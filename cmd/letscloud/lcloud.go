package main

import (
	"fmt"

	"github.com/letscloud-community/letscloud-cli/internal/commands"
	"github.com/letscloud-community/letscloud-cli/internal/helpers"
	"github.com/letscloud-community/letscloud-go"
)

func initLetscloud() (*commands.Commands, error) {
	ak, err := helpers.GetAPIKey()
	if err != nil {
		fmt.Println("WARNING:", err)
	}

	sdkClient, err := letscloud.New(ak)
	if err != nil {
		return nil, err
	}

	return commands.New(sdkClient), nil
}
