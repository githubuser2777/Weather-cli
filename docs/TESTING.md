<!-- generated-by: gsd-doc-writer -->
## TESTING.md

**Test framework and setup:**
- Framework: `go test` (standard library)
- Setup: Ensure `go mod tidy` has been run.

**Running tests:**
```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...
```

**Writing new tests:**
- Append `_test.go` to the file name you are testing (e.g., `main_test.go`).
- Use the standard `testing` package.

**Coverage requirements:**
No coverage threshold configured.

**CI integration:**
No CI/CD pipeline detected.
