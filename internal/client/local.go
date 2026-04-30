package client

import (
	"context"

	"taskforge-cli/internal/runner"
)

type LocalClient struct{}

func NewLocalClient() *LocalClient {
	return &LocalClient{}
}

func (c *LocalClient) RunTask(ctx context.Context, executorName string, params map[string]any) (any, error) {
	return runner.Run(ctx, executorName, params)
}

func (c *LocalClient) Watch(ctx context.Context, taskID string) (<-chan Event, error) {
	events := make(chan Event)
	close(events)
	return events, nil
}
