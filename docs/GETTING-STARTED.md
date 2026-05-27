<!-- generated-by: gsd-doc-writer -->
## GETTING-STARTED.md

**Prerequisites:**
- `Go >= 1.20`

**Installation steps:**
1. Clone the repository:
   ```bash
   git clone <!-- VERIFY: repository_url -->
   cd weather-cli
   ```
2. Build the project:
   ```bash
   go build -o weather-cli cmd/weather/main.go
   ```

**First run:**
```bash
./weather-cli
```

**Common setup issues:**
- Ensure `WEATHER_API_KEY` is set in your environment if the API requires authentication.

**Next steps:**
- See [DEVELOPMENT.md](DEVELOPMENT.md) for local development setup.
- See [TESTING.md](TESTING.md) for running the test suite.
