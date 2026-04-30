# TaskForge CLI - Full Skeleton

## Repo Structure

```
taskforge-cli/
├── cmd/
│   ├── root.go
│   ├── run.go
│   ├── simulate.go
│   ├── watch.go
│   └── analyze.go
├── internal/
│   ├── client/
│   ├── runner/
│   ├── retry/
│   ├── stream/
│   ├── dlq/
│   └── output/
├── pkg/
│   └── sdk/
├── examples/
├── docs/
├── .github/
├── go.mod
├── Makefile
├── README.md
├── CONTRIBUTING.md
└── LICENSE
```

## Core Idea

TaskForge CLI is:

- Sandbox (local mode)
- Gateway (remote mode)

## Local Mode

- No Kafka
- No DB
- No backend required

## Remote Mode

- Connects to TaskForge API
- Uses REST + WebSocket

## Quick Start

```bash
go mod init github.com/yourname/taskforge-cli
go get github.com/spf13/cobra@v1.8.1
go mod tidy
make run
```

## Philosophy

- CLI must stay lightweight
- No distributed logic inside CLI
- Backend owns orchestration

## Good First Issues

- Add retry strategy
- Improve CLI output
- Implement remote client
- Enhance DLQ analyzer

