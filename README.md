# GoListen 🎧

A lightweight DNS traffic sniffer written in Go that monitors and displays DNS queries on your network interface.

## Overview

GoListen is a simple yet powerful tool that captures and displays DNS queries in real-time from your network interface. It uses the `gopacket` library to capture network packets and specifically focuses on DNS traffic, making it useful for network debugging and monitoring purposes.

## Features

- 🔍 Real-time DNS query monitoring
- 🌐 Automatic network interface detection
- 📝 Clean and readable output format
- 🚀 Minimal resource usage
  
## Installation

### Option 1: Direct Installation (Recommended)

Install directly using Go:

```bash
go install github.com/iamhouser/golisten@latest
```

After installation, the `golisten` binary will be available in your `$GOPATH/bin` directory. Make sure your `$GOPATH/bin` is in your system's PATH.

### Usage

Simply run:

```bash
sudo golisten
```

### Option 2: Building from Source

If you want to modify the code or build from source:

1. Clone the repository:
```bash
git clone https://github.com/iamhouser/golisten.git
cd golisten
```

2. Build and install:
```bash
go build
```
