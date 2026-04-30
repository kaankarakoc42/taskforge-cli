package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"taskforge-cli/internal/client"
)

var paramsFile string
var useRemote bool
var remoteBaseURL string

var runCmd = &cobra.Command{
	Use:   "run <executor_name>",
	Short: "Run a registered executor with params JSON",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		executorName := args[0]

		raw, err := os.ReadFile(paramsFile)
		if err != nil {
			return fmt.Errorf("read params file: %w", err)
		}

		params := map[string]any{}
		if err := json.Unmarshal(raw, &params); err != nil {
			return fmt.Errorf("parse params JSON: %w", err)
		}

		var c client.Client = client.NewLocalClient()
		if useRemote {
			c = client.NewRemoteClient(remoteBaseURL)
		}

		result, err := c.RunTask(context.Background(), executorName, params)
		if err != nil {
			return err
		}

		out, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return fmt.Errorf("encode result JSON: %w", err)
		}

		fmt.Println(string(out))
		return nil
	},
}

func init() {
	runCmd.Flags().StringVarP(&paramsFile, "params", "p", "", "Path to JSON params file")
	runCmd.Flags().BoolVar(&useRemote, "remote", false, "Use remote gateway mode (TODO)")
	runCmd.Flags().StringVar(&remoteBaseURL, "api-base-url", "http://localhost:8080", "Backend API base URL for remote mode")
	_ = runCmd.MarkFlagRequired("params")
	rootCmd.AddCommand(runCmd)
}
