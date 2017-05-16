package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yoshihara/kachi/lib"
	"os"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start task tracking",
	Long:  "start [task name]: Start task tracking",
	RunE: func(cmd *cobra.Command, args []string) error {
		task, error := lib.StartTask(args[0])

		if error != nil {
			return error
		}

		fmt.Fprintf(os.Stdout, "Task: "+task.Name+" start\n")
		return nil
	},
}
