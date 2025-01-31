package main

import (
	"encoding/json"
	"os"
)

type Pattern struct {
	Label    string   `json:"label,omitempty"`
	Syllable []string `json:"syllable"`
}

type Config struct {
	Sounds  map[string][]string `json:"sounds"`
	Pattern []Pattern           `json:"patterns"`
}

func loadConfig(path string) (Config, error) {
	configFile, err := os.ReadFile(path)

	if err != nil {
		return Config{}, err
	}

	var config Config
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
