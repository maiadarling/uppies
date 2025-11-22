package commands

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	Token string `json:"token"`
}

var GlobalConfig Config

func LoadConfig() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	configDir := filepath.Join(home, ".uppies")
	configPath := filepath.Join(configDir, "config")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// Create default config
		GlobalConfig = Config{
			Token: "",
		}
		// Ensure directory exists
		os.MkdirAll(configDir, 0755)
		// Write to file
		file, err := os.Create(configPath)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		json.NewEncoder(file).Encode(GlobalConfig)
	} else {
		// Read from file
		file, err := os.Open(configPath)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		json.NewDecoder(file).Decode(&GlobalConfig)
	}
}