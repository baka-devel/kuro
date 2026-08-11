package cmd

import (
	"fmt"
	"os"

	"github.com/baka-devel/kuro/cmd/task"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kuro",
	Short: "kuro is a simple task management tool written in go.",
	Long:  "kuro is a simple task management tool written in go, uses .json files to export your tasks, and works entirly from the terminal.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Excute() {
	rootCmd.AddCommand(task.TaskCmd)
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
