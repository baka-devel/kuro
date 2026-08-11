package task

import "github.com/spf13/cobra"

// TaskCmd represents the base task command that groups all task management subcommands.
var TaskCmd = &cobra.Command{
	Use:   "task",
	Short: "Manage Your Tasks",
	Long:  "Manage Your Tasks",
}
