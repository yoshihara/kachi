package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yoshihara/kachi/lib"
	"os"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop task tracking",
	Long:  "stop: Stop task tracking",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !lib.IsJSONDirExists() {
			return lib.ErrorNotFoundJSONDir()
		}

		task, error := lib.ReadCurrentTask()
		if error != nil {
			return error
		}

		lib.CompleteTask(task)

		fmt.Fprintf(os.Stdout, "Task: "+task.Name+" stop\n")
		return nil
	},
}
