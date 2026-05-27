# Phase 1: Foundation (MVP) - Context

**Gathered:** 2026-05-27
**Status:** Ready for planning

<domain>
## Phase Boundary

Establish the foundational structure for the Weather CLI, setting up the repository according to the standard Go project layout and implementing basic CLI argument parsing using the standard library.

</domain>

<decisions>
## Implementation Decisions

### CLI Structure
- Use `internal/cli` to keep `main.go` thin. This ensures testability and keeps the entrypoint clean.

</decisions>

<canonical_refs>
## Canonical References

**Downstream agents MUST read these before planning or implementing.**

### Project Rules
- `AI_POLICY_RULES.md` — Core constraints (standard library only, cross-platform)

</canonical_refs>

<specifics>
## Specific Ideas

- The `internal/cli` package should export a `Config` struct that holds parsed flags like `--city`, `--help`, and `--version`.

</specifics>

<deferred>
## Deferred Ideas

None — Phase is scoped correctly.

</deferred>

---

*Phase: 01-foundation-mvp*
*Context gathered: 2026-05-27*
