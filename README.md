# port-doctor 🩺🔌

`port-doctor` is a lightning-fast, zero-dependency CLI tool built in Go designed for developers. It diagnoses port conflicts instantly, inspects the exact process name and Process ID (PID) holding a port, and safely or forcefully frees it up so you can get back to coding without the headache of stuck background servers.

---

## Features

- **Port Diagnostics (`check`):** Quickly scan any port to see if it's clear or occupied.
- **Process Identification:** Reveals the exact name of the application (e.g., `node`, `python`, `java`) and its PID hogging your port.
- **Interactive & Force Cleanup (`free`):** Safely prompt before killing a blocking process, or use `-f` to instantly force-terminate it.
- **Cross-Platform Support:** Fully compatible with macOS (`lsof`), Linux (`ss` / `lsof`), and Windows (`netstat` / `taskkill`).
- **Single Binary:** Compiled entirely in Go with zero external dependencies.

---

## Project Structure

```text
port-doctor/
├── cmd/
│   └── root.go                 # CLI command router & argument parser
├── internal/
│   ├── doctor/
│   │   └── doctor.go           # Diagnostics & cleanup orchestration
│   └── osutil/
│       ├── process_darwin.go   # macOS process inspection & termination
│       ├── process_linux.go    # Linux process inspection & termination
│       └── process_windows.go  # Windows process inspection & termination
├── go.mod
└── main.go                     # Application entry point