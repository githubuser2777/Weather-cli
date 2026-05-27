# Phase 1: Foundation (MVP)

**Status:** Planned

## Overview
This phase establishes the foundational structure for the Weather CLI, setting up the repository according to the standard Go project layout and implementing basic CLI argument parsing using the standard library.

## Constraints
- **NO THIRD-PARTY CLI FRAMEWORKS**: Must use standard library `flag`, not Cobra/Viper.
- **CROSS-PLATFORM**: Code must compile and run flawlessly on both Windows and Linux.
- **NO PANICS**: Handle errors gracefully using standard `fmt.Errorf` and `os.Exit(1)`.

## Task List

### 1. Initialize Go Module
- Run `go mod init github.com/githubuser2777/Weather-cli`.

### 2. Scaffold Standard Project Layout
- Create `cmd/weather/main.go`.
- Create `internal/cli/parser.go`.

### 3. Implement CLI Parsing (`internal/cli/parser.go`)
- Define a `Config` struct (e.g. holding `City string`, `ShowVersion bool`).
- Use the `flag` package to parse `--city`, `--help`, and `--version`.
- Return the parsed struct and any validation errors.

### 4. Implement Main Entrypoint (`cmd/weather/main.go`)
- Call `cli.ParseFlags()`.
- Handle the `--version` output.
- Print a dummy placeholder output (e.g. `fmt.Printf("Weather for %s: (Not implemented yet)\n", config.City)`).
- Implement graceful error handling (exit code 1 on parsing failure).

### 5. Verification
- Build binary: `go build -o weather-cli ./cmd/weather`.
- Run tests: `./weather-cli --help`, `./weather-cli --version`, `./weather-cli --city "London"`.
