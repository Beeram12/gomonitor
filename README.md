# Gomonitor
Gomonitor is a lightweight and concurrent website uptime monitoring service built with Go. It starts as a simple command-line application that uses Go’s concurrency to efficiently check if a list of websites is online. It then progressively evolves into a full-fledged monitoring service complete with a database for storing historical data, real-time Slack notifications for outages, a REST API, and a web dashboard for visualizing uptime.

## Features
- High-Performance Concurrent Checking: Uses Go’s goroutines to check hundreds of websites simultaneously without blocking.
- Interactive TUI Dashboard: A real-time Terminal User Interface built with Bubble Tea to monitor services live from your terminal.
- Robust CLI Commands: A powerful and scalable command structure built with Cobra for easy management (start, add, list, remove).
- Simple Configuration: Manages monitored URLs through a simple config.yaml file.
- Future-Proof Architecture: Designed to evolve with a clean separation of concerns, ready for database integration, a REST API, and a full web frontend.

## Architecture
The project follows the standard Go project layout to ensure a clean separation between the application code and reusable libraries. This makes the project scalable, maintainable, and easy to contribute to.

<img width="3533" height="1498" alt="go-monitor" src="https://github.com/user-attachments/assets/f75ae46f-347f-46b2-98ea-6516c346e764" />


## Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

###  Prerequisites
- Go (version 1.18 or higher)
- Make (for using the Makefile commands)

### Installation
1. Clone the repository to your local machine:
    ```
    git clone git@github.com:Beeram12/gomonitor.git
    ```
2. Navigate into the project directory:
    ```
    cd gomonitor
    ```
## Usage (with Makefile)
This project uses a Makefile to simplify common development tasks.

**Build the application binary**:
```
make build
```
**Build and run the application:**
This will start the main interactive TUI dashboard.
```
make run
```
**Clean up binary from the last build:**
This removes the gomonitor binary and the config.yaml file.
```
make clean
```
## Directory Structure
The project’s directory structure is organized for modularity and scalability:
```
gomonitor/
├── cmd/
│   └── gomonitor/
│       └── main.go        # Main entry point for the CLI application.
├── pkg/
│   ├── checker/           # Public, reusable package for the core website checking logic.
│   └── config/            # Public, reusable package for loading/saving the config file.
├── internal/
│   ├── commands/          # All private Cobra command definitions (start, add, etc.).
│   └── tui/               # All private Bubble Tea logic for the TUI (model, view, update).
├── go.mod                 # Manages the project’s dependencies.
├── Makefile               # Automates common development tasks.
└── README.md              # You are here!
```
## License
This project is licensed under the MIT License - see the LICENSE.md file for details.
