# Project Gomonitor

The project starts as a simple command-line application that uses Go's concurrency to efficiently check if a list of websites is online. It then progressively evolves into a full-fledged monitoring service complete with a database for storing historical data, real-time Slack notifications for outages, a REST API, and a web dashboard for visualizing uptime.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

Run build make command with tests
```bash
make all
```

Build the application
```bash
make build
```

Run the application
```bash
make run
```

Live reload the application:
```bash
make watch
```

Run the test suite:
```bash
make test
```

Clean up binary from the last build:
```bash
make clean
```
