package config

import (
	"errors"
	"fmt"
	"os"
)

func FindConfigPath() (string, error) {
	return findConfigDir(".")
}

func findConfigDir(currentDirectory string) (string, error) {
	entries, err := os.ReadDir(currentDirectory)
	if err != nil {
		return "", errors.New("clint.yaml not found")
	}

	for _, entry := range entries {
		if entry.Name() == CONFIG_NAME {
			return currentDirectory, nil
		}
	}

	return findConfigDir(fmt.Sprintf("../%v", currentDirectory))
}
