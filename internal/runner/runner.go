package runner

import (
	"context"
	"fmt"

	"taskforge-cli/internal/executor"
	_ "taskforge-cli/internal/executors"
)

func Run(ctx context.Context, executorName string, params map[string]any) (any, error) {
	e := executor.Get(executorName)
	if e == nil {
		return nil, fmt.Errorf("executor not found: %s", executorName)
	}

	result, err := e.Execute(ctx, params)
	if err != nil {
		return result, fmt.Errorf("executor %q failed: %w", executorName, err)
	}

	return result, nil
}
