# Architecture Specification — catnet CLI

**Target:** `github.com/catnet-io/catnet`  
**Binary Name:** `catnet`  
**Language:** Go 1.26.5 (Pure Go, `CGO_ENABLED=0`)  
**Ecosystem Dependencies:** `github.com/catnet-io/engine`  

---

## 1. Architectural Philosophy & Boundaries

`catnet` is the command-line interface frontend for the CatNet network scanning ecosystem. Its design is governed by three primary tenets:

1. **Zero Internal Scanning Logic:** The CLI contains no raw socket manipulation, ICMP handling, ARP tables, or TCP connection routines. All scanning, host discovery, reverse DNS resolution, and port enumeration logic are strictly isolated in `github.com/catnet-io/engine`.
2. **First-Class Pipeline Interoperability:** Output is strictly separated between streams:
   - `stdout`: Machine-readable scan results (JSON) or clean export tables.
   - `stderr`: Human-oriented progress indicators, scanning metadata, warnings, and error diagnostics.
3. **Deterministic Exit Codes:** Every termination path maps to a documented, contractually enforced numeric exit code.

```
                  +--------------------------------+
                  |         catnet (CLI)           |
                  |  cmd/catnet/main.go            |
                  +---------------+----------------+
                                  |
                                  v
                  +--------------------------------+
                  |       internal/cli/root.go     |
                  |     (Cobra Command Dispatcher) |
                  +---+------------+-----------+---+
                      |            |           |
            +---------+            |           +---------+
            |                      |                     |
            v                      v                     v
    +---------------+      +---------------+     +---------------+
    |    scanCmd    |      |   exportCmd   |     |  versionCmd   |
    | (scan.go)     |      | (export.go)   |     | (version.go)  |
    +-------+-------+      +-------+-------+     +---------------+
            |                      |
            |                      |
            v                      v
    +---------------+      +---------------+
    | Output Layer  |      |   Engine      |
    | (Human/JSON)  |      | Exporter Pkg  |
    +-------+-------+      +---------------+
            |
            v
    +----------------------------------------------+
    |           catnet-io/engine (Shared)          |
    |  - engine.StartScan (Scan Execution)         |
    |  - targets.ParseRange (Target Parsing)       |
    |  - results.ScanReport (Data Model)           |
    +----------------------------------------------+
```

---

## 2. Directory & Package Layout

```
catnet/
├── cmd/
│   └── catnet/
│       └── main.go               # Application entrypoint; translates ExitError to os.Exit
├── internal/
│   └── cli/
│       ├── root.go               # Cobra root command; global flags (--no-color)
│       ├── scan.go               # 'scan' command implementation and flag binding
│       ├── export.go             # 'export' command; converts JSON scans to CSV/XML/JSON
│       ├── version.go            # 'version' command; reports build, commit, and engine version
│       ├── signals.go            # Context cancellation on OS signals (SIGINT/SIGTERM)
│       ├── errors.go             # ExitError definition and error formatting
│       └── output/
│           ├── human.go          # Human-readable progress bar and tabular device formatter
│           └── json.go           # JSON streaming event adapter
├── tests/
│   └── integration_test.go       # End-to-end subprocess integration tests
└── docs/
    ├── architecture.md           # This architecture specification
    ├── audit_202606.md           # Historical June 2026 technical audit
    ├── audit_202609.md           # September 2026 technical audit
    └── index.md                  # GitHub Pages documentation portal
```

---

## 3. Subcommand Architecture

### 3.1 Root Command (`internal/cli/root.go`)
- Initializes the Cobra root instance.
- Configures `SilenceErrors: true` and `SilenceUsage: true` so error output is routed cleanly through `cmd/catnet/main.go` without auto-printing usage help on runtime errors.
- Binds `--no-color` as a persistent flag across subcommands.

### 3.2 Scan Command (`internal/cli/scan.go`)
1. **Target Processing:** Accepts single IPs, CIDR blocks, ranges, or `'auto'` (default if no targets provided) via `targets.ParseRange`. Deduplicates input arguments.
2. **Configuration Construction:** Assembles `engine.Config` with port lists, concurrency limits (`--threads`), and timeout parameters (`--ping-timeout`, `--port-timeout`). Calls `cfg.Sanitize()`.
3. **Signal & Context Wrapping:** Wraps the command context with `WithCancelOnSignal(cmd.Context())`.
4. **Event Adapter Dispatch:** Instantiates either `HumanOutput` or `JSONOutput` based on `--format`.
5. **Execution:** Invokes `engine.StartScan(ctx, allIPs, cfg, eventHandler)`.
6. **File Persistence:** If `--output` is specified, writes JSON report to disk with owner-only `0600` permissions.

### 3.3 Export Command (`internal/cli/export.go`)
- Reads an existing scan JSON report from disk.
- Validates the `schemaVersion` (issues warning if major version > 2).
- Delegates serialization to `exporter.ExportJSON`, `exporter.ExportCSV`, or `exporter.ExportXML`.
- Outputs directly to `stdout` or writes to file with `0600` permissions.

### 3.4 Version Command (`internal/cli/version.go`)
- Displays injected build variables (`Version`, `Commit`, `Date`) populated via GoReleaser `-ldflags`.
- Inspects runtime debug metadata (`debug.ReadBuildInfo()`) to report the exact linked module version of `github.com/catnet-io/engine`.
- Supports `--short` flag for automated scripts.

---

## 4. Engine Integration & Event Streaming

### 4.1 Event Lifecycle
The CLI attaches an `engine.EventCallback` to `engine.StartScan`. As the engine processes IP targets asynchronously across worker threads, it dispatches discrete `engine.ScanEvent` items:

| Event Type | Payload | CLI Handling |
|---|---|---|
| `EventLifecycleStart` | Host count | Logs initial banner; initializes timer and table headers |
| `EventProgress` | Progress float (0.0 – 1.0) | Updates progress indicator on `stderr` |
| `EventResult` | `*results.Device` | Formats and prints host IP, MAC, hostname, status, and ports |
| `EventWarning` | Message string | Emits formatted `[WARN]` to `stderr` |
| `EventLifecycleCancel` | Message string | Emits `[CANCELLED]` banner to `stderr` |
| `EventLifecycleComplete` | Total time / summary | Finalizes table; prints elapsed scan duration |

### 4.2 Migration to Canonical `ScanStream`
The CLI currently invokes the callback-based `engine.StartScan` (suppressed with `//nolint:staticcheck`). Under Milestone M5 (*Core Engine Alignment*), this will be refactored to consume the channel-based `pkg/scan.Engine.ScanStream` once stabilized upstream.

---

## 5. Signal Handling, Cancellation & Exit Codes

### 5.1 Signal Trap
Signals are trapped via Go's `signal.NotifyContext`:
- Non-Windows: Traps `os.Interrupt` (`SIGINT`) and `syscall.SIGTERM`.
- Windows: Traps `os.Interrupt` only (preventing runtime crashes due to undefined `SIGTERM` signals).

When a signal arrives:
1. The `context.Context` is cancelled immediately.
2. The running `engine.StartScan` halts active goroutines.
3. The CLI exits with code `130` (`ExitCodeInterrupted`).

### 5.2 Exit Code Standards

| Exit Code | Constant | Meaning | Typical Causes |
|:---:|---|---|---|
| `0` | `ExitCodeSuccess` | Clean completion | Scan completed, export finished |
| `1` | `ExitCodeInputError` | Invalid user input | Bad CIDR range, missing arguments, unsupported format |
| `2` | `ExitCodeRuntimeError` | Runtime failure | Engine error, disk permission failure, unreadable input |
| `130` | `ExitCodeInterrupted` | Execution cancelled | Received SIGINT (Ctrl+C), SIGTERM, or context deadline exceeded |

---

## 6. Output Formatting & Testability

Both output adapters utilize dependency injection for streaming destinations:

- `HumanOutput`: Wraps `tabwriter.Writer`. Injected with `out io.Writer` (defaults to `os.Stdout`) and `errOut io.Writer` (defaults to `os.Stderr`). Performs terminal capability checks (`os.ModeCharDevice`) to disable ANSI color escape sequences automatically when pipes or file redirects are detected.
- `JSONOutput`: Injected with `errOut io.Writer` to output lifecycle notifications while reserving `os.Stdout` for JSON data.

This abstraction enables unit tests (`internal/cli/output/human_test.go`) to execute in parallel without mutating global file descriptors.

---

## 7. Security & DevSecOps Architecture

- **Static Binary Compilation:** Built with `CGO_ENABLED=0` for Linux, macOS, and Windows (`amd64` and `arm64`).
- **File System Defenses:** Output files (`-o`) written via `os.WriteFile` explicitly specify permission mode `0600` (readable/writable only by the owner).
- **Vulnerability Scanning:** Automated weekly and pull request scanning via `govulncheck`, pinned to immutable commit SHAs.
- **Supply Chain Integrity:** No local `replace` directives are permitted in `go.mod` on public branches; automated CI checks (`ci.yml`) enforce this constraint.
- **Commit Signing Policy:** All pull requests merging to `main` must originate from `develop`, authored by `github-actions[bot]`, with every commit cryptographically signed (`gpgsig`).
