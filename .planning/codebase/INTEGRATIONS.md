# External Integrations

**Analysis Date:** 2026-05-27

## APIs & External Services

**Weather API:**
- OpenWeatherMap or WeatherAPI (planned) - Fetch current weather data
  - SDK/Client: Built-in `net/http` (planned)
  - Auth: `WEATHER_API_KEY` (planned)

**Geolocation API:**
- ipapi.co or ipify (planned) - Auto-detect user location via IP address
  - SDK/Client: Built-in `net/http` (planned)

## Data Storage

**Databases:**
- None

**File Storage:**
- Local filesystem for saving default location to `~/.weather-cli/config.json` (planned)

**Caching:**
- None

## Authentication & Identity

**Auth Provider:**
- Custom (API Keys for third-party services)

## Monitoring & Observability

**Error Tracking:**
- None

**Logs:**
- Built-in logging or fmt (planned)

## CI/CD & Deployment

**Hosting:**
- Local execution

**CI Pipeline:**
- None detected

## Environment Configuration

**Required env vars:**
- `WEATHER_API_KEY` (planned)

**Secrets location:**
- Environment variables or local config

## Webhooks & Callbacks

**Incoming:**
- None

**Outgoing:**
- None

---

*Integration audit: 2026-05-27*
