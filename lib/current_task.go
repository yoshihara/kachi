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
	taskJSON, err := json.Marshal(&task)
	if err != nil {
		return nil, errors.New("[ERROR] Couldn't create 'current.json'. Please check your task name")
	}

	_, err = os.Stat("current.json")
	if !os.IsNotExist(err) {
		return nil, errors.New("[ERROR] 'current.json' exist already. Please stop before start")
	}

	f, error := os.Create("current.json")
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
	task.End = time.Now().Format(DateTimeLayout)
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
