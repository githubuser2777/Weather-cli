# Weather CLI - Features Roadmap

This document tracks the capabilities of the CLI application. The AI agent (Antigravity) MUST update this file as new features are implemented.

## Implemented Features
*(None yet. Awaiting initial scaffolding.)*

## Work in Progress
- [ ] Setup initial Go project structure (`cmd/weather/main.go`).
- [ ] Basic CLI argument parsing using the `flag` standard library (e.g., `--help`, `--city`).
- [ ] Auto-detect user location via IP address API (e.g., ipapi.co or ipify).

## Planned Features
- [ ] Fetch current weather data from an external API (e.g., OpenWeatherMap or WeatherAPI).
- [ ] Parse JSON response and map it to Go structs.
- [ ] Display clean, formatted weather output (Temperature, Humidity, Conditions) in the terminal.
- [ ] Implement color-coded terminal output (e.g., blue for cold, red for hot) compatible with Windows Command Prompt and Linux.
- [ ] Allow saving a default location to a local config file (e.g., `~/.weather-cli/config.json`) so the IP detection isn't needed every run.

## Backlog / Ideas
- [ ] Add forecasting capabilities (e.g., `--forecast 3` for a 3-day forecast).
- [ ] Support multiple units (Celsius/Fahrenheit) via flag (e.g., `--unit C`).
