package task

import (
	"encoding/json"
	"os"
)

// LoadTasks reads and parses the tasks JSON file from the default storage path.
func LoadTasks() (Tasks, error) {
	fpath := "./taskes.json"

	data, err := os.ReadFile(fpath)
	if err != nil {
		return nil, err
	}

	var tasks Tasks
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// SaveTasks saves the given tasks into JSON file in the default storage path.
func SaveTasks(t Tasks) error {
	fpath := "./taskes.json"
	data, err := json.MarshalIndent(t, "", "\t")
	if err != nil {
		return err
	}
	err = os.WriteFile(fpath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
