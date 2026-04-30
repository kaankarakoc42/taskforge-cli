package client

import (
	"context"
	"fmt"
)

type RemoteClient struct {
	baseURL string
}

func NewRemoteClient(baseURL string) *RemoteClient {
	return &RemoteClient{baseURL: baseURL}
}

func (c *RemoteClient) RunTask(ctx context.Context, executorName string, params map[string]any) (any, error) {
	return nil, fmt.Errorf("remote mode is TODO: this client should only proxy run requests to backend API")
}

func (c *RemoteClient) Watch(ctx context.Context, taskID string) (<-chan Event, error) {
	return nil, fmt.Errorf("remote mode is TODO: this client should only open websocket stream from backend")
}
