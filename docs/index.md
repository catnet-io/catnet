---
layout: default
title: catnet — Scriptable Network Scanner CLI
nav_order: 1
description: A fast, scriptable network scanner for the command line. Built in Go. Zero engine dependencies. Made for pipelines.
---

> **catnet** — A fast, scriptable network scanner for the command line. Built in Go. Zero engine dependencies. Made for pipelines.

## What catnet does

| 🔍 Discover              | 🔒 Enumerate                | 📤 Export                |
|---------------------------|-----------------------------|--------------------------|
| ICMP ping sweep           | TCP port scanning           | JSON · CSV · XML         |
| ARP resolution            | Reverse DNS lookup          | Pipeline-ready stdout    |

## Install in 30 seconds

**Linux / macOS:**
```bash
curl -sSL https://github.com/catnet-io/catnet/releases/latest/download/catnet_Linux_x86_64.tar.gz | tar xz
sudo mv catnet /usr/local/bin/
catnet version
```

**Windows:**
Download `catnet_Windows_x86_64.zip` from [Releases](https://github.com/catnet-io/catnet/releases), extract, and add to PATH.

**Using Go:**
```bash
go install github.com/catnet-io/catnet/cmd/catnet@latest
```

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

### Explore the Interfaces

#### 📟 Terminal UI (TUI)
For an interactive, keyboard-driven dashboard inside the console, check out [catnet-io/tui](https://github.com/catnet-io/tui). It connects to the same underlying scanning engine to render real-time progress bars and network status panels.

#### 🖥️ Desktop Application
For engineers who prefer a graphical dashboard, history persistence, and visual comparison of changes over time, check out the [catnet-io/app](https://github.com/catnet-io/app) desktop wrapper built using Wails and React.

---

## Badges & Status

[![Release](https://img.shields.io/github/v/release/catnet-io/catnet)](https://github.com/catnet-io/catnet/releases)
[![CI](https://github.com/catnet-io/catnet/actions/workflows/ci.yml/badge.svg)](https://github.com/catnet-io/catnet/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

- [Full documentation on the Wiki](https://github.com/catnet-io/catnet/wiki)
- [GitHub Repository](https://github.com/catnet-io/catnet)
- [Report an Issue](https://github.com/catnet-io/catnet/issues/new)
