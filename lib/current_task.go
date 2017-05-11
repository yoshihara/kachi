package lib

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"time"
)

// ReadCurrentTask return the current Task
func ReadCurrentTask() (*Task, error) {
	bytes, err := ioutil.ReadFile("current.json")
	if err != nil {
		return nil, errors.New("[ERROR] Couldn't read current task. Please check if 'current.json' exists")
	}

	var task = Task{}
	err = json.Unmarshal(bytes, &task)

	if err != nil {
		return nil, errors.New("[ERROR] Couldn't parse current task JSON. Please check if 'current.json' contains valid JSON")
	}

	return &task, nil
}

// CompleteTask add the current Task to log and remove
func CompleteTask(task *Task) (*Task, error) {
	layout := "Mon Jan 2 15:04:05 MST 2006"

	task.End = time.Now().Format(layout)
	taskJSON, err := json.Marshal(task)
	if err != nil {
		return nil, err
	}

	f, err := os.OpenFile("log.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	_, err = f.WriteString(string(taskJSON) + "\n")
	if err != nil {
		return nil, err
	}

	os.Remove("current.json")
	if err != nil {
		return nil, err
	}

	return task, nil
}
