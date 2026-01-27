package config

import (
	"encoding/json"
	"os"
)

func write(cfg Config) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	_ = os.Remove(filePath)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	encoder.SetIndent("", "    ")

	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}
