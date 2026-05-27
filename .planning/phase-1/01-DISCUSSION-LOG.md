# Phase 1: Foundation (MVP) - Discussion Log

**Date:** 2026-05-27

## CLI Structure
**Options presented:**
- Use `internal/cli` to keep `main.go` thin (Recommended for testability)
- Keep everything in `main.go` for simplicity since the app is small

**Selection:** Use `internal/cli` to keep `main.go` thin (Recommended for testability)
**Notes:** Decided to separate concerns right from Phase 1 to make testing easier later.
