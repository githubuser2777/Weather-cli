<!-- generated-by: gsd-doc-writer -->
## CONFIGURATION.md

**Environment variables:**
| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `WEATHER_API_KEY` | Optional | - | API key for fetching weather data |

**Config file format:**
Local config may be stored at `~/.weather-cli/config.json`.
```json
{
  "default_location": "New York",
  "unit": "C"
}
```

**Required vs optional settings:**
No settings are strictly required for the application to start, but some features (like API calls) may fail without `WEATHER_API_KEY`.

**Defaults:**
- Location: Auto-detected via IP API
- Unit: Celsius

**Per-environment overrides:**
Not applicable for this CLI tool.
