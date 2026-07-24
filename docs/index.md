---
layout: default
title: catnet — Scriptable Network Scanner CLI
nav_order: 1
description: A fast, scriptable network scanner for the command line. Built in Go. Zero engine dependencies. Made for pipelines.
---

> **catnet** — A fast, scriptable network scanner for the command line. Built in Go. Zero engine dependencies. Made for pipelines.

[![Release](https://img.shields.io/github/v/release/catnet-io/catnet)](https://github.com/catnet-io/catnet/releases)
[![CI](https://github.com/catnet-io/catnet/actions/workflows/ci.yml/badge.svg)](https://github.com/catnet-io/catnet/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## What catnet does

| 🔍 Discover              | 🔒 Enumerate                | 📤 Export                |
|---------------------------|-----------------------------|--------------------------|
| ICMP ping sweep           | TCP port scanning           | JSON · CSV · XML         |
| ARP resolution            | Reverse DNS lookup          | Pipeline-ready stdout    |

## Install in 30 seconds

**macOS / Linux (Homebrew):**
```bash
brew install catnet-io/tap/catnet
```

**Windows (Scoop):**
```powershell
scoop bucket add catnet https://github.com/catnet-io/scoop-bucket
scoop install catnet
```

**Linux / macOS (Binary Download):**
```bash
curl -sSL https://github.com/catnet-io/catnet/releases/latest/download/catnet_Linux_x86_64.tar.gz | tar xz
sudo mv catnet /usr/local/bin/
catnet version
```

**Using Go:**
```bash
go install github.com/catnet-io/catnet/cmd/catnet@latest
```

## Quick CLI & Flag Reference

### `catnet scan [targets] [flags]`

| Flag | Shorthand | Default | Description |
|:---|:---:|:---:|:---|
| `--ports` | `-p` | `22,80,443,139,445,3389` | Comma-separated TCP ports to probe |
| `--threads` | `-t` | `64` | Concurrency thread count |
| `--ping-timeout` | — | `1000` | ICMP ping timeout (in ms) |
| `--port-timeout` | — | `500` | TCP connect timeout (in ms) |
| `--timeout` | — | (none) | Scan deadline (e.g. `30s`) |
| `--no-ports` | — | `false` | ICMP ping sweep only |
| `--format` | — | `human` | Output format: `human` or `json` |
| `--output` | `-o` | (none) | Write scan output file with `0600` permissions (JSON format) |
| `--quiet` | `-q` | `false` | Suppress progress output to stderr |
| `--no-color` | — | `false` | Disable ANSI color codes |

### `catnet export [input.json] [flags]`

| Flag | Shorthand | Default | Description |
|:---|:---:|:---:|:---|
| `--format` | `-f` | (required) | Target export format: `json`, `csv`, or `xml` |
| `--output` | `-o` | (none) | Write exported results to file with `0600` permissions |

### Exit Codes Contract

| Exit Code | Classification | Meaning |
|:---:|:---|:---|
| `0` | Success | Scan completed successfully. |
| `1` | Input Error | Invalid targets, unsupported format, or flag error. |
| `2` | Runtime Error | Scan failure or network socket error. |
| `130` | Interrupted | Scan cancelled via Ctrl+C / SIGINT or timeout. |

## Designed for pipelines

`catnet` outputs human-readable progress indicators to `stderr` while streaming structured results (or silent quiet outputs) to `stdout`.

### Basic JSON scan with jq integration
```bash
catnet scan 192.168.1.0/24 --format json | jq '.devices[] | select(.isAlive) | {ip, hostname, openPorts}'
```

### Quiet mode (ideal for scripting/CI)
```bash
catnet scan 192.168.1.0/24 --format json --quiet
```

### Custom port scanning
```bash
# Scan specific ports
catnet scan 192.168.1.0/24 --ports 22,80,443,8080

# Skip port scanning entirely (ping sweep only)
catnet scan 192.168.1.0/24 --no-ports
```

### Save and re-export
```bash
catnet scan 192.168.1.0/24 --format json -o result.json
catnet export result.json --format csv -o result.csv
```

---

## Part of the CatNet Ecosystem

CatNet is a complete network scanning suite designed for terminal users, automation scripts, and graphical desktops.

| | Repository | Role | Description |
|---|---|---|---|
| ⚙️ | [catnet-io/engine](https://github.com/catnet-io/engine) | Shared scanning engine | High-performance, asynchronous scanning library in Go. |
| 💻 | [catnet-io/catnet](https://github.com/catnet-io/catnet) | **Scriptable CLI** | This CLI client, optimized for terminal pipelining. |
| 🖥️ | [catnet-io/app](https://github.com/catnet-io/app) | Desktop GUI | Cross-platform desktop application (Wails + React) with local SQLite history and scan comparison diffing. |
| 📟 | [catnet-io/tui](https://github.com/catnet-io/tui) | Terminal UI | Keyboard-centric interactive Terminal UI built with Bubble Tea. |

---

- [Full documentation on the Wiki](https://github.com/catnet-io/catnet/wiki)
- [GitHub Repository](https://github.com/catnet-io/catnet)
- [Report an Issue](https://github.com/catnet-io/catnet/issues/new)
