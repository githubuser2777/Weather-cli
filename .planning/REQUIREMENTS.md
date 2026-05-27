# Requirements

## User Stories

1. As a user, I want to type `weather-cli` and get the weather for my current location without passing any arguments.
2. As a user, I want to see a clean, color-coded display so it's easy to read at a glance.
3. As a user, I want to query a specific city with `--city "London"`.
4. As a user, I want the CLI to run smoothly on my Windows PowerShell and my Linux terminal.

## Acceptance Criteria

### Core Interactions
- `weather-cli --city "Paris"` fetches weather for Paris.
- `weather-cli` auto-detects IP location and fetches weather.
- `weather-cli --help` prints usage instructions.

### Presentation
- Output includes Temperature, Humidity, and Conditions.
- Colors are used to represent conditions (e.g., Red for hot, Blue for cold).
- Windows CMD/PowerShell display colors correctly.

## Non-Functional Requirements

- **Performance:** Command must execute and display within 2 seconds.
- **Dependencies:** Standard library only (`net/http`, `encoding/json`, `flag`).

## Definition of Done

- All code is in `cmd/` and `internal/`.
- Tests are written for `internal/weather` and `internal/location`.
- Build succeeds for both `windows/amd64` and `linux/amd64`.
