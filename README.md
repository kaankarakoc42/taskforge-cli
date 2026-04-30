# 🚀 TaskForge CLI

> CLI tool to simulate and debug distributed task execution — locally or with a real backend.

---

## ⚡ What is this?

TaskForge CLI is a **developer tool** for experimenting with:

- async task execution  
- retry strategies (backoff, failures)  
- event streams  
- DLQ analysis  

Without needing:

- Kafka  
- PostgreSQL  
- Kubernetes  

---

## 🧠 Core Idea

TaskForge CLI has **two modes**:

### 🟢 Local Mode (default)

- Runs tasks locally  
- Simulates retries & failures  
- Streams fake events  
- No setup required  

### 🔵 Remote Mode

- Connects to TaskForge backend  
- Submits real tasks  
- Streams real-time events via WebSocket  

---

## 📦 Repo Structure

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

---

## ⚡ Quick Start

```bash
go mod init github.com/yourname/taskforge-cli
go get github.com/spf13/cobra@v1.8.1
go mod tidy

go run . run examples/hello.sh
go run . simulate --fail-rate=0.5
go run . watch
```

---

## 🎯 Example Usage

```bash
taskforge run examples/hello.sh
taskforge simulate --fail-rate=0.7
taskforge watch
```

---

## 🔌 Remote Usage (coming soon)

```bash
taskforge run examples/hello.sh --remote
taskforge watch --remote
```

---

## 🧱 Philosophy

- CLI = sandbox + gateway  
- No distributed orchestration inside CLI  
- Backend owns scheduling, queues, and execution  
- Local mode must always work without infrastructure  

---

## 🚧 Roadmap

- [ ] Remote API integration  
- [ ] WebSocket event streaming  
- [ ] Advanced retry strategies (jitter, policies)  
- [ ] Plugin system for extensions  

---

## 🧩 Good First Issues

- Add new retry strategy  
- Improve CLI output formatting  
- Implement remote client  
- Enhance DLQ analyzer  

---

## 🔥 Why this exists

Distributed systems are hard to test.

TaskForge CLI lets you:

- experiment locally  
- debug retry behavior  
- visualize task execution  

Before deploying anything.

---

## 📌 Scope

This CLI does NOT:

- manage Kafka  
- handle scheduling  
- persist tasks  
- run distributed workers  

Those belong to the TaskForge backend.

---

## 🤝 Contributing

See CONTRIBUTING.md for guidelines.
