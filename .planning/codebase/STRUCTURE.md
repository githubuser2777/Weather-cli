# Codebase Structure

**Analysis Date:** 2026-05-27

## Directory Layout

```
weather-cli/
├── cmd/
│   └── weather/          # Entry point
├── internal/
│   ├── config/           # Configuration management
│   ├── location/         # IP-based location detection
│   ├── weather/          # API client for fetching weather data
│   └── display/          # Terminal output rendering
├── docs/                 # Documentation and feature tracking
├── AGENTS.md             # AI Agent Mission & Directives
└── AI_POLICY_RULES.md    # AI Coding Rules
```

## Directory Purposes

**cmd/:**
- Purpose: Main applications for the project.
- Contains: `main.go` entry points.

**internal/:**
- Purpose: Private application and library code.
- Contains: Business logic packages that shouldn't be imported by other repositories.

**docs/:**
- Purpose: Project documentation.
- Contains: Markdown files for architecture, features, roadmap.

## Key File Locations

**Entry Points:**
- `cmd/weather/main.go`: Main CLI executable.

**Configuration:**
- `internal/config/`: Configuration handling logic.
- `AGENTS.md`, `AI_POLICY_RULES.md`: Agent configuration and policies.

**Core Logic:**
- `internal/weather/`: Weather fetching.
- `internal/location/`: Location detecting.
- `internal/display/`: Output rendering.

## Naming Conventions

**Files:**
- Standard Go conventions: snake_case for filenames (e.g., `main.go`, `weather_client.go`).

**Directories:**
- Standard Go conventions: short, concise lowercase names (e.g., `cmd`, `internal`, `weather`).

## Where to Add New Code

**New Feature:**
- Primary code: `internal/[package]/`
- CLI binding: `cmd/weather/main.go`

**New API Integration:**
- Implementation: `internal/[package]/` (e.g., `internal/forecast/`)

**Utilities:**
- Shared helpers: `internal/util/` or specific package

## Special Directories

**cmd/:**
- Purpose: Executables
- Generated: No
- Committed: Yes

**internal/:**
- Purpose: Private packages
- Generated: No
- Committed: Yes

---

*Structure analysis: 2026-05-27*
