package runner

import (
	"context"
	"fmt"

	"github.com/kaankarakoc42/taskforge-sdk/pkg/executor"
	"github.com/kaankarakoc42/taskforge-sdk/pkg/logger"
	_ "taskforge-cli/internal/executors"
)

func Run(ctx context.Context, executorName string, params map[string]any) (any, error) {
	result, err := executor.Invoke(ctx, executorName, params, &executor.InvokeOptions{
		Logger: logger.NewConsoleLogger(),
		Metadata: &executor.TaskMetadata{
			TaskType: executorName,
		},
	})
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
