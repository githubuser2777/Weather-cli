# Coding Conventions

**Analysis Date:** 2026-05-27

## Naming Patterns

**Files:**
- Standard Go: snake_case for files, lowercase for directories

**Functions:**
- Standard Go: PascalCase for exported, camelCase for unexported

**Variables:**
- Standard Go: camelCase, short names for local scopes

**Types:**
- Standard Go: PascalCase for exported structs/interfaces

## Code Style

**Formatting:**
- `gofmt` or `goimports`

**Linting:**
- `golangci-lint` (recommended)

## Import Organization

**Order:**
1. Standard library
2. Blank line
3. Project internal packages

## Error Handling

**Patterns:**
- Standard Go `error` returns.
- Fail fast for fatal errors using `log.Fatalf` or returning error to `main`.

## Logging

**Framework:** `log` or `fmt` from standard library.

**Patterns:**
- Print simple output to stdout using `fmt`.
- Print errors to stderr using `log`.

## Function Design

**Size:** Keep functions small and focused on one responsibility.

**Parameters:** Prefer structs for configuration objects over many parameters.

**Return Values:** Return `(type, error)` for functions that can fail.

---

*Convention analysis: 2026-05-27*
