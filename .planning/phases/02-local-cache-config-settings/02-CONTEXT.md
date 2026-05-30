# Phase 2: Local Cache & Config Settings - Context

**Gathered:** 2026-05-30
**Status:** Ready for planning

<domain>
## Phase Boundary

This phase implements API response caching and CLI configuration management flags.

### In Scope
- Lightweight local JSON cache at `~/.weather-cli/cache.json` with a 10-minute expiration (TTL) for weather API responses.
- Environment variables or CLI parameter checks to bypass caching (`--force` or `-f`).
- Explicit configuration view and write flags: `--config-show` and `--config-set`.
- Unit tests for caching mechanisms and config flag processors.

### Out of Scope
- Support for multiple config formats (JSON is the only format).
- Distributed cache configurations (local only).

</domain>

<decisions>
## Implementation Decisions

### Local API Caching
- File location: Store local cache at `~/.weather-cli/cache.json` alongside the config.json.
- TTL duration: 10 minutes. If the current time is less than 10 minutes from the cached timestamp, use cache.
- Bypassing: Add a `--force` flag (or check environment variable `FORCE_CACHE_BYPASS=1`) to skip cache and force-fetch fresh API data.

### Configuration Extensions
- Flag `--config-show`: Parse config and print keys and values clearly in the console.
- Flag `--config-set`: Accept format `key=value`. Valid keys: `city` (setting `default_city` in JSON) and `unit` (setting `unit` in JSON).
- Validation: Ensure values for `unit` are either `celsius` or `fahrenheit`.

### the agent's Discretion
- Exact formatting of `--config-show` console list.

</decisions>

<code_context>
## Existing Code Insights

### Reusable Assets
- `config.LoadConfig` and `config.SaveConfig` inside `internal/config/config.go`.
- `weather.FetchWeather` inside `internal/weather/weather.go`.

### Established Patterns
- Single entrypoint flag parsing inside `cmd/weather/main.go`.

### Integration Points
- Create a new package or add functions in `internal/config/` (e.g. `cache.go` and `config.go`) to handle caching logic.

</code_context>

<specifics>
## Specific Ideas

- Check if `~/.weather-cli` directory creation is isolated and robust.
- Cache should be skipped automatically if coordinates differ.

</specifics>

<deferred>
## Deferred Ideas

- None.

</deferred>
