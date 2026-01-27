package config

import (
	"encoding/json"
	"os"
)

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	configfile, err := os.Open(filePath)
	if err != nil {
		return Config{}, err
	}
	defer configfile.Close()

	exportconfig := Config{}
	decoder := json.NewDecoder(configfile)
	err = decoder.Decode(&exportconfig)
	if err != nil {
		return Config{}, err
	}

	return exportconfig, nil
}
