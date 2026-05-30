# Weather CLI - Development Roadmap

This document outlines the phased development plan for the Weather CLI project. It provides a strategic view of how the project will evolve from scaffolding to a feature-rich application.

---

## 🟢 Phase 1: Foundation (MVP) [COMPLETE]
**Goal:** Establish the project structure and basic CLI interactions.

- [x] **Project Scaffolding:** Create the standard Go directory structure (`cmd/weather`, `internal/`).
- [x] **CLI Argument Parsing:** Implement the `flag` package to handle basic inputs like `--city`, `--help`, and `--version`.
- [x] **Basic Output & Error Handling:** Ensure the CLI can run, print dummy text, and handle basic execution errors gracefully (no panics).

---

## 🟢 Phase 2: Core Data Integration [COMPLETE]
**Goal:** Connect to external services to fetch real data.

- [x] **IP Geolocation Detection:** Implement logic in `internal/location` to auto-detect the user's city/coordinates if `--city` is not provided (using services like ip-api or ipify).
- [x] **Weather API Client:** Implement the HTTP client in `internal/weather` to call a weather service (e.g., OpenWeatherMap).
- [x] **Data Parsing:** Unmarshal the JSON API responses into strongly typed Go structs.

---

## 🟢 Phase 3: Presentation & User Experience [COMPLETE]
**Goal:** Make the output beautiful and cross-platform compatible.

- [x] **Terminal Formatting:** Create clean, aligned text output for temperature, humidity, and weather conditions.
- [x] **Cross-Platform Colors:** Implement color-coded output (e.g., Blue for cold, Red for hot, Yellow for sunny) ensuring it works on both Linux terminals and Windows (CMD/PowerShell).
- [x] **Robust Error Messages:** Format network errors or "city not found" errors into user-friendly terminal messages.

---

## 🟢 Phase 4: Advanced Features & Polish [COMPLETE]
**Goal:** Add convenience features that make the tool a daily driver.

- [x] **Local Configuration:** Allow users to save a default location to a config file (`~/.weather-cli/config.json`) to bypass IP detection and speed up execution.
- [x] **Unit Support:** Add flags to switch between Metric and Imperial units (e.g., `--unit fahrenheit`).
- [x] **Extended Forecasts:** Add a flag (e.g., `--forecast 3`) to show a multi-day weather forecast.
- [x] **ASCII Art / Icons:** Optionally integrate simple terminal icons (e.g., ⛅, 🌧️) if the terminal supports UTF-8.

---

## 🔵 Phase 5: Interactive Terminal User Interface (TUI) [IN PROGRESS]
**Goal:** Create a real-time console dashboard with interactive settings.

- [ ] **Interactive Input Processing:** Read console inputs in raw or line-buffered mode to execute dynamic key commands.
- [ ] **ANSI Dashboard Redrawing:** Use escape codes to clear screen and redeliver beautiful layouts on city, unit, and forecast changes.
- [ ] **Console Help Screen:** Implement an interactive page mapping commands and explaining caching and config storage.
- [ ] **Graceful Program Exit:** Support clean termination on quit key presses.

---

> *Note: For a granular list of individual tasks and their current status, please refer to [features.md](features.md).*
