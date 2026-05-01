# TaskForge CLI

TaskForge is a lightweight task execution runtime with a public Go SDK for building and sharing executors.

It is designed for two use cases:
- run executors locally for fast development and testing
- forward tasks to a remote backend from the same CLI

## Why TaskForge

Task systems are often hard to iterate on because they depend on infrastructure.
TaskForge lets you build and validate executor logic as plain Go code first, then run it through the same CLI in local or remote mode.
The SDK now lives in a separate module: `github.com/kaankarakoc42/taskforge-sdk`.

## Project Structure

```text
taskforge-cli/
├── cmd/                     # cobra commands
├── internal/                # CLI internals and built-ins
│   ├── client/
│   ├── executors/
│   └── runner/
├── examples/
│   └── api_health.json
├── CONTRIBUTING.md
└── README.md
```

## SDK Dependency

TaskForge CLI consumes the external SDK module:

```go
import "github.com/kaankarakoc42/taskforge-sdk/pkg/executor"
```

Interface:

```go
type Executor interface {
    Name() string
    Execute(ctx context.Context, params map[string]interface{}) (Result, error)
}

type Result struct {
    Output  interface{}
    Success bool
    Error   string
}
```

Registry:

```go
func Register(e Executor)
func Get(name string) (Executor, bool)
func List() []Executor
```

## Writing an Executor

Example:

```go
package myexec

import (
    "context"
    "fmt"

    "github.com/kaankarakoc42/taskforge-sdk/pkg/executor"
)

type HelloExecutor struct{}

func (h *HelloExecutor) Name() string { return "hello" }

func (h *HelloExecutor) Execute(ctx context.Context, params map[string]interface{}) (executor.Result, error) {
    name, _ := params["name"].(string)
    if name == "" {
        name = "world"
    }

    return executor.Result{
        Output: map[string]interface{}{"message": fmt.Sprintf("hello, %s", name)},
        Success: true,
    }, nil
}

func init() {
    executor.Register(&HelloExecutor{})
}
```

See runnable SDK sample in the SDK repository:
`taskforge-sdk/examples/custom-executor/main.go`.

## CLI Usage

Run built-in `api_health` locally:

```bash
go run . run api_health --params examples/api_health.json
```

Run via remote backend:

```bash
go run . run api_health --params examples/api_health.json --remote --api-base-url https://taskforge.kaankarakoc.com
```

Watch remote task events:

```bash
go run . watch --remote --api-base-url https://taskforge.kaankarakoc.com
```

## Built-in Executor

`api_health` checks endpoint health.

Params:
- `url` (required)
- `expected_status` (optional, default `200`)
- `timeout_seconds` (optional, default `5`)

## External Contributor Workflow

1. Implement an executor using `github.com/kaankarakoc42/taskforge-sdk/pkg/executor`.
2. Register it with `executor.Register(...)`.
3. Run locally in your own Go program, or include it in a TaskForge-based binary.

## Development

```bash
go test ./...
```

For contribution guidance, see `CONTRIBUTING.md`.
