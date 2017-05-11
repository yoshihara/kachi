package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yoshihara/kachi/lib"
	"os"
)

// TODO: エラーチェックは予想不可能なものはpanic、それ以外はメッセージを出す
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start task tracking",
	Long:  "start [task name]: start task tracking",
	RunE: func(cmd *cobra.Command, args []string) error {
		task, error := lib.StartTask(args[0])

		if error != nil {
			return error
		}

		fmt.Fprintf(os.Stdout, "Task: "+task.Name+" start\n")
		return nil
	},
}
