package lib

import (
	"github.com/mitchellh/go-homedir"
	"path/filepath"
)

// Task JSONにするときのタスクの持つ情報
type Task struct {
	Name  string
	Start string
	End   string
}

// DateTimeLayout using datetime format
const DateTimeLayout string = "Mon Jan 2 15:04:05 MST 2006"

// JSONDir return directory path for each JSON file
func JSONDir() string {
	dir, error := homedir.Dir()
	if error != nil {
		panic("[ERROR] Couldn't find home dir. abort")
	}

	return filepath.Join(dir, ".kachi")
}

// JSONPath return filepath for each JSON file
func JSONPath(filename string) string {
	return filepath.Join(JSONDir(), filename)
}
