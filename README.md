<p align="center">
  <img src="assets/logo.png" alt="zp logo" width="200"/>
</p>

<h1 align="center">zp</h1>

<p align="center">
  Lightweight, fast CLI tool written in Go to monitor open ports and inspect the processes using them.
</p>

---
## ✨ Features

- Built with Go
- Fast and lightweight
- Designed as a CLI tool
- Systemd-friendly
- Minimal dependencies
- Identify the process bound to each port
- View process information (PID, port, etc.)
- Kill a process directly from the CLI
- See how many connections hit a specific port

## 🚧 Upcoming Features

- Resource usage monitoring (CPU, memory)
- Better tracking of how many clients hit a port
- Improved statistics and summaries

## 📦 Requirements

- Go 1.20+
- Linux
- Root privileges (required for eBPF and low-level system access)

## 🏗 Architecture

```
internals/
├── ebpf/
│   ├── C source files for eBPF programs
│   └── Compiled eBPF object files
│
├── loadebpf/
│   └── Loads eBPF object files and pins them to the kernel
│
├── process/
│   └── Reads data from pinned maps to identify processes
```
## ⚙️ How It Works

- `zp` uses **eBPF programs** to hook into kernel-level networking events such as socket creation and port binding.
- These eBPF programs collect lightweight metadata (PID, port, protocol, etc.) without affecting system performance.
- Collected data is stored inside **eBPF maps**, which are **pinned to the kernel** so they persist independently of the CLI process.
- The Go CLI attaches to these pinned maps and **reads data without reloading the eBPF programs**.
- By correlating socket information with process IDs, `zp` can identify which process is using a specific port.
- Optional actions like **killing a process** are executed safely from user space based on this data.

This design allows `zp` to be fast, safe, and suitable for long-running system monitoring.
## 🚀 Installation

> **Note:** Installation will be simplified in future releases.

For now, build from source:

```bash
git clone https://github.com/moundher122/zp.git
cd zp
go build -o zp
```

Run the CLI:

```bash
sudo ./zp
```

## ▶️ Usage

Display help:

```bash
sudo ./zp --help
```

> Root access is required to load and read eBPF programs.

## 🛠 Development

Run locally:

```bash
sudo go run .
```

Format code:

```bash
go fmt ./...
```
## Architecture

For detailed architecture information, please refer to [Architecture.md](./Architecture.md).

## 📄 License

MIT License
