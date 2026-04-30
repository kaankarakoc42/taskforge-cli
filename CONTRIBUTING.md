# 🤝 Contributing to TaskForge CLI

Thanks for your interest in contributing.

This project is designed to be **contributor-friendly**, especially for developers who want to explore:

- async task execution
- retry strategies
- event-driven systems
- CLI tooling in Go

---

## 🚀 Getting Started

```bash
git clone https://github.com/yourname/taskforge-cli
cd taskforge-cli

go mod tidy
make run
```

Test basic functionality:

```bash
go run . run examples/hello.sh
go run . simulate --fail-rate=0.5
go run . watch
```

---

## 🧠 Project Philosophy

TaskForge CLI follows a strict boundary:

### ✅ What belongs here

- Local task execution (sandbox)
- Retry simulation
- Event visualization
- CLI UX
- Offline analysis (DLQ, logs)
- Remote API interaction (thin layer)

### ❌ What does NOT belong here

- Kafka integration
- Scheduler logic
- Distributed orchestration
- Database access
- Worker coordination

👉 These belong to the **TaskForge backend**, not the CLI.

---

## 🧱 Architecture Overview

The CLI operates in two modes:

### 🟢 Local Mode (default)

- In-memory execution
- Fake event stream
- Retry simulation
- No infrastructure required

### 🔵 Remote Mode

- Connects to TaskForge API
- Streams real events via WebSocket
- Delegates execution to backend

---

## 📂 Where to Contribute

### 🟢 Beginner Friendly

- `internal/retry/` → add new strategies
- `internal/output/` → improve formatting
- `cmd/` → enhance CLI UX & help messages

### 🟡 Intermediate

- `internal/stream/` → event streaming improvements
- `internal/dlq/` → analysis & grouping logic

### 🔴 Advanced

- `internal/client/remote.go` → API integration
- WebSocket streaming implementation
- Plugin system for retry strategies

---

## 🧩 Good First Issues

Look for issues labeled:

- `good first issue`
- `help wanted`

---

## 🧪 Testing

```bash
go test ./...
```

---

## 📏 Contribution Rules

- Keep PRs small and focused  
- Follow existing project structure  
- Avoid adding heavy dependencies  
- Do not mix backend logic into CLI  

---

## 🧨 Common Mistakes

❌ Adding Kafka logic into CLI  
❌ Implementing distributed execution  
❌ Tight coupling with backend  

---

## 🧾 Pull Request Process

1. Fork the repo  
2. Create a branch  
3. Make changes  
4. Test locally  
5. Submit PR  

---

## 💡 Final Note

Small, focused contributions are preferred.
