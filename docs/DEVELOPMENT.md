<!-- generated-by: gsd-doc-writer -->
## DEVELOPMENT.md

**Local setup:**
1. Clone the repository.
2. Run `go mod tidy` to download dependencies.
3. Use `go run cmd/weather/main.go` for local testing.

**Build commands:**
| Command | Description |
|---------|-------------|
| `go build -o weather-cli cmd/weather/main.go` | Builds the CLI executable |
| `go run cmd/weather/main.go` | Runs the CLI without compiling an executable |
| `go fmt ./...` | Formats all Go source files |

**Code style:**
- Use standard `go fmt` for formatting code.

**Branch conventions:**
No convention documented.

**PR process:**
- Please ensure tests pass before submitting.
- Follow the guidelines in [CONTRIBUTING.md](CONTRIBUTING.md).
