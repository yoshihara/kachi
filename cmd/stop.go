package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yoshihara/kachi/lib"
	"os"
	"time"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop task tracking",
	Long:  "stop: stop task tracking",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: 関数内部はlib/配下に作ってテストを書きたい
		layout := "Mon Jan 2 15:04:05 MST 2006"

		task, err := lib.ReadCurrentTask()
		if err != nil {
			return err
		}

		task.End = time.Now().Format(layout)
		taskJSON, err := json.Marshal(task)
		if err != nil {
			panic(err)
		}

		f, err := os.OpenFile("log.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		_, err = f.WriteString(string(taskJSON) + "\n")
		if err != nil {
			panic(err)
		}

		os.Remove("current.json")
		fmt.Fprintf(os.Stdout, "Task: "+task.Name+" stop\n")
		return nil
	},
}
