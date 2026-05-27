# Contributing Guidelines

Thank you for investing your time in contributing to the Weather CLI project!

## AI-Assisted Development
This project is actively developed with the help of AI coding agents (like Antigravity). All contributors (human or AI) must strictly follow the rules defined in `AI_POLICY_RULES.md` and the instructions in `AGENTS.md`.

## Code Standards
- **Standard Library Only:** Avoid introducing 3rd-party dependencies unless absolutely necessary and previously discussed.
- **Formatting:** All Go code must be formatted using `gofmt` before committing.
- **Error Handling:** Avoid panics. Return errors explicitly and format them nicely for the CLI user.

## Process for Adding Features
1. Check `docs/features.md` to pick up a task or add a new planned feature.
2. Discuss the architectural approach if it requires a new package or major refactoring.
3. Write clean, cross-platform code.
4. Update `docs/features.md` to reflect your progress (move from 'Planned' to 'Work in Progress' to 'Implemented').
