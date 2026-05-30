package location

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDetectLocationSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp := IPApiResponse{
			Status:  "success",
			Country: "Vietnam",
			City:    "Hanoi",
			Lat:     21.0245,
			Lon:     105.8412,
		}
		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	oldAPIURL := apiURL
	apiURL = server.URL + "/"
	defer func() { apiURL = oldAPIURL }()

	loc, err := DetectLocation()
	if err != nil {
		t.Fatalf("DetectLocation() failed: %v", err)
	}

	if loc.City != "Hanoi" {
		t.Errorf("Expected city 'Hanoi', got '%s'", loc.City)
	}
	if loc.Country != "Vietnam" {
		t.Errorf("Expected country 'Vietnam', got '%s'", loc.Country)
	}
	if loc.Lat != 21.0245 || loc.Lon != 105.8412 {
		t.Errorf("Expected lat/lon 21.0245, 105.8412, got %f, %f", loc.Lat, loc.Lon)
	}
}

func TestDetectLocationFailure(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp := IPApiResponse{
			Status:  "fail",
			Message: "invalid query",
		}
		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	oldAPIURL := apiURL
	apiURL = server.URL + "/"
	defer func() { apiURL = oldAPIURL }()

	_, err := DetectLocation()
	if err == nil {
		t.Error("Expected error from DetectLocation when API returns failure status, got nil")
	}
}

func TestDetectLocationHttpError(t *testing.T) {
	// Point to an invalid address
	oldAPIURL := apiURL
	apiURL = "http://invalid-endpoint-that-should-fail-immediately.local"
	defer func() { apiURL = oldAPIURL }()

	_, err := DetectLocation()
	if err == nil {
		t.Error("Expected error from DetectLocation when HTTP call fails, got nil")
	}
}
