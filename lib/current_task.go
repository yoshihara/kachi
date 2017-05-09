package lib

import (
	"encoding/json"
	"errors"
	"io/ioutil"
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
