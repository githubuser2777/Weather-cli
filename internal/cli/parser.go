package cli

import (
	"flag"
	"fmt"
	"io"
)

// Config holds all the command-line options parsed from flags
type Config struct {
	City        string
	ShowVersion bool
	ShowHelp    bool
}

// ParseFlags parses command-line flags and returns a Config struct.
// It takes args slice (typically os.Args[1:]) and an io.Writer for custom output.
func ParseFlags(args []string, output io.Writer) (*Config, error) {
	config := &Config{}
	
	// Create a new FlagSet to allow better testability (vs flag.CommandLine)
	fs := flag.NewFlagSet("weather-cli", flag.ContinueOnError)
	fs.SetOutput(output)

	fs.StringVar(&config.City, "city", "", "City name to fetch weather for (e.g., 'London')")
	fs.BoolVar(&config.ShowVersion, "version", false, "Print version information")
	
	// Override the default Usage
	fs.Usage = func() {
		fmt.Fprintf(output, "Usage of weather-cli:\n")
		fs.PrintDefaults()
	}

	// Parse the arguments
	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	return config, nil
}
