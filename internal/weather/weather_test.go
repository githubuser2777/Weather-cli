package weather

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDecodeWeatherCode(t *testing.T) {
	tests := []struct {
		code         int
		expectedCond string
		expectedIcon string
	}{
		{0, "Clear sky", "󰖙"},
		{1, "Partly cloudy", "󰖕"},
		{2, "Partly cloudy", "󰖕"},
		{3, "Partly cloudy", "󰖕"},
		{45, "Fog", "󰖑"},
		{48, "Fog", "󰖑"},
		{51, "Drizzle", "󰖗"},
		{55, "Drizzle", "󰖗"},
		{56, "Freezing Drizzle", "󰖗"},
		{61, "Rain", "󰖖"},
		{65, "Rain", "󰖖"},
		{66, "Freezing Rain", "󰖖"},
		{71, "Snow fall", "󰖘"},
		{77, "Snow grains", "󰖘"},
		{80, "Rain showers", "󰖗"},
		{85, "Snow showers", "󰖘"},
		{95, "Thunderstorm", "󰖓"},
		{99, "Thunderstorm", "󰖓"},
		{-1, "Unknown", "󰖕"},
	}

	for _, tt := range tests {
		cond, icon := decodeWeatherCode(tt.code)
		if cond != tt.expectedCond {
			t.Errorf("For code %d, expected condition '%s', got '%s'", tt.code, tt.expectedCond, cond)
		}
		if icon != tt.expectedIcon {
			t.Errorf("For code %d, expected icon '%s', got '%s'", tt.code, tt.expectedIcon, icon)
		}
	}
}

func TestGeocodeCitySuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp := GeocodeResponse{}
		resp.Results = append(resp.Results, struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
			Name      string  `json:"name"`
			Country   string  `json:"country"`
		}{
			Latitude:  21.0245,
			Longitude: 105.8412,
			Name:      "Hanoi",
			Country:   "Vietnam",
		})
		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	oldBaseURL := geocodeBaseURL
	geocodeBaseURL = server.URL
	defer func() { geocodeBaseURL = oldBaseURL }()

	lat, lon, name, err := GeocodeCity("Hanoi")
	if err != nil {
		t.Fatalf("GeocodeCity() failed: %v", err)
	}

	if lat != 21.0245 || lon != 105.8412 {
		t.Errorf("Expected coords 21.0245, 105.8412, got %f, %f", lat, lon)
	}
	if name != "Hanoi, Vietnam" {
		t.Errorf("Expected name 'Hanoi, Vietnam', got '%s'", name)
	}
}

func TestGeocodeCityNotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp := GeocodeResponse{Results: nil}
		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	oldBaseURL := geocodeBaseURL
	geocodeBaseURL = server.URL
	defer func() { geocodeBaseURL = oldBaseURL }()

	_, _, _, err := GeocodeCity("NonExistentCity")
	if err == nil {
		t.Error("Expected error for non-existent city, got nil")
	}
}

func TestFetchWeatherSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp := WeatherResponse{}
		resp.Current.Temperature2m = 28.2
		resp.Current.RelativeHumidity = 91
		resp.Current.WeatherCode = 95
		
		resp.Daily.Time = []string{"2026-05-30", "2026-05-31"}
		resp.Daily.Temperature2mMax = []float64{32.0, 31.0}
		resp.Daily.Temperature2mMin = []float64{25.0, 24.0}
		resp.Daily.WeatherCode = []int{95, 3}

		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	oldBaseURL := weatherBaseURL
	weatherBaseURL = server.URL
	defer func() { weatherBaseURL = oldBaseURL }()

	data, err := FetchWeather(21.0245, 105.8412, "celsius", 1)
	if err != nil {
		t.Fatalf("FetchWeather() failed: %v", err)
	}

	if data.Temperature != 28.2 {
		t.Errorf("Expected temperature 28.2, got %f", data.Temperature)
	}
	if data.Humidity != 91 {
		t.Errorf("Expected humidity 91, got %d", data.Humidity)
	}
	if data.Conditions != "Thunderstorm" || data.Icon != "󰖓" {
		t.Errorf("Expected 'Thunderstorm' / '󰖓', got '%s' / '%s'", data.Conditions, data.Icon)
	}
	if data.Unit != "C" {
		t.Errorf("Expected unit symbol 'C', got '%s'", data.Unit)
	}

	if len(data.Forecast) != 1 {
		t.Errorf("Expected 1 forecast day, got %d", len(data.Forecast))
	} else {
		f := data.Forecast[0]
		if f.Date != "2026-05-31" {
			t.Errorf("Expected forecast date '2026-05-31', got '%s'", f.Date)
		}
		if f.MaxTemp != 31.0 || f.MinTemp != 24.0 {
			t.Errorf("Expected Min/Max 24.0/31.0, got %f/%f", f.MinTemp, f.MaxTemp)
		}
		if f.Conditions != "Partly cloudy" {
			t.Errorf("Expected forecast condition 'Partly cloudy', got '%s'", f.Conditions)
		}
	}
}
