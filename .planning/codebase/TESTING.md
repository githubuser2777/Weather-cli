# Testing Patterns

**Analysis Date:** 2026-05-27

## Test Framework

**Runner:**
- Standard Go `testing` package

**Run Commands:**
```bash
go test ./...              # Run all tests
go test -v ./...           # Verbose mode
go test -cover ./...       # Coverage
```

## Test File Organization

**Location:**
- Co-located with implementation (e.g., `weather.go` -> `weather_test.go`)

**Naming:**
- `*_test.go`

## Test Structure

**Suite Organization:**
```go
func TestWeather(t *testing.T) {
    // Standard table-driven tests
}
```

## Mocking

**Framework:** None (Manual mocking using interfaces)

**What to Mock:**
- External HTTP calls to Weather/IP APIs.
- OS-specific formatting if it interacts directly with stdout.

## Coverage

**Requirements:** None enforced yet.

---

*Testing analysis: 2026-05-27*
