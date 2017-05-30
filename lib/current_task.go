package lib

import (
	"bufio"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"time"
)

// StartTask create and start current Task
func StartTask(taskName string) (*Task, error) {
	var task = Task{Name: taskName, Start: time.Now().Format(DateTimeLayout)}
	taskJSON, error := json.Marshal(&task)
	if error != nil {
		return nil, errors.New("[ERROR] Couldn't create 'current.json'. Please check your task name")
	}

	_, error = os.Stat(JSONPath("current.json"))
	if !os.IsNotExist(error) {
		return nil, errors.New("[ERROR] 'current.json' exist already. Please stop before start")
	}

	f, error := os.Create(CurrentTaskPath())
	if error != nil {
		return nil, errors.New("[ERROR] Couldn't create 'current.json' file")
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	writer.WriteString(string(taskJSON))
	writer.Flush()

	return &task, nil
}

// ReadCurrentTask return the current Task
func ReadCurrentTask() (*Task, error) {
	bytes, error := ioutil.ReadFile(CurrentTaskPath())
	if error != nil {
		return nil, errors.New("[ERROR] Couldn't read current task. Please check if 'current.json' exists")
	}

	var task = Task{}
	error = json.Unmarshal(bytes, &task)

	if error != nil {
		return nil, errors.New("[ERROR] Couldn't parse current task JSON. Please check if 'current.json' contains valid JSON")
	}

	return &task, nil
}

// CompleteTask add the current Task to log and remove
func CompleteTask(task *Task) (*Task, error) {
	task.End = time.Now().Format(DateTimeLayout)
	taskJSON, error := json.Marshal(task)
	if error != nil {
		return nil, error
	}

	f, error := os.OpenFile(JSONPath("log.json"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if error != nil {
		return nil, error
	}

	defer f.Close()

	_, error = f.WriteString(string(taskJSON) + "\n")
	if error != nil {
		return nil, error
	}

	error = os.Remove(CurrentTaskPath())
	if error != nil {
		return nil, error
	}

	return task, nil
}
