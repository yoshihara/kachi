package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yoshihara/kachi/lib"
	"os"
)

var refreshCmd = &cobra.Command{
	Use:   "refresh",
	Short: "Move completed tasks to archive",
	Long:  "refresh: Move completed tasks to archive",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !lib.IsJSONDirExists() {
			return lib.ErrorNotFoundJSONDir()
		}

		error := lib.RefreshTasks()

		if error != nil {
			return error
		}
		fmt.Fprintf(os.Stdout, "Refresh log.json and move completed tasks to archive.json\n")

		return nil
	},
}
