package executor

import "context"

type Executor interface {
	Execute(ctx context.Context, params map[string]any) (any, error)
}
