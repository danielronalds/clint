package directories

import (
	"errors"
	"fmt"
	"os"

	"github.com/danielronalds/clint/internal"
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
		if entry.Name() == internal.CONFIG_NAME {
			return fmt.Sprintf("%v/%v", currentDirectory, internal.CONFIG_NAME), nil
		}
	}

	return findConfigDir(fmt.Sprintf("../%v", currentDirectory))
}
