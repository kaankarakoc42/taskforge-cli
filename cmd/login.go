package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"taskforge-cli/internal/config"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Save remote API credentials for TaskForge CLI",
	RunE: func(cmd *cobra.Command, args []string) error {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("API base URL [http://localhost:8080]: ")
		baseURLInput, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("read API base URL: %w", err)
		}
		baseURL := strings.TrimSpace(baseURLInput)
		if baseURL == "" {
			baseURL = "http://localhost:8080"
		}

		fmt.Print("Email: ")
		emailInput, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("read email: %w", err)
		}
		email := strings.TrimSpace(emailInput)
		if email == "" {
			return fmt.Errorf("email is required")
		}

		fmt.Print("Password: ")
		passwordInput, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("read password: %w", err)
		}
		password := strings.TrimSpace(passwordInput)
		if password == "" {
			return fmt.Errorf("password is required")
		}

		payload := map[string]string{
			"email":    email,
			"password": password,
		}
		body, err := json.Marshal(payload)
		if err != nil {
			return fmt.Errorf("encode login payload: %w", err)
		}

		loginURL := strings.TrimRight(baseURL, "/") + "/api/auth/login"
		req, err := http.NewRequestWithContext(cmd.Context(), http.MethodPost, loginURL, bytes.NewReader(body))
		if err != nil {
			return fmt.Errorf("build login request: %w", err)
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return fmt.Errorf("login request failed: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode >= 400 {
			return fmt.Errorf("login failed with status %d", resp.StatusCode)
		}

		var loginResp struct {
			Success bool `json:"success"`
			Data    struct {
				Token string `json:"token"`
			} `json:"data"`
			Token string `json:"token"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&loginResp); err != nil {
			return fmt.Errorf("decode login response: %w", err)
		}

		token := strings.TrimSpace(loginResp.Data.Token)
		if token == "" {
			// Backward compatibility for flat response payloads.
			token = strings.TrimSpace(loginResp.Token)
		}
		if token == "" {
			return fmt.Errorf("login response did not include token")
		}

		cfg := config.Config{
			APIBaseURL:   baseURL,
			WebSocketURL: "ws://localhost:8090/ws",
			Token:        token,
		}
		if err := config.SaveConfig(cfg); err != nil {
			return err
		}

		fmt.Println("Login saved.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
