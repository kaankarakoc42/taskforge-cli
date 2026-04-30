package cmd

import "github.com/spf13/cobra"

var debug bool

var rootCmd = &cobra.Command{
	Use:   "taskforge",
	Short: "TaskForge CLI executes dynamic local task executors",
	Long:  "TaskForge CLI is a lightweight runtime for executing registered task executors with JSON input parameters.",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Enable debug output for request/response details")
}
