package task

// Tasks represents a mapping of group names to their respective list of tasks.
type Tasks map[string]Group

// Group represents a slice of Task instances to be put under the group name in Tasks instance.
type Group []Task

// Task represents a single task, with its name and state(whether it is done or not).
type Task struct {
	TaskName string `json:"taskName"`
	Done     bool   `json:"done"`
}
