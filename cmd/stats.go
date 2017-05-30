package cmd

import (
	"bytes"
	"github.com/spf13/cobra"
	"github.com/yoshihara/kachi/lib"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Show tracked tasks",
	Long:  "stats: Show tracked tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !lib.IsJSONDirExists() {
			return lib.ErrorNotFoundJSONDir()
		}

		tasksJSON, error := lib.ReadCompletedTasks()
		if error != nil {
			return error
		}

		for _, taskJSON := range bytes.Split(tasksJSON, []byte("\n")) {
			if len(taskJSON) < 1 {
				continue
			}

			error = lib.ShowTaskStat(taskJSON, scale)
			if error != nil {
				return error
			}

		}
		return nil
	},
}
