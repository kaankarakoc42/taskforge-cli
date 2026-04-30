# 🚀 TaskForge CLI

> Run, simulate, and debug distributed task executors — locally or via a backend.

---

## ⚡ What is this?

TaskForge CLI is a **developer tool** for experimenting with:

- async task execution  
- retry / failure behavior  
- event-driven patterns  
- executor-based systems  

Without needing:

- Kafka  
- PostgreSQL  
- Kubernetes  

---

## 🧠 Core Idea

TaskForge CLI is built around a **dynamic executor system**.

```bash
taskforge run <executor_name> --params <file.json>
```

Executors:

- are registered dynamically  
- receive JSON input  
- return structured results  

---

## 🧱 Architecture Boundaries

This CLI has **strict responsibilities**:

### 🟢 Local Mode (default)

- Runs executors directly  
- Acts as a sandbox  
- No infrastructure required  

### 🔵 Remote Mode (`--remote`)

- Sends tasks to backend API  
- Streams events  
- **Must stay thin (no logic)**  

---

### ❌ Non-Goals

The CLI does NOT:

- implement orchestration  
- implement scheduling  
- manage workers  
- connect to Kafka  
- access databases  

---

## 📦 Project Structure

```
taskforge-cli/
├── main.go
├── go.mod
├── cmd/
│   ├── root.go
│   └── run.go
├── internal/
│   ├── executor/
│   │   ├── executor.go
│   │   └── registry.go
│   ├── executors/
│   │   └── api_health.go
│   └── runner/
│       └── runner.go
├── examples/
│   └── api_health.json
└── README.md
```

---

## ⚙️ Dynamic Executor System

### 1. Executor Interface

```go
type Executor interface {
    Execute(ctx context.Context, params map[string]any) (any, error)
}
```

---

### 2. Registration

```go
func init() {
    executor.Register("api_health", &APIHealthExecutor{})
}
```

---

### 3. Execution

```bash
taskforge run <executor_name> --params <file.json>
```

---

## 🎯 Example

```bash
taskforge run api_health --params examples/api_health.json
```

### Params

```json
{
  "url": "https://google.com",
  "expected_status": 200
}
```

### Output

```json
{
  "url": "https://google.com",
  "status_code": 200,
  "healthy": true,
  "latency_ms": 120
}
```

---

## 🔌 Built-in Executor

### `api_health`

Checks if an API endpoint is healthy.

**Params:**

- `url` (required)
- `expected_status` (default: 200)
- `timeout_seconds` (default: 5)

---

## 🚀 Quick Start

```bash
go mod tidy

go run . run api_health --params examples/api_health.json
```

---

## 🔌 Remote Mode (Gateway)

```bash
go run . run api_health \
  --params examples/api_health.json \
  --remote \
  --api-base-url http://localhost:8080
```

### Status

- 🚧 Not fully implemented yet  
- Designed as a thin API/WebSocket client  

### Design Rule

Remote mode:

- does NOT execute tasks  
- does NOT implement logic  
- only forwards and observes  

---

## 🧩 Why this is useful

- test task logic locally  
- build executors without backend  
- debug failure scenarios  
- experiment with distributed system patterns  

---

## 🚧 Roadmap

- [ ] Retry simulation (fail-rate, backoff)
- [ ] Event streaming (`watch`)
- [ ] Remote API integration
- [ ] Plugin system for executors

---

## 🤝 Contributing

Good areas:

- new executors  
- retry strategies  
- CLI UX improvements  
- output formatting  

See `CONTRIBUTING.md`

---

## 🔥 TL;DR

TaskForge CLI is a **playground for task execution logic** — without the overhead of distributed systems infrastructure.
