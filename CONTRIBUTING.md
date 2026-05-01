# Contributing to TaskForge

Thanks for helping improve TaskForge.

TaskForge is now organized as a CLI plus a separate public SDK module (`github.com/kaankarakoc42/taskforge-sdk/pkg/executor`).
Contributions that improve extensibility and developer experience are especially welcome.

## Contribution Paths

### 1) Add New Executors

- Add built-in executors under `internal/executors`.
- Implement the public `executor.Executor` interface.
- Register in `init()` with `executor.Register(...)`.
- Include sample params in `examples/` when useful.

### 2) Improve the SDK

- Work in the `taskforge-sdk` repository (`pkg/executor`).
- Keep APIs simple, stable, and Go-idiomatic.
- Favor backwards compatibility where possible.

### 3) Add Integrations

- Improve remote API and WebSocket flow in `internal/client`.
- Keep CLI behavior thin: transport and UX only, no orchestration logic.

## Local Setup

```bash
git clone <your-fork-url>
cd taskforge-cli
go mod tidy
go test ./...
```

## Coding Guidelines

- Keep PRs focused and small.
- Avoid unnecessary dependencies.
- Preserve compatibility for existing CLI commands.
- Prefer clear naming over abstraction-heavy designs.

## Adding an Executor Checklist

- [ ] Implements `Name() string`
- [ ] Implements `Execute(context.Context, map[string]interface{}) (executor.Result, error)`
- [ ] Registers itself in `init()`
- [ ] Has basic docs or example input
- [ ] Works via `taskforge run <name> --params <file>`

## Pull Requests

1. Fork and create a feature branch.
2. Add tests when behavior changes.
3. Run `go test ./...`.
4. Open a PR with:
   - problem statement
   - approach
   - how to test

Thanks for building TaskForge with us.
