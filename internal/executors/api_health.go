package executors

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"taskforge-cli/internal/executor"
)

type APIHealthExecutor struct{}

func (e *APIHealthExecutor) Execute(ctx context.Context, params map[string]any) (any, error) {
	urlValue, ok := params["url"].(string)
	if !ok || urlValue == "" {
		return nil, fmt.Errorf("missing required param: url (string)")
	}

	expectedStatus := intValueOrDefault(params["expected_status"], 200)
	timeoutSeconds := intValueOrDefault(params["timeout_seconds"], 5)

	reqCtx, cancel := context.WithTimeout(ctx, time.Duration(timeoutSeconds)*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(reqCtx, http.MethodGet, urlValue, nil)
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
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
		return result, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	healthy := resp.StatusCode == expectedStatus
	result := map[string]any{
		"url":         urlValue,
		"status_code": resp.StatusCode,
		"healthy":     healthy,
		"latency_ms":  latencyMS,
	}

	return result, nil
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
	executor.Register("api_health", &APIHealthExecutor{})
}
