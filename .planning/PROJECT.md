# Weather CLI

**Context:**
A robust Command Line Interface (CLI) application in Go that fetches and displays current weather data based on the user's location.

**Core Value:**
Quick, beautiful, cross-platform weather checking directly from the terminal without needing to specify a city every time.

## Target Audience
Developers and terminal users who want fast weather updates on Windows and Linux/macOS.

## Requirements

### Validated
- Cross-platform compatible (Windows, Linux, macOS).
- Native Go execution (no external wrappers like Cobra or Viper).

### Active
- Setup initial Go project structure
- Basic CLI argument parsing
- Auto-detect user location via IP address API
- Fetch current weather data from external API
- Display formatted output with terminal colors
- Local configuration for default location

### Out of Scope
- GUI interface — this is strictly a CLI tool.

## Key Decisions
| Decision | Rationale | Outcome |
|----------|-----------|---------|
| No third-party CLI frameworks | Keeping the binary small and mastering standard library | — Pending |

## Evolution

This document evolves at phase transitions and milestone boundaries.

**After each phase transition** (via `/gsd-transition`):
1. Requirements invalidated? → Move to Out of Scope with reason
2. Requirements validated? → Move to Validated with phase reference
3. New requirements emerged? → Add to Active
4. Decisions to log? → Add to Key Decisions
5. "What This Is" still accurate? → Update if drifted

**After each milestone** (via `/gsd-complete-milestone`):
1. Full review of all sections
2. Core Value check — still the right priority?
3. Audit Out of Scope — reasons still valid?
4. Update Context with current state

---
*Last updated: 2026-05-27 after initialization*
