package executor

import "context"

type Executor interface {
	Name() string
	Execute(ctx context.Context, params map[string]interface{}) (Result, error)
}

type Result struct {
	Output  interface{} `json:"output"`
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
}
