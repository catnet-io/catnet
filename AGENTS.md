# AGENTS.md — catnet-io/catnet

This file provides persistent context for AI coding agents working in `catnet-io/catnet`.

---

## What this repository is

`catnet-io/catnet` is the CLI frontend for the CatNet ecosystem.
It is a pure consumer of `catnet-io/engine` — it contains zero scanning logic.

**Module path:** `github.com/catnet-io/catnet`  
**Binary name:** `catnet`  
**Go version:** 1.26.5  
**Engine dependency:** `github.com/catnet-io/engine` (see `go.mod` for current version)

---

## Architecture

```
cmd/catnet/main.go
  └── internal/cli/root.go         (Cobra root command, persistent flags)
       ├── internal/cli/scan.go    (scan command → engine.StartScan)
       ├── internal/cli/export.go  (export helpers)
       ├── internal/cli/signals.go (SIGINT/SIGTERM → context cancel)
       ├── internal/cli/version.go (version command)
       └── internal/cli/output/
            ├── human.go           (human-readable output handler)
            └── json.go            (JSON streaming output handler)
```

### Engine API used

This CLI uses `pkg/engine.StartScan` (callback-based API) from `catnet-io/engine`.
It does NOT use `pkg/scan.Engine.ScanStream` (channel API).
This is known and intentional for the current version.
When Milestone 5 designates `pkg/scan.Engine` as canonical, this CLI will be updated.

---

## Hard rules — never violate

1. **No scanning logic in this repository.** All discovery, port scanning, and fingerprinting
   happens in `catnet-io/engine`. This package only calls `engine.StartScan`.
2. **No CGO.** This is a pure Go binary.
3. **English only** in all Go source files.
4. **No local `replace` directives in `main` branch.** Use `scripts/dev-replace.sh on/off`.
5. **Exit codes are contracts.** See `internal/cli/errors.go`:
   - `0` — success
   - `1` — input error (invalid targets, unsupported format)
   - `2` — runtime error (scan failure)
   - `130` — interrupted (SIGINT/context cancel)
   Do not add new exit codes without updating `docs/cli-reference.md`.

---

## Conventions

### Commit messages — Conventional Commits

```
feat(scan): add --timeout flag for global scan timeout
fix(output): handle empty device list in human output
chore(deps): update engine to v0.4.0
test(integration): add integration test for CIDR range scan
docs(cli-reference): document --no-ports flag behavior
```

Scopes: `scan`, `output`, `export`, `signals`, `version`, `root`, `deps`, `ci`, `docs`.

### Changelog — Keep a Changelog

Update `CHANGELOG.md` under `[Unreleased]` for every behavioral change.

### Release — GoReleaser

Releases are automated via `.github/workflows/release.yml` + `.goreleaser.yml`.
Do not manually create release artifacts.
Trigger a release by pushing a tag: `git tag -a vX.Y.Z -m "Release vX.Y.Z"`.
Changelog section is auto-generated from Conventional Commits.

### Testing

- `tests/integration_test.go` contains end-to-end tests using `testdata/`.
- `internal/cli/output/human_test.go` tests the human output handler.
- New commands must have at least one integration test.
- Use `go test -race ./...` locally before pushing.

---

## CI requirements — all must pass before merge

- `go build ./...`
- `go test -race ./...`
- `go vet ./...`
- GoReleaser snapshot: `goreleaser release --snapshot --clean` (on release PRs)
