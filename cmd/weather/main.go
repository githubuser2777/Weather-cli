package main

import (
	"fmt"
	"os"

	"github.com/githubuser2777/Weather-cli/internal/cli"
)

const version = "0.1.0"

func main() {
	config := cli.ParseFlags()

	if config.ShowVersion {
		fmt.Printf("weather-cli version %s\n", version)
		os.Exit(0)
	}

	// Phase 1 Dummy Output
	if config.City != "" {
		fmt.Printf("Weather for %s: (Not implemented yet - Phase 1 Foundation)\n", config.City)
	} else {
		// In future phases, this will trigger IP Geolocation detection
		fmt.Println("No city provided. IP Geolocation auto-detect: (Not implemented yet - Phase 1 Foundation)")
	}

	// Exit gracefully without panics
	os.Exit(0)
}
