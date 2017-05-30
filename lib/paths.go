package lib

import (
	"errors"
	"github.com/mitchellh/go-homedir"
	"os"
	"path/filepath"
)

// ErrorNotFoundJSONDir Error for JSONDir not found
func ErrorNotFoundJSONDir() error {
	return errors.New("[ERROR] couldn't fild ~/.kachi directory. Please run 'init' command")
}

// IsJSONDirExists return if kachi directory exists
func IsJSONDirExists() bool {
	_, error := os.Stat(JSONDir())
	return !os.IsNotExist(error)
}

// JSONDir return directory path for each JSON file
func JSONDir() string {
	dir, error := homedir.Dir()
	if error != nil {
		panic("[ERROR] Couldn't find home dir. abort")
	}

	return filepath.Join(dir, ".kachi")
}

// CurrentTaskPath return path for the current task file
func CurrentTaskPath() string {
	return JSONPath("current.json")
}

// LogPath return path for the tasks log file
func LogPath() string {
	return JSONPath("log.json")
}

// JSONPath return filepath for each JSON file
func JSONPath(filename string) string {
	return filepath.Join(JSONDir(), filename)
}
