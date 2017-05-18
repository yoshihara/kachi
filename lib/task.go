package lib

import (
	"github.com/mitchellh/go-homedir"
	"os"
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

// JSONPath return filepath for each JSON file
func JSONPath(filename string) string {
	dir, error := homedir.Dir()
	if error != nil {
		panic("[ERROR] Couldn't find home dir. abort")
	}

	kachiDir := filepath.Join(dir, ".kachi")

	_, error = os.Stat(kachiDir)

	if os.IsNotExist(error) {
		error = os.Mkdir(kachiDir, os.ModeDir|0755)
		if error != nil {
			panic("[ERROR] Couldn't create '.kachi' dir in home dir. abort")
		}
	}

	return filepath.Join(kachiDir, filename)
}
