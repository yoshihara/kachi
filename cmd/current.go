package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yoshihara/kachi/lib"
	"os"
)

var currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Show current task",
	Long:  "current: Show current task",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !lib.IsJSONDirExists() {
			return lib.ErrorNotFoundJSONDir()
		}

		task, error := lib.ReadCurrentTask()
		if error != nil {
			return error
		}

		fmt.Fprintf(os.Stdout, "Task: "+task.Name+"\nStart:"+task.Start+"\n")
		return nil
	},
}
