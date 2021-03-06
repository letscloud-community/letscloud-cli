package helpers

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	noApiKeyErr = errors.New(`Use export "LETSCLOUD_API_KEY=your_api_key" or use "letscloud api-key set your_api_key"`)
)

// return the saved API Key
func GetAPIKey() (string, error) {
	// check if API KEY exists in the env
	if apiKey := os.Getenv("LETSCLOUD_API_KEY"); apiKey != "" {
		// if yes, take it and save it to file
		if err := SaveAPIKey(apiKey); err != nil {
			return "", err
		}

		// then return it
		return apiKey, nil
	}

	apiKey, err := readAPIKeyFromFile()
	if err != nil {
		return "", noApiKeyErr
	}

	if apiKey == "" {
		return "", noApiKeyErr
	}

	return apiKey, nil
}

func readAPIKeyFromFile() (string, error) {
	usrHomeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	f, err := os.Open(filepath.Join(usrHomeDir, ".letscloud"))
	if err != nil {
		return "", err
	}

	fi, err := f.Stat()
	if err != nil {
		return "", err
	}

	if fi.Size() == 0 {
		return "", errors.New("No API Key found, please save first")
	}

	var b = make([]byte, 128)

	nb, err := f.Read(b)
	if err != nil {
		return "", err
	}

	tok := string(b[:nb])

	return splitAPIKey(tok), nil
}

func splitAPIKey(tok string) string {
	st := strings.Split(tok, "=")

	if len(st) != 2 {
		return ""
	}

	return st[1]
}

//SaveAPIKey saves the API key to disk
func SaveAPIKey(value string) error {
	usrHomeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	f, err := os.Create(filepath.Join(usrHomeDir, ".letscloud"))
	if err != nil {
		return err
	}

	_, err = f.Write([]byte(fmt.Sprintf("LETSCLOUD_API_KEY=%s", value)))
	if err != nil {
		return err
	}

	return nil
}
