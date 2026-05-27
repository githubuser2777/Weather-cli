# Weather CLI - Development Roadmap

This document outlines the phased development plan for the Weather CLI project. It provides a strategic view of how the project will evolve from scaffolding to a feature-rich application.

---

## 🟢 Phase 1: Foundation (MVP)
**Goal:** Establish the project structure and basic CLI interactions.

- [ ] **Project Scaffolding:** Create the standard Go directory structure (`cmd/weather`, `internal/`).
- [ ] **CLI Argument Parsing:** Implement the `flag` package to handle basic inputs like `--city`, `--help`, and `--version`.
- [ ] **Basic Output & Error Handling:** Ensure the CLI can run, print dummy text, and handle basic execution errors gracefully (no panics).

---

## 🟡 Phase 2: Core Data Integration
**Goal:** Connect to external services to fetch real data.

- [ ] **IP Geolocation Detection:** Implement logic in `internal/location` to auto-detect the user's city/coordinates if `--city` is not provided (using services like ip-api or ipify).
- [ ] **Weather API Client:** Implement the HTTP client in `internal/weather` to call a weather service (e.g., OpenWeatherMap).
- [ ] **Data Parsing:** Unmarshal the JSON API responses into strongly typed Go structs.

---

## 🟠 Phase 3: Presentation & User Experience
**Goal:** Make the output beautiful and cross-platform compatible.

- [ ] **Terminal Formatting:** Create clean, aligned text output for temperature, humidity, and weather conditions.
- [ ] **Cross-Platform Colors:** Implement color-coded output (e.g., Blue for cold, Red for hot, Yellow for sunny) ensuring it works on both Linux terminals and Windows (CMD/PowerShell).
- [ ] **Robust Error Messages:** Format network errors or "city not found" errors into user-friendly terminal messages.

---

## 🔵 Phase 4: Advanced Features & Polish
**Goal:** Add convenience features that make the tool a daily driver.

- [ ] **Local Configuration:** Allow users to save a default location to a config file (`~/.weather-cli/config.json`) to bypass IP detection and speed up execution.
- [ ] **Unit Support:** Add flags to switch between Metric and Imperial units (e.g., `--unit fahrenheit`).
- [ ] **Extended Forecasts:** Add a flag (e.g., `--forecast 3`) to show a multi-day weather forecast.
- [ ] **ASCII Art / Icons:** Optionally integrate simple terminal icons (e.g., ⛅, 🌧️) if the terminal supports UTF-8.

---

> *Note: For a granular list of individual tasks and their current status, please refer to [features.md](features.md).*
