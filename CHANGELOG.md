# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

> **For maintainers:** The GitHub Release body is generated
> automatically by GoReleaser from commit messages grouped by
> Conventional Commits type, plus a static install footer.
> The `[Unreleased]` section in this file is the canonical
> human-written summary and is linked from the release body.

## [Unreleased]

### Fixed
- Added Content Security Policy (CSP) meta tag to GitHub Pages documentation site (`docs/_includes/head-custom.html`).

### Changed
- Bumped GitHub Actions `actions/setup-go` to v7 and `actions/checkout` to v7 across CI workflows.

### Added
- Added hard rule in `AGENTS.md` prohibiting unauthorized dependency downgrades.

## [0.4.0] - 2026-07-18

### Added
- Scaffolding and setup of GitHub Pages landing page using the Jekyll Cayman theme.
- GoReleaser template customization with a release footer and statically-linked binary matrix.
- Linter integration (`golangci-lint`) in GitHub Actions workflows.
- Integration tests coverage for export commands error handling.
- Local development helper script `dev-replace.sh` to manage engine module replacement.

### Changed
- Updated `engine` dependency to `v0.5.1` (bringing async dispatcher, performance improvements, and unified worker paths).
- Updated minimum Go version requirement to `1.26.5`.
- Updated repository and ecosystem references to `catnet-io` organization.

### Fixed
- Fixed Cobra flag check in `export` subcommand to handle error on marking format required.
- Resolved deprecation warning for `engine.StartScan` by adding inline `nolint:staticcheck` annotations.
- Handled potential file write and signal errors in integration tests to resolve code review findings.
- Cleaned up unused JSON output helper function.

## [0.2.0] - 2026-06-23

### Changed
- `export --format` flag is now local to the export subcommand and required. It no longer inherits the root `--format` value. Existing scripts that relied on `catnet --format csv export input.json` must be updated to `catnet export input.json --format csv`. (Resolves analysis finding C8.)
- Updated catnet-core dependency from development pseudo-version to stable v0.2.0.
- All repository comments, documentation strings, and user-facing messages standardized to English. No functional changes.

### Fixed
- Integration tests no longer share Cobra flag state between cases. `TestMain` now resets `rootCmd` before each test. (Resolves analysis finding C6.)
- `os.Stdout` pipe is now always restored via `defer` in integration tests, preventing file descriptor leaks on early test failure. (Resolves analysis finding C7.)
- `TestScanCancelledByContext` rewritten as a subprocess test. Signal is sent only to the child process, eliminating risk of terminating the test runner. (Resolves analysis finding C10.)

### Added
- Unit tests for `output/human.go` and `output/json.go` using an injected `io.Writer`. Coverage now includes TTY detection fallback and color flag propagation. (Resolves analysis finding C9.)

## [0.1.0] - 2026-06-06

### Added
- Initial scaffolding of the CLI repository (`github.com/catnet-io/catnet`).
- Cobra CLI structure with `root`, `scan`, `export`, and `version` subcommands.
- Graceful cancellation handling via `context` and `os/signal` (Exit Code 130).
- Human-readable output formatting with progress bars and terminal color support.
- JSON output formatting for CI/CD and scriptability.
- Feature to re-export previous JSON scans into CSV, XML, and JSON.
- GitHub Actions CI pipelines for cross-compilation (Windows, macOS, Linux).
- GoReleaser configuration for automated binary publishing.
- Integration tests simulating End-to-End behavior via `127.0.0.1` loopback.

[unreleased]: https://github.com/catnet-io/catnet/compare/v0.4.0...HEAD
[0.4.0]: https://github.com/catnet-io/catnet/compare/v0.2.0...v0.4.0
[0.2.0]: https://github.com/catnet-io/catnet/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/catnet-io/catnet/compare/2721a3346032d02831f4f0594ad6332a57c4f145...v0.1.0
