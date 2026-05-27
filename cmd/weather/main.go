package main

import (
	"fmt"
	"os"

	"github.com/githubuser2777/Weather-cli/internal/cli"
)

var Version = "0.1.0"

func main() {
	// Parse CLI arguments
	config, err := cli.ParseFlags(os.Args[1:], os.Stderr)
	if err != nil {
		// flag package already prints the error if ContinueOnError is used
		// or ExitOnError exits. We used ContinueOnError to allow testing.
		os.Exit(1)
	}

	if config.ShowVersion {
		fmt.Printf("weather-cli version %s\n", Version)
		os.Exit(0)
	}

	if config.City != "" {
		fmt.Printf("Weather for %s: (Not implemented yet)\n", config.City)
	} else {
		// Fallback for when we implement IP geolocation
		fmt.Println("Auto-detecting location... (Not implemented yet)")
	}
}
