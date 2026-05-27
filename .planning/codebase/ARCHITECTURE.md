<!-- refreshed: 2026-05-27 -->
# Architecture

**Analysis Date:** 2026-05-27

## System Overview

```text
┌─────────────────────────────────────────────────────────────┐
│                        CLI Entry Point                       │
│                      `cmd/weather/main.go`                   │
└────────┬──────────────────────┬──────────────────────┬──────┘
         │                      │                      │
         ▼                      ▼                      ▼
┌──────────────────┐   ┌──────────────────┐   ┌──────────────────┐
│  Config Manager  │   │ Location Service │   │  Weather Client  │
│ `internal/config`│   │`internal/location`│  │`internal/weather`│
└────────┬─────────┘   └────────┬─────────┘   └────────┬─────────┘
         │                      │                      │
         └──────────────────────┼──────────────────────┘
                                ▼
                       ┌──────────────────┐
                       │  Display Engine  │
                       │`internal/display`│
                       └──────────────────┘
```

## Component Responsibilities

| Component | Responsibility | File |
|-----------|----------------|------|
| Entry Point | Parses flags, initializes app, handles fatal errors | `cmd/weather/main.go` |
| Config Manager | Configuration management (env vars, default config file) | `internal/config/` |
| Location Service | IP-based location detection logic | `internal/location/` |
| Weather Client | API client for fetching weather data | `internal/weather/` |
| Display Engine | Terminal output rendering and color coding for OS | `internal/display/` |

## Pattern Overview

**Overall:** CLI Command Pattern

**Key Characteristics:**
- Strict adherence to Standard Go Project Layout.
- OS-specific terminal formatting via `runtime.GOOS`.
- Built solely with the Go standard library, no third-party framework wrappers.

## Layers

**Entry Layer:**
- Purpose: CLI interface and initialization
- Location: `cmd/weather/`

**Business Logic Layer:**
- Purpose: Orchestrate external API calls and parse data
- Location: `internal/location/`, `internal/weather/`

**Presentation Layer:**
- Purpose: Output results to console
- Location: `internal/display/`

## Data Flow

### Primary Request Path

1. User invokes `weather-cli` or `weather-cli --city "Paris"` (`cmd/weather/main.go`)
2. Config Manager reads `WEATHER_API_KEY` and defaults (`internal/config/`)
3. If no city provided, Location Service fetches location via IP (`internal/location/`)
4. Weather Client fetches weather for location (`internal/weather/`)
5. Display Engine formats output and prints to stdout (`internal/display/`)

## Key Abstractions

**[Not defined yet]**

## Entry Points

**CLI Command:**
- Location: `cmd/weather/main.go`
- Triggers: User execution in terminal
- Responsibilities: Parse flags, initiate core workflow

## Architectural Constraints

- **Dependencies:** Strictly rely on standard library (`net/http`, `encoding/json`, `flag`). Do not import external frameworks.
- **Cross-Compilation:** Use `runtime.GOOS` for OS-specific logic and `path/filepath` for file operations.

## Anti-Patterns

### Using Third-Party CLI Frameworks
**What happens:** Developer imports Cobra or Viper.
**Why it's wrong:** Violates project constraints in `AGENTS.md`.
**Do this instead:** Use the `flag` standard library.

## Error Handling

**Strategy:** Standard Go error returns.
**Patterns:** Fail fast at the entry point for fatal errors.

## Cross-Cutting Concerns

**Logging:** Standard `log` package or `fmt`.
**Validation:** Basic argument validation via `flag`.

---

*Architecture analysis: 2026-05-27*
