# 🔍 port-scanner

A fast and concurrent port scanner built with Go. Scans a target IP address for open TCP ports using up to **600 goroutines** simultaneously for high performance.

## 📁 Project Structure

```
port-scanner/
├── cmd/
│   └── main/
│       └── main.go         # Entry point of the application
├── scanner/
│   └── context.go          # Scanning logic and goroutine management
├── go.mod                  # Go module file
└── README.md               # Project documentation
```

## 🚀 Features

- Scans ports **0 to 65535** (or any custom range)
- Utilizes **600 goroutines** for efficient scanning
- Simple command-line interface
- Clean modular structure

## 🛠️ Requirements

- Go 1.18 or later

## 🧪 Usage

### Build and Run

```bash
go run cmd/main/main.go -ip=localhost -range=0:65535
```

### Flags

| Flag     | Description                | Example                   |
|----------|----------------------------|---------------------------|
| `-ip`    | IP address to scan         | `-ip=192.168.1.1`         |
| `-range` | Port range (start:end)     | `-range=20:1024`          |

### Example Output

```
Scanning started, target IP: localhost, port range: 0-65535
[0] port 22 open
[1] port 80 open
[2] port 443 open
Scan complete.
```

## 🧠 How it Works

- Accepts IP and port range as CLI arguments
- Creates a buffered channel and distributes the ports among **600 goroutines**
- Each goroutine tries to open a TCP connection and reports back if the port is open
- Results are collected and displayed after the scan is complete

## 🧰 Future Improvements

- Support for UDP scanning
- Web UI for real-time scanning status
- Output to JSON/CSV
- Scan multiple IPs concurrently

