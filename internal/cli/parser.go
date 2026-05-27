package cli

import (
	"flag"
	"fmt"
	"os"
)

// Config holds the parsed command-line arguments.
type Config struct {
	City        string
	ShowVersion bool
	ShowHelp    bool
}

// ParseFlags parses the command line arguments and returns a Config.
// If there is an error during parsing or if the user requests help/version,
// this function may exit the program or print to stdout/stderr.
func ParseFlags() Config {
	var cfg Config

	// We use a new FlagSet to ensure clean testing and separation
	fs := flag.NewFlagSet("weather-cli", flag.ContinueOnError)

	fs.StringVar(&cfg.City, "city", "", "City name to fetch weather for (e.g., \"London\")")
	fs.BoolVar(&cfg.ShowVersion, "version", false, "Print version information and exit")
	fs.BoolVar(&cfg.ShowHelp, "help", false, "Print usage information and exit")

	// Custom usage function
	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of weather-cli:\n")
		fs.PrintDefaults()
	}

	err := fs.Parse(os.Args[1:])
	if err != nil {
		if err == flag.ErrHelp {
			os.Exit(0)
		}
		fmt.Fprintf(os.Stderr, "Error parsing flags: %v\n", err)
		os.Exit(1)
	}

	// Handle the standard help flag (which might not trigger ErrHelp if defined manually)
	if cfg.ShowHelp {
		fs.Usage()
		os.Exit(0)
	}

	return cfg
}
