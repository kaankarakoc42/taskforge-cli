package runner

import (
	"context"
	"fmt"

	"github.com/kaankarakoc42/taskforge-sdk/pkg/executor"
	_ "taskforge-cli/internal/executors"
)

func Run(ctx context.Context, executorName string, params map[string]any) (any, error) {
	e, ok := executor.Get(executorName)
	if !ok {
		return nil, fmt.Errorf("executor not found: %s", executorName)
	}

	result, err := e.Execute(ctx, params)
	if err != nil {
		output := result.Output
		if output == nil {
			output = map[string]any{
				"success": false,
				"error":   result.Error,
			}
		}
		return output, fmt.Errorf("executor %q failed: %w", executorName, err)
	}

	return result.Output, nil
}
