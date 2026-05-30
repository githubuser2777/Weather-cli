# Phase 1: Unit Testing & Robustness - Context

**Gathered:** 2026-05-30
**Status:** Ready for planning

<domain>
## Phase Boundary

This phase establishes full unit testing coverage for all internal modules of the Weather CLI and improves layout rendering robustness to prevent dynamic box layout breakage.

### In Scope
- Core testing suite for `internal/config`, `internal/location`, `internal/weather`, and `internal/display`.
- Standard library mock testing using `net/http/httptest`.
- Dynamic console box-width calculations in `internal/display` to accommodate long city names.
- Multi-byte character width verification using UTF-8 rune counts.

### Out of Scope
- Adding external CLI frameworks like Cobra or Viper.
- Setting up permanent cloud CI/CD pipelines (local testing only).
- Caching weather responses (deferred to Phase 2).

</domain>

<decisions>
## Implementation Decisions

### Unit Testing Scope and Strategy
- Package coverage: Implement Go unit tests in all core packages (`config`, `location`, `weather`, `display`).
- Mocking HTTP: Use Go's standard library `net/http/httptest` to mock remote IP geolocation and Open-Meteo weather/geocoding API servers, ensuring no real network requests are made during tests.
- Mocking Config: Use Go `t.TempDir()` in testing config files to isolate config file load/save testing.
- Target Coverage: Prioritize functional coverage (success, errors, boundary conditions, invalid inputs) over an exact coverage percentage.

### Visual Layout Polish
- Dynamic Box Width: Calculate the display width dynamically based on the longest row length to ensure right-border alignment is never broken.
- Minimum Box Width: Set a minimum default `boxWidth` of 50 to maintain a visually appealing baseline layout.
- Unicode Safety: Use `unicode/utf8.RuneCountInString` to calculate visual lengths accurately, taking into account multi-byte Unicode characters and Nerd Font symbols.

### the agent's Discretion
- Choice of mock payloads and custom table-driven inputs.
- Precise implementation details of the dynamic width padding function.

</decisions>

<code_context>
## Existing Code Insights

### Reusable Assets
- `config.LoadConfig` and `config.SaveConfig` inside `internal/config/config.go`.
- `location.DetectLocation` inside `internal/location/location.go`.
- `weather.FetchWeather` and `weather.GeocodeCity` inside `internal/weather/weather.go`.
- `display.RenderWeather` and `display.PrintError` inside `internal/display/display.go`.

### Established Patterns
- Clean CLI interactions relying solely on Go's standard library (`flag`, `net/http`, `encoding/json`).
- Handlers wrapping HTTP client operations with a 5-second timeout limit.
- Color formatting using raw ANSI escape codes.

### Integration Points
- Add testing files directly into the package folders: `internal/config/config_test.go`, `internal/location/location_test.go`, `internal/weather/weather_test.go`, `internal/display/display_test.go`.
- Modify `internal/display/display.go` to compute box width dynamically before printing header and content rows.

</code_context>

<specifics>
## Specific Ideas

- Implement table-driven tests for the `weather.decodeWeatherCode` function to ensure all weather codes map to the correct Nerd Font icon and description.
- Implement table-driven tests for formatting functions such as `formatTemp` and `formatHumidity`.

</specifics>

<deferred>
## Deferred Ideas

- Weather API response caching (deferred to Phase 2).
- Extended environment/CLI parameters for custom API keys (deferred to Phase 2).

</deferred>
