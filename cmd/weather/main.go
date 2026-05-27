package main

import (
	"flag"
	"fmt"
	"os"

	"githubuser2777/Weather-cli/internal/config"
	"githubuser2777/Weather-cli/internal/display"
	"githubuser2777/Weather-cli/internal/location"
	"githubuser2777/Weather-cli/internal/weather"
)

func main() {
	// Parse CLI flags
	cityFlag := flag.String("city", "", "Specify a city to fetch the weather for")
	unitFlag := flag.String("unit", "", "Unit of temperature: 'celsius' or 'fahrenheit'")
	forecastFlag := flag.Int("forecast", 0, "Number of upcoming days to forecast (1-7)")
	saveConfigFlag := flag.Bool("save-config", false, "Save current settings as default")
	versionFlag := flag.Bool("version", false, "Print the version of Weather CLI")

	flag.Parse()

	if *versionFlag {
		fmt.Println("Weather CLI v1.0.0")
		os.Exit(0)
	}

	// 1. Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		display.PrintError(fmt.Errorf("failed to load config: %v", err))
		// Non-fatal, continue with defaults
	}

	// 2. Cascade logic for Unit
	finalUnit := "celsius"
	if *unitFlag != "" {
		finalUnit = *unitFlag
	} else if cfg.Unit != "" {
		finalUnit = cfg.Unit
	}

	if finalUnit != "celsius" && finalUnit != "fahrenheit" {
		display.PrintError(fmt.Errorf("invalid unit '%s'. Must be 'celsius' or 'fahrenheit'", finalUnit))
		os.Exit(1)
	}

	// 3. Cascade logic for City
	finalCity := *cityFlag
	if finalCity == "" {
		finalCity = cfg.DefaultCity
	}

	var targetLat, targetLon float64
	var targetName string

	if finalCity == "" {
		// IP Detection
		loc, err := location.DetectLocation()
		if err != nil {
			display.PrintError(err)
			os.Exit(1)
		}
		targetLat = loc.Lat
		targetLon = loc.Lon
		targetName = fmt.Sprintf("%s, %s", loc.City, loc.Country)
	} else {
		// Geocoding
		lat, lon, name, err := weather.GeocodeCity(finalCity)
		if err != nil {
			display.PrintError(err)
			os.Exit(1)
		}
		targetLat = lat
		targetLon = lon
		targetName = name
	}

	// 4. Save Config if requested
	if *saveConfigFlag {
		newCfg := config.Config{
			DefaultCity: finalCity,
			Unit:        finalUnit,
		}
		if err := config.SaveConfig(newCfg); err != nil {
			display.PrintError(fmt.Errorf("failed to save config: %v", err))
		} else {
			fmt.Println("Config saved successfully.")
		}
	}

	// 5. Fetch weather data
	data, err := weather.FetchWeather(targetLat, targetLon, finalUnit, *forecastFlag)
	if err != nil {
		display.PrintError(err)
		os.Exit(1)
	}

	// 6. Render beautiful ASCII widget
	display.RenderWeather(targetName, data)
}
