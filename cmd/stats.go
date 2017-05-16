package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yoshihara/kachi/lib"
	"io/ioutil"
	"math"
	"os"
	"strings"
	"time"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Show tracked tasks",
	Long:  "stats: Show tracked tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: 関数内部はlib/配下に作ってテストを書きたい
		layout := "Mon Jan 2 15:04:05 MST 2006"

		bytes, error := ioutil.ReadFile("log.json")
		if error != nil {
			return errors.New("[ERROR] Couldn't read 'log.json'. Please check if 'log.json' exists")
		}

		tasks := string(bytes)
		task := lib.Task{}

		for _, taskJSON := range strings.Split(tasks, "\n") {
			if len(taskJSON) < 1 {
				continue
			}
			error = json.Unmarshal([]byte(taskJSON), &task)
			if error != nil {
				return errors.New("[ERROR] Couldn't parse task log as JSON:" + taskJSON)
			}
			startDateTime, _ := time.Parse(layout, task.Start)
			endDateTime, _ := time.Parse(layout, task.End)

			// TODO: もうちょっとフォーマット何とかする
			duration := float64(endDateTime.Sub(startDateTime).Seconds()) * scale
			minutes := duration / 60.0
			hours := minutes / 60.0

			// 小数点丸め
			fmt.Fprintf(os.Stdout, "%s %v\n", task.Name, math.Trunc(hours*100)/100.0)
		}
		return nil
	},
}
