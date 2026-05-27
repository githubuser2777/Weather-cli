package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Config represents user preferences.
type Config struct {
	DefaultCity string `json:"default_city"`
	Unit        string `json:"unit"` // "celsius" or "fahrenheit"
}

// getConfigPath resolves the path to ~/.weather-cli/config.json
func getConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}
	return filepath.Join(home, ".weather-cli", "config.json"), nil
}

// LoadConfig reads the configuration from disk. Returns an empty Config if it doesn't exist.
func LoadConfig() (Config, error) {
	path, err := getConfigPath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return Config{Unit: "celsius"}, nil // Default fallback
		}
		return Config{}, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("failed to parse config JSON: %w", err)
	}

	if cfg.Unit == "" {
		cfg.Unit = "celsius"
	}
	return cfg, nil
}

// SaveConfig writes the given Config to disk.
func SaveConfig(cfg Config) error {
	path, err := getConfigPath()
	if err != nil {
		return err
	}

	// Create directory if missing
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to encode config JSON: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
