package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestConfigLoadAndSave(t *testing.T) {
	// Set up temporary home directory for testing
	tempDir := t.TempDir()
	t.Setenv("HOME", tempDir)        // Unix
	t.Setenv("USERPROFILE", tempDir) // Windows

	// 1. Load config when file does not exist (should return defaults)
	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig() failed when config file was missing: %v", err)
	}

	if cfg.Unit != "celsius" {
		t.Errorf("Expected default unit 'celsius', got '%s'", cfg.Unit)
	}
	if cfg.DefaultCity != "" {
		t.Errorf("Expected default city empty, got '%s'", cfg.DefaultCity)
	}

	// 2. Save config
	newCfg := Config{
		DefaultCity: "Hanoi",
		Unit:        "fahrenheit",
	}
	err = SaveConfig(newCfg)
	if err != nil {
		t.Fatalf("SaveConfig() failed: %v", err)
	}

	// Verify file is actually created
	expectedPath := filepath.Join(tempDir, ".weather-cli", "config.json")
	if _, err := os.Stat(expectedPath); os.IsNotExist(err) {
		t.Errorf("Expected config file to be created at %s, but it does not exist", expectedPath)
	}

	// 3. Load config again (should return saved values)
	cfg, err = LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig() failed after save: %v", err)
	}

	if cfg.DefaultCity != "Hanoi" {
		t.Errorf("Expected city 'Hanoi', got '%s'", cfg.DefaultCity)
	}
	if cfg.Unit != "fahrenheit" {
		t.Errorf("Expected unit 'fahrenheit', got '%s'", cfg.Unit)
	}
}

func TestLoadConfigInvalidJSON(t *testing.T) {
	tempDir := t.TempDir()
	t.Setenv("HOME", tempDir)
	t.Setenv("USERPROFILE", tempDir)

	configDir := filepath.Join(tempDir, ".weather-cli")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		t.Fatalf("failed to create mock config dir: %v", err)
	}

	configPath := filepath.Join(configDir, "config.json")
	if err := os.WriteFile(configPath, []byte("{invalid-json"), 0644); err != nil {
		t.Fatalf("failed to write invalid config file: %v", err)
	}

	_, err := LoadConfig()
	if err == nil {
		t.Error("Expected error when loading malformed JSON config, but got nil")
	}
}
