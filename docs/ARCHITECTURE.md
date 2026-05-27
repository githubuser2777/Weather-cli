# Software Architecture

This document describes the high-level architecture and planned Go project structure for the Weather CLI.

## Directory Structure

We adhere to standard Go project layouts to keep the codebase clean and maintainable.

```text
weather-cli/
├── cmd/
│   └── weather/
│       └── main.go         # Entry point: Parses flags, initializes app, and handles fatal errors
├── internal/
│   ├── config/             # Configuration management (env vars, default config file)
│   ├── location/           # Logic for IP-based location detection
│   ├── weather/            # API client for fetching weather data
│   └── display/            # Terminal output rendering and color coding
├── docs/                   # Documentation and feature tracking
├── AGENTS.md               # AI Agent Mission & Directives
└── AI_POLICY_RULES.md      # AI Coding Rules
```

## ⚙️ Core Components

1. **Config Manager:** Responsible for reading the `WEATHER_API_KEY` from the environment or a local config, and managing saved default locations.
2. **Location Service:** Reaches out to a public IP geolocation API to determine coordinates/city if the user doesn't pass a `--city` flag.
3. **Weather Client:** Sends an HTTP GET request to the chosen Weather API provider. Uses `net/http` and `encoding/json` to decode the response.
4. **Display Engine:** Takes the structured weather data and formats it for terminal output. It detects the OS (`runtime.GOOS`) to ensure color codes and alignments render correctly on Windows (CMD/PowerShell) and Linux/macOS.
