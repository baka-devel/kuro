package task

// AddTask adds all the given tasks to the specified group (if given, defaults to "main") to the Tasks map.
func (t Tasks) AddTask(groupName string, taskNames ...string) {
	for _, taskName := range taskNames {
		t[groupName] = append(t[groupName], Task{TaskName: taskName})
	}
}
