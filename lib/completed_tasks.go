package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"time"
)

// ReadCompletedTasks return completed tasks
func ReadCompletedTasks() ([]byte, error) {
	bytes, error := ioutil.ReadFile(JSONPath("log.json"))
	if error != nil {
		return nil, errors.New("[ERROR] Couldn't read 'log.json'. Please check if 'log.json' exists")
	}

	return bytes, nil
}

// RefreshTasks move completed tasks to archive.json
func RefreshTasks() error {
	bytes, error := ReadCompletedTasks()
	if error != nil {
		return error
	}

	f, error := os.OpenFile(JSONPath("archive.json"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if error != nil {
		return error
	}

	defer f.Close()

	_, error = f.WriteString(string(bytes))
	if error != nil {
		return error
	}

	lf, error := os.OpenFile(JSONPath("log.json"), os.O_TRUNC, 0666)
	if error != nil {
		return error
	}

	defer lf.Close()

	return nil
}

// ShowTaskStat show task stat
func ShowTaskStat(taskJSON []byte, scale float64) error {
	task := Task{}

	error := json.Unmarshal([]byte(taskJSON), &task)
	if error != nil {
		return errors.New("[ERROR] Couldn't parse task log as JSON:" + string(taskJSON))
	}

	startDateTime, _ := time.Parse(DateTimeLayout, task.Start)
	endDateTime, _ := time.Parse(DateTimeLayout, task.End)

	// TODO: もうちょっとフォーマット何とかする
	duration := float64(endDateTime.Sub(startDateTime).Seconds()) * scale
	minutes := duration / 60.0
	hours := minutes / 60.0

	// 小数点丸め
	fmt.Fprintf(os.Stdout, "%s %v\n", task.Name, math.Trunc(hours*100)/100.0)
	return nil
}
