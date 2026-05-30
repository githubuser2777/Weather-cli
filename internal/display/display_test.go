package display

import (
	"bytes"
	"io"
	"os"
	"testing"
	"unicode/utf8"

	"githubuser2777/Weather-cli/internal/weather"
)

func TestFormatTempCelsius(t *testing.T) {
	tests := []struct {
		temp        float64
		expectedRaw string
	}{
		{10.0, "10.0°C"},
		{20.0, "20.0°C"},
		{30.0, "30.0°C"},
	}

	for _, tt := range tests {
		raw, col := formatTemp(tt.temp, "C")
		if raw != tt.expectedRaw {
			t.Errorf("Expected raw '%s', got '%s'", tt.expectedRaw, raw)
		}
		if col == "" {
			t.Errorf("Expected colored string, got empty")
		}
	}
}

func TestFormatTempFahrenheit(t *testing.T) {
	raw, col := formatTemp(50.0, "F")
	if raw != "50.0°F" {
		t.Errorf("Expected raw '50.0°F', got '%s'", raw)
	}
	if col == "" {
		t.Error("Expected colored fahrenheit string, got empty")
	}
}

func TestFormatHumidity(t *testing.T) {
	raw, col := formatHumidity(85)
	if raw != "85%" {
		t.Errorf("Expected raw '85%%', got '%s'", raw)
	}
	if col == "" {
		t.Error("Expected colored humidity string, got empty")
	}
}

func TestFormatConditions(t *testing.T) {
	raw, col := formatConditions("Clear sky", "󰖙")
	if raw != "󰖙 Clear sky" {
		t.Errorf("Expected raw '󰖙 Clear sky', got '%s'", raw)
	}
	if col == "" {
		t.Error("Expected colored conditions string, got empty")
	}
}

func TestRenderWeatherDynamicBoxWidth(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	longCity := "Taumatawhakatangihangakoauauotamateaturipukakapikimaungahoronukupokaiwhenuakitanatahu, New Zealand"
	data := weather.WeatherData{
		Temperature: 25.0,
		Humidity:    80,
		Conditions:  "Clear sky",
		Icon:        "󰖙",
		Unit:        "C",
	}

	RenderWeather(longCity, data)

	w.Close()
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	os.Stdout = oldStdout

	expectedHeaderLen := utf8.RuneCountInString(" Weather for: ") + utf8.RuneCountInString(longCity)
	if boxWidth < expectedHeaderLen {
		t.Errorf("Expected boxWidth to be at least %d, but got %d", expectedHeaderLen, boxWidth)
	}
}
