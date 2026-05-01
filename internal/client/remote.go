package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"

	"taskforge-cli/internal/config"
)

type RemoteClient struct {
	baseURLOverride string
	debug           bool
}

func NewRemoteClient(baseURL string, debug bool) *RemoteClient {
	return &RemoteClient{
		baseURLOverride: baseURL,
		debug:           debug,
	}
}

func (c *RemoteClient) RunTask(ctx context.Context, executorName string, params map[string]any) (any, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	if strings.TrimSpace(cfg.Token) == "" {
		return nil, fmt.Errorf("Not logged in. Run: taskforge login")
	}

	baseURL := strings.TrimSpace(c.baseURLOverride)
	if baseURL == "" {
		baseURL = strings.TrimSpace(cfg.APIBaseURL)
	}
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	payload := map[string]any{
		// Newer backend contract requires "type"; keep "handler" for compatibility.
		"type":    executorName,
		"handler": executorName,
		"params":  params,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("encode request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, strings.TrimRight(baseURL, "/")+"/api/tasks", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+cfg.Token)

	if c.debug {
		printRequestDebug(req.Method, req.URL.String(), cfg.Token, body)
		printCurlDebug(req.Method, req.URL.String(), cfg.Token, body)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("remote request failed: %w", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body: %w", err)
	}

	if c.debug {
		printResponseDebug(resp.Status, responseBody)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("remote request failed (%d)\nResponse: %s", resp.StatusCode, strings.TrimSpace(string(responseBody)))
	}

	if len(bytes.TrimSpace(responseBody)) == 0 {
		return map[string]any{}, nil
	}

	var result any
	if err := json.Unmarshal(responseBody, &result); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}

func (c *RemoteClient) Watch(ctx context.Context, taskID string) (<-chan Event, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}
	if strings.TrimSpace(cfg.Token) == "" {
		return nil, fmt.Errorf("Not logged in. Run: taskforge login")
	}

	wsURL := strings.TrimSpace(cfg.WebSocketURL)
	if strings.TrimSpace(c.baseURLOverride) != "" {
		wsURL = config.DeriveWebSocketURL(c.baseURLOverride)
	}
	if wsURL == "" {
		wsURL = config.DeriveWebSocketURL(cfg.APIBaseURL)
	}
	if wsURL == "" {
		wsURL = config.DeriveWebSocketURL("")
	}

	header := http.Header{}
	header.Set("Authorization", "Bearer "+cfg.Token)
	conn, _, err := websocket.DefaultDialer.DialContext(ctx, wsURL, header)
	if err != nil {
		return nil, fmt.Errorf("websocket connect failed: %w", err)
	}

	ch := make(chan Event)
	go func() {
		defer close(ch)
		defer conn.Close()

		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			var envelope map[string]any
			if err := conn.ReadJSON(&envelope); err != nil {
				return
			}

			eventType, _ := envelope["type"].(string)
			data, _ := envelope["data"].(map[string]any)
			if data == nil {
				data = envelope
			}

			event := Event{
				Type: eventType,
				Data: data,
			}

			if taskID != "" {
				if id, ok := data["task_id"].(string); ok && id != taskID {
					continue
				}
			}

			select {
			case <-ctx.Done():
				return
			case ch <- event:
			}
		}
	}()

	return ch, nil
}

func printRequestDebug(method, url, token string, body []byte) {
	fmt.Println("---- REQUEST ----")
	fmt.Printf("%s %s\n", method, url)
	fmt.Println("Headers:")
	fmt.Printf("Authorization: Bearer %s\n", maskedToken(token))
	fmt.Println("Content-Type: application/json")
	fmt.Println()
	fmt.Println("Body:")
	fmt.Println(prettyJSON(body))
	fmt.Println("------------")
}

func printCurlDebug(method, url, token string, body []byte) {
	fmt.Println("---- CURL ----")
	fmt.Printf("curl -X %s %s \\\n", method, url)
	fmt.Printf("-H \"Authorization: Bearer %s\" \\\n", maskedToken(token))
	fmt.Println("-H \"Content-Type: application/json\" \\")
	fmt.Printf("-d '%s'\n", escapeSingleQuotes(string(body)))
	fmt.Println("-----------")
}

func printResponseDebug(status string, body []byte) {
	fmt.Println("---- RESPONSE ----")
	fmt.Printf("Status: %s\n", status)
	fmt.Println("Body:")
	trimmed := strings.TrimSpace(string(body))
	if trimmed == "" {
		fmt.Println("{}")
	} else {
		fmt.Println(prettyJSON(body))
	}
	fmt.Println("--------------------")
}

func prettyJSON(raw []byte) string {
	var out bytes.Buffer
	if err := json.Indent(&out, raw, "", "  "); err != nil {
		return strings.TrimSpace(string(raw))
	}
	return out.String()
}

func maskedToken(token string) string {
	if strings.TrimSpace(token) == "" {
		return ""
	}
	return "*****"
}

func escapeSingleQuotes(input string) string {
	return strings.ReplaceAll(input, "'", "'\\''")
}
