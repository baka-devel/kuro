package task

import (
	"fmt"

	"github.com/baka-devel/kuro/internal/task"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add tasks to your task lists",
	Long:  "add tasks to your task lists",
	RunE: func(cmd *cobra.Command, args []string) error {
		group, err := cmd.Flags().GetString("group")
		if err != nil {
			return err
		}
		tasks, err := task.LoadTasks()
		if err != nil {
			return err
		}
		fmt.Println("Adding Tasks To Group: ", group)
		tasks.AddTask(group, args...)
		err = task.SaveTasks(tasks)
		if err != nil {
			return err
		}
		fmt.Println("Done!")
		return nil
	},
}

func init() {
	addCmd.Flags().StringP("group", "g", "main", "the group name in which the task will be added.")
	TaskCmd.AddCommand(addCmd)
}
