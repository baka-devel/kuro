package main

import (
	"fmt"

	"github.com/baka-devel/kuro/cmd"
	"github.com/baka-devel/kuro/internal/task"
)

func main() {
	cmd.Excute()
	t, err := task.LoadTasks()
	if err != nil {
		fmt.Println(err)
		return
	}
	// t.AddTask("main", "task3")
	// err = tasks.SaveTasks(t)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	fmt.Println(t)
}
