# Contributing to catnet

First off, thank you for considering contributing to `catnet`! It's people like you that make the CatNet ecosystem a great tool.

## Rules
1. **Do not implement scanning logic in this repository.** All network operations must go through `catnet-core`. If you need a new feature that involves scanning, DNS, or parsing, please open an issue or PR in `github.com/mendsec/catnet-core` first.
2. The CLI is just a thin wrapper.

## Local Development
To test multiplatform builds locally, you can use:
```bash
GOOS=windows go build ./...
GOOS=linux go build ./...
GOOS=darwin go build ./...
```

To inject a local development version during build:
```bash
go build -ldflags "-X github.com/mendsec/catnet/internal/cli.Version=dev-local" ./cmd/catnet
```

## Testing
Run the integration tests before submitting PRs:
```bash
go test -race -v ./...
```

## Branching & Commit Policy (DevSecOps)
- **Collaboration Branch**: The `develop` branch is the primary integration branch for development. All contributor pull requests must target `develop`.
- **Main Branch Restrictions**: The `main` branch is reserved for stable releases. Pull requests targeting `main` must:
  - Come exclusively from `develop`.
  - Be automatically created by `github-actions[bot]`.
- **Signed Commits**: All commits in pull requests targeting `main` must be signed (GPG or SSH signature) to ensure verification and integrity.
