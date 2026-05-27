# Weather CLI

A robust, cross-platform Command Line Interface (CLI) application for fetching and displaying current weather data, written in Go.

## Features

- **Auto-Location Detection:** Detects your location automatically via IP if no city is specified.
- **Accurate Weather Data:** Fetches real-time weather information (temperature, humidity, conditions) from reliable APIs.
- **Cross-Platform:** Runs seamlessly on both Linux and Windows.
- **Beautiful Output:** Clean, color-coded, and well-aligned terminal display.

## Getting Started

### Prerequisites
- [Go](https://golang.org/doc/install) (1.20+ recommended)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/weather-cli.git
   cd weather-cli
   ```
2. Build the project:
   ```bash
   go build -o weather-cli cmd/weather/main.go
   ```

### Configuration
Set up your API key via an environment variable. Do NOT hardcode secrets!
```bash
export WEATHER_API_KEY="your_api_key_here"
```
*(Windows users can use `set WEATHER_API_KEY=...` or `$env:WEATHER_API_KEY="..."` in PowerShell)*

## Documentation
- [Development Roadmap](docs/ROADMAP.md)
- [Features Tracker](docs/features.md)
- [Architecture & Design](docs/ARCHITECTURE.md)
- [Agent Directives](AGENTS.md)
- [AI Policy Rules](AI_POLICY_RULES.md)

## Contributing
Please see our [Contributing Guidelines](docs/CONTRIBUTING.md) for details on how to propose changes.
