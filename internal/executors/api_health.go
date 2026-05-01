package executors

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/kaankarakoc42/taskforge-sdk/pkg/executor"
)

type APIHealthExecutor struct{}

func (e *APIHealthExecutor) Name() string {
	return "api_health"
}

func (e *APIHealthExecutor) Execute(ctx context.Context, params map[string]interface{}) (executor.Result, error) {
	urlValue, ok := params["url"].(string)
	if !ok || urlValue == "" {
		err := fmt.Errorf("missing required param: url (string)")
		return executor.Result{
			Success: false,
			Error:   err.Error(),
		}, err
	}

	expectedStatus := intValueOrDefault(params["expected_status"], 200)
	timeoutSeconds := intValueOrDefault(params["timeout_seconds"], 5)

	reqCtx, cancel := context.WithTimeout(ctx, time.Duration(timeoutSeconds)*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(reqCtx, http.MethodGet, urlValue, nil)
	if err != nil {
		wrappedErr := fmt.Errorf("build request: %w", err)
		return executor.Result{
			Success: false,
			Error:   wrappedErr.Error(),
		}, wrappedErr
	}

	start := time.Now()
	resp, err := http.DefaultClient.Do(req)
	latencyMS := time.Since(start).Milliseconds()

	if err != nil {
		result := map[string]any{
			"url":        urlValue,
			"healthy":    false,
			"latency_ms": latencyMS,
			"error":      err.Error(),
		}
		wrappedErr := fmt.Errorf("request failed: %w", err)
		return executor.Result{
			Output:  result,
			Success: false,
			Error:   wrappedErr.Error(),
		}, wrappedErr
	}
	defer resp.Body.Close()

	healthy := resp.StatusCode == expectedStatus
	result := map[string]any{
		"url":         urlValue,
		"status_code": resp.StatusCode,
		"healthy":     healthy,
		"latency_ms":  latencyMS,
	}

	return executor.Result{
		Output:  result,
		Success: true,
	}, nil
}

func intValueOrDefault(raw any, fallback int) int {
	switch v := raw.(type) {
	case nil:
		return fallback
	case float64:
		return int(v)
	case int:
		return v
	default:
		return fallback
	}
}

func init() {
	executor.Register(&APIHealthExecutor{})
}
