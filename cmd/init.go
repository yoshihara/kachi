package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yoshihara/kachi/lib"
	"os"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize directory for kachi",
	Long:  "init: Initialize directory for kachi",
	RunE: func(cmd *cobra.Command, args []string) error {
		kachiDir := lib.JSONDir()

		_, error := os.Stat(kachiDir)
		if error == nil {
			return errors.New("[ERROR] .kachi directory is aleady exists")
		}

		error = os.Mkdir(kachiDir, os.ModeDir|0755)
		if error != nil {
			panic("[ERROR] Couldn't create '.kachi' dir in home dir. abort")
		}
		fmt.Fprintf(os.Stdout, ".kachi directory was created. Enjoy\n")
		return nil
	},
}
