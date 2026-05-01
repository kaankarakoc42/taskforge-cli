package main

import (
	"context"
	"fmt"

	"taskforge-cli/pkg/executor"
)

type HelloExecutor struct{}

func (h *HelloExecutor) Name() string {
	return "hello"
}

func (h *HelloExecutor) Execute(_ context.Context, params map[string]interface{}) (executor.Result, error) {
	name, _ := params["name"].(string)
	if name == "" {
		name = "TaskForge"
	}

	return executor.Result{
		Output: map[string]interface{}{
			"message": fmt.Sprintf("hello, %s", name),
		},
		Success: true,
	}, nil
}

func main() {
	executor.Register(&HelloExecutor{})

	e, ok := executor.Get("hello")
	if !ok {
		panic("hello executor is not registered")
	}

	result, err := e.Execute(context.Background(), map[string]interface{}{
		"name": "SDK developer",
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("success=%t output=%v\n", result.Success, result.Output)
}
