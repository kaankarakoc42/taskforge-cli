package client

import "context"

type Event struct {
	Type string         `json:"type"`
	Data map[string]any `json:"data"`
}

type Client interface {
	RunTask(ctx context.Context, executorName string, params map[string]any) (any, error)
	Watch(ctx context.Context, taskID string) (<-chan Event, error)
}
