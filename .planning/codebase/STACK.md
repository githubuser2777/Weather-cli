# Technology Stack

**Analysis Date:** 2026-05-27

## Languages

**Primary:**
- Go (planned) - Core logic and CLI implementation

**Secondary:**
- None detected

## Runtime

**Environment:**
- Native (compiled binary for Windows/Linux)

**Package Manager:**
- Go modules (`go mod`) (planned)
- Lockfile: missing (no go.mod yet)

## Frameworks

**Core:**
- Standard library (`net/http`, `encoding/json`, `flag`)
- No external frameworks like Cobra or Viper allowed per AGENTS.md

**Testing:**
- Standard library (`testing`) (planned)

**Build/Dev:**
- `go build` (planned)

## Key Dependencies

**Critical:**
- None detected yet.

**Infrastructure:**
- None detected yet.

## Configuration

**Environment:**
- `WEATHER_API_KEY` environment variable (planned)
- Local config file `~/.weather-cli/config.json` (planned)

**Build:**
- Standard Go build process (planned)

## Platform Requirements

**Development:**
- Go SDK

**Production:**
- Windows (CMD/PowerShell)
- Linux/macOS

---

*Stack analysis: 2026-05-27*
