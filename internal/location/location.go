package location

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Location represents the geographic location of the user.
type Location struct {
	City        string  `json:"city"`
	Country     string  `json:"country"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Coordinates string  // Formatted string if needed
}

// IPApiResponse maps the JSON response from ip-api.com
type IPApiResponse struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	City        string  `json:"city"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Message     string  `json:"message"`
}

// DetectLocation uses ip-api.com to detect the user's location based on their IP address.
func DetectLocation() (Location, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get("http://ip-api.com/json/")
	if err != nil {
		return Location{}, fmt.Errorf("failed to fetch IP geolocation: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Location{}, fmt.Errorf("IP geolocation API returned status: %d", resp.StatusCode)
	}

	var ipResp IPApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&ipResp); err != nil {
		return Location{}, fmt.Errorf("failed to decode IP geolocation response: %w", err)
	}

	if ipResp.Status != "success" {
		return Location{}, fmt.Errorf("IP geolocation failed: %s", ipResp.Message)
	}

	return Location{
		City:        ipResp.City,
		Country:     ipResp.Country,
		Lat:         ipResp.Lat,
		Lon:         ipResp.Lon,
		Coordinates: fmt.Sprintf("%f, %f", ipResp.Lat, ipResp.Lon),
	}, nil
}
