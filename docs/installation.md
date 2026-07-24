---
layout: default
title: Installation
nav_order: 2
description: Install catnet — binary download, go install, or build from source.
---

# Installation

## Requirements

- **Go 1.26 or later** (for source builds only)
- **Linux, macOS, or Windows** (amd64 or arm64)
- **Root / administrator privileges** for ICMP scanning on some systems

## Method 1 — Package Managers (Recommended)

### Homebrew (macOS / Linux)

```bash
brew install catnet-io/tap/catnet
```

### Scoop (Windows)

```powershell
scoop bucket add catnet https://github.com/catnet-io/scoop-bucket
scoop install catnet
```

*Note: Additional Windows package managers (Chocolatey and Winget) are planned for Phase 2.*

## Method 2 — Pre-built Binary

### Automated Download & Install (Linux / macOS)

```bash
OS="$(uname -s)"
ARCH="$(uname -m)"
case "$ARCH" in
  x86_64) ARCH="x86_64" ;;
  aarch64|arm64) ARCH="arm64" ;;
  *) echo "Unsupported architecture: $ARCH" && exit 1 ;;
esac

TAG=$(curl -sSL https://api.github.com/repos/catnet-io/catnet/releases/latest | grep '"tag_name":' | cut -d'"' -f4)
ARCHIVE="catnet_${OS}_${ARCH}.tar.gz"

curl -sSLO "https://github.com/catnet-io/catnet/releases/download/${TAG}/${ARCHIVE}"
curl -sSLO "https://github.com/catnet-io/catnet/releases/download/${TAG}/checksums.txt"

grep "${ARCHIVE}" checksums.txt | sha256sum -c -
tar -xzf "${ARCHIVE}" catnet
sudo mv catnet /usr/local/bin/
rm -f "${ARCHIVE}" checksums.txt
catnet version
```

### Manual Download

Download the appropriate archive for your platform from [GitHub Releases](https://github.com/catnet-io/catnet/releases):

| Platform | Architecture | Archive |
|---|---|---|
| Linux | x86_64 (amd64) | `catnet_Linux_x86_64.tar.gz` |
| Linux | ARM64 | `catnet_Linux_arm64.tar.gz` |
| macOS | Intel (x86_64) | `catnet_Darwin_x86_64.tar.gz` |
| macOS | Apple Silicon (arm64) | `catnet_Darwin_arm64.tar.gz` |
| Windows | x86_64 / ARM64 | `catnet_Windows_x86_64.zip` / `catnet_Windows_arm64.zip` |

Verify your download against `checksums.txt`:

```bash
sha256sum -c checksums.txt --ignore-missing
```

## Method 3 — `go install`

```bash
go install github.com/catnet-io/catnet/cmd/catnet@latest
```

For reproducible builds, pin to a specific tag:

```bash
go install github.com/catnet-io/catnet/cmd/catnet@v0.4.0
```

## Method 4 — Build from Source

```bash
git clone https://github.com/catnet-io/catnet.git
cd catnet
go build -o catnet ./cmd/catnet
./catnet version
```

## Advanced Options

For build-from-source with version injection and shell completion installation, see the [Installation wiki page](https://github.com/catnet-io/catnet/wiki/Installation).

