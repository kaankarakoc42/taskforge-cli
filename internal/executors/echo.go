package executors

import (
	"context"
	"fmt"
	"strings"

	"github.com/kaankarakoc42/taskforge-sdk/pkg/executor"
)

type EchoExecutor struct{}

func (e *EchoExecutor) Name() string {
	return "echo"
}

func (e *EchoExecutor) Execute(goCtx context.Context, params map[string]interface{}) (executor.Result, error) {
	log := executor.LoggerFromContext(goCtx)
	log.Debug("running echo executor")

	meta, ok := executor.TaskMetadataFromContext(goCtx)
	if ok && meta.TaskType != "" {
		log.Debug("task type: " + meta.TaskType)
	}

	message, _ := params["message"].(string)
	if strings.TrimSpace(message) == "" {
		err := fmt.Errorf("missing required param: message (string)")
		log.Error("echo failed: missing message param")
		return executor.Result{
			Success: false,
			Error:   err.Error(),
		}, err
	}

	uppercase, _ := params["uppercase"].(bool)
	if uppercase {
		message = strings.ToUpper(message)
	}

	log.Info("echo completed successfully")

	return executor.Result{
		Output: map[string]interface{}{
			"message": message,
			"length":  len(message),
		},
		Success: true,
	}, nil
}

func init() {
	executor.Register(&EchoExecutor{})
}
