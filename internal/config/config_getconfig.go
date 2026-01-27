package config

import (
	"os"
	"path/filepath"
)

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	filePath := filepath.Join(homeDir, ".gatorconfig.json")
	return filePath, nil
}
