package config

import (
	"testing"

	"githubuser2777/Weather-cli/internal/weather"
)

func TestCacheSaveAndLoadSuccess(t *testing.T) {
	tempDir := t.TempDir()
	t.Setenv("HOME", tempDir)
	t.Setenv("USERPROFILE", tempDir)

	data := weather.WeatherData{
		Temperature: 28.2,
		Humidity:    91,
		Conditions:  "Thunderstorm",
		Icon:        "󰖓",
		Unit:        "C",
	}

	err := SaveCache(21.0245, 105.8412, "Hanoi, Vietnam", "C", data)
	if err != nil {
		t.Fatalf("SaveCache() failed: %v", err)
	}

	entry, ok := LoadCache(21.0245, 105.8412, "C")
	if !ok {
		t.Fatal("Expected cache hit, but got miss")
	}

	if entry.WeatherData.Temperature != 28.2 {
		t.Errorf("Expected temperature 28.2, got %f", entry.WeatherData.Temperature)
	}
	if entry.Location != "Hanoi, Vietnam" {
		t.Errorf("Expected location 'Hanoi, Vietnam', got '%s'", entry.Location)
	}
}

func TestCacheMismatch(t *testing.T) {
	tempDir := t.TempDir()
	t.Setenv("HOME", tempDir)
	t.Setenv("USERPROFILE", tempDir)

	data := weather.WeatherData{
		Temperature: 28.2,
		Unit:        "C",
	}

	_ = SaveCache(21.0245, 105.8412, "Hanoi, Vietnam", "C", data)

	_, ok := LoadCache(20.0, 105.8412, "C")
	if ok {
		t.Error("Expected cache miss for coordinate mismatch, but got hit")
	}

	_, ok = LoadCache(21.0245, 105.8412, "F")
	if ok {
		t.Error("Expected cache miss for unit mismatch, but got hit")
	}
}
