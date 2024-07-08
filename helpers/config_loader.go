package helpers

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config interface {
	DefaultConfig() any
}

// Check if database.yml exists, if not create it with default config
func CheckOrParseConfig[T Config](filename string) (T, error) {
	var options T
	configPath := filepath.Join("configs", filename)
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// File doesn't exist, create it with default config

		defaultConfig := options.DefaultConfig()

		data, err := yaml.Marshal(&defaultConfig)
		if err != nil {
			return options, err
		}

		err = os.MkdirAll(filepath.Dir(configPath), 0755)
		if err != nil {
			return options, err
		}

		err = os.WriteFile(configPath, data, 0644)
		if err != nil {
			return options, err
		}

		return options, fmt.Errorf("%s created in configs folder. Please update with your details", filename)
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return options, err
	}

	err = yaml.Unmarshal(data, &options)
	if err != nil {
		return options, err
	}
	return options, nil
}
