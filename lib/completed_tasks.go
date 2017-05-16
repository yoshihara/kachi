package lib

import (
	"errors"
	"io/ioutil"
	"os"
)

// RefreshTasks move completed tasks to archive.json
func RefreshTasks() error {
	bytes, error := ioutil.ReadFile("log.json")
	if error != nil {
		return errors.New("[ERROR] Couldn't read 'log.json'. Please check if 'log.json' exists")
	}

	f, error := os.OpenFile("archive.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if error != nil {
		return error
	}

	defer f.Close()

	_, error = f.WriteString(string(bytes))
	if error != nil {
		return error
	}

	lf, error := os.OpenFile("log.json", os.O_TRUNC, 0666)
	if error != nil {
		return error
	}

	defer lf.Close()

	return nil
}
