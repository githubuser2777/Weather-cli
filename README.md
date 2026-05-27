# Weather CLI

Weather CLI is a small, cross-platform Go command line tool for checking current weather from the terminal. It can use your IP address to detect a location automatically, or it can fetch weather for a city you provide.

The project uses only the Go standard library and public JSON APIs.

## Features

- Detects your current city and coordinates with IP geolocation when no city is provided.
- Looks up city coordinates with the Open-Meteo geocoding API.
- Fetches current temperature, humidity, and weather conditions from Open-Meteo.
- Supports Celsius and Fahrenheit output.
- Supports optional multi-day forecasts.
- Renders a formatted terminal weather panel with ANSI colors.
- Saves default city and unit preferences in a local JSON config file.

## Requirements

- Go 1.22 or newer
- Network access for location, geocoding, and weather API calls

No API key is required for the current Open-Meteo based implementation.

## Build

```bash
go build -o weather-cli ./cmd/weather
```

On Windows PowerShell:

```powershell
go build -o weather-cli.exe .\cmd\weather
```

## Usage

Run with automatic IP-based location detection:

```bash
./weather-cli
```

Run for a specific city:

```bash
./weather-cli --city "London"
```

Use Fahrenheit:

```bash
./weather-cli --city "New York" --unit fahrenheit
```

Show a forecast:

```bash
./weather-cli --city "Tokyo" --forecast 3
```

Save your current options as defaults:

```bash
./weather-cli --city "Paris" --unit celsius --save-config
```

Show the version:

```bash
./weather-cli --version
```

## Flags

| Flag | Description |
| --- | --- |
| `--city` | City name to fetch weather for. If omitted, the CLI uses config defaults or IP geolocation. |
| `--unit` | Temperature unit. Accepted values: `celsius`, `fahrenheit`. |
| `--forecast` | Number of forecast days to request, from 1 to 7. |
| `--save-config` | Save the selected city and unit to the local config file. |
| `--version` | Print the CLI version. |

## Configuration

Saved preferences are stored at:

```text
~/.weather-cli/config.json
```

Example:

```json
{
  "default_city": "Paris",
  "unit": "celsius"
}
```

Configuration is optional. Without a saved city, Weather CLI attempts IP-based location detection.

## Development

Run all tests:

```bash
go test ./...
```

Format Go files:

```bash
gofmt -w ./cmd ./internal
```

## Documentation

- [Getting Started](docs/GETTING-STARTED.md)
- [Configuration](docs/CONFIGURATION.md)
- [Development Roadmap](docs/ROADMAP.md)
- [Features Tracker](docs/FEATURES.md)
- [Architecture & Design](docs/ARCHITECTURE.md)
- [Testing](docs/TESTING.md)
- [Contributing](docs/CONTRIBUTING.md)
- [Agent Directives](AGENTS.md)
- [AI Policy Rules](AI_POLICY_RULES.md)
