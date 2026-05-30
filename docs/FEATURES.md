# Weather CLI - Features Roadmap

This document tracks the capabilities of the CLI application. The AI agent (Antigravity) MUST update this file as new features are implemented.

## Implemented Features
- [x] Setup initial Go project structure (`cmd/weather/main.go`).
- [x] Basic CLI argument parsing using the `flag` standard library (e.g., `--help`, `--city`).
- [x] Auto-detect user location via IP address API (e.g., ipapi.co or ipify).
- [x] Fetch current weather data from an external API (e.g., OpenWeatherMap or WeatherAPI).
- [x] Parse JSON response and map it to Go structs.
- [x] Display clean, formatted weather output (Temperature, Humidity, Conditions) in the terminal.
- [x] Implement color-coded terminal output (e.g., blue for cold, red for hot) compatible with Windows Command Prompt and Linux.
- [x] Allow saving a default location to a local config file (e.g., `~/.weather-cli/config.json`) so the IP detection isn't needed every run.
- [x] Add forecasting capabilities (e.g., `--forecast 3` for a 3-day forecast).
- [x] Support multiple units (Celsius/Fahrenheit) via flag (e.g., `--unit C`).
- [x] Add offline-first unit tests for core packages (`weather`, `display`, `location`, `config`) mocking HTTP integrations.
- [x] Dynamically adjust console box-width to prevent layout border wrapping for long city names.

## Work in Progress
*(None currently. Ready for new features!)*

## Planned Features
*(None currently. Ready for new features!)*

## Backlog / Ideas
- [ ] Add caching for weather API responses to avoid rate limits.
- [ ] Add hourly forecast display.
- [ ] Add weather alerts (e.g., severe weather warnings).
- [ ] Support custom API keys via config/ENV.
