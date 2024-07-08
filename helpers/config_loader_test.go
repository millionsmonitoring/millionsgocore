package helpers_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/millionsmonitoring/millionsgocore/helpers"
	"gopkg.in/yaml.v2"
)

// MockConfig implements the Config interface for testing
type MockConfig struct {
	Name string `yaml:"name"`
	Age  int    `yaml:"age"`
}

func (m MockConfig) DefaultConfig() any {
	return MockConfig{
		Name: "John Doe",
		Age:  30,
	}
}

func TestCheckOrParseConfig(t *testing.T) {
	configsDir := "configs"
	t.Run("Config file doesn't exist", func(t *testing.T) {
		filename := "test_config.yml"
		config, err := helpers.CheckOrParseConfig[MockConfig](filename)

		if err == nil {
			t.Errorf("Expected error, got nil")
		}

		expectedErrMsg := "test_config.yml created in configs folder. Please update with your details"
		if err.Error() != expectedErrMsg {
			t.Errorf("Expected error message %q, got %q", expectedErrMsg, err.Error())
		}

		// Check if file was created
		if _, err := os.Stat(filepath.Join(configsDir, filename)); os.IsNotExist(err) {
			t.Errorf("Config file was not created")
		}

		// Check default values
		if config.Name != "" || config.Age != 0 {
			t.Errorf("Default config not set correctly. Got %+v", config)
		}

		if err := os.Remove(filepath.Join(configsDir, filename)); err != nil {
			t.Fatalf("Failed to remove test config file: %v", err)
		}
	})

	t.Run("Config file exists and is valid", func(t *testing.T) {
		filename := "existing_config.yml"
		existingConfig := MockConfig{Name: "Jane Doe", Age: 25}
		data, _ := yaml.Marshal(existingConfig)
		err := os.WriteFile(filepath.Join(configsDir, filename), data, 0644)
		if err != nil {
			t.Fatalf("Failed to write test config file: %v", err)
		}

		config, err := helpers.CheckOrParseConfig[MockConfig](filename)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if config.Name != "Jane Doe" || config.Age != 25 {
			t.Errorf("Config not parsed correctly. Got %+v, want %+v", config, existingConfig)
		}
	})

	t.Run("Config file exists but is invalid", func(t *testing.T) {
		filename := "invalid_config.yml"
		invalidData := []byte("invalid: yaml: content")
		err := os.WriteFile(filepath.Join(configsDir, filename), invalidData, 0644)
		if err != nil {
			t.Fatalf("Failed to write invalid test config file: %v", err)
		}

		_, err = helpers.CheckOrParseConfig[MockConfig](filename)
		if err == nil {
			t.Errorf("Expected error for invalid config, got nil")
		}
	})
}
