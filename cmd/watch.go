package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"taskforge-cli/internal/client"
)

var watchRemote bool
var watchTaskID string
var watchRemoteBaseURL string

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch remote backend websocket events",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !watchRemote {
			return fmt.Errorf("watch currently supports remote mode only. Use: taskforge watch --remote")
		}

		c := client.NewRemoteClient(watchRemoteBaseURL, debug)
		events, err := c.Watch(context.Background(), watchTaskID)
		if err != nil {
			return err
		}

		for evt := range events {
			out, err := json.MarshalIndent(evt, "", "  ")
			if err != nil {
				return fmt.Errorf("encode event: %w", err)
			}
			fmt.Println(string(out))
		}

		return nil
	},
}

func init() {
	watchCmd.Flags().BoolVar(&watchRemote, "remote", false, "Watch backend websocket in remote mode")
	watchCmd.Flags().StringVar(&watchTaskID, "task-id", "", "Optional task ID filter")
	watchCmd.Flags().StringVar(&watchRemoteBaseURL, "api-base-url", "", "Optional backend API base URL override for remote mode")
	rootCmd.AddCommand(watchCmd)
}
