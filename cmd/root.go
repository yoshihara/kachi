package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// RootCmd 引数無しで実行したコマンド
var RootCmd = &cobra.Command{
	Use:   "kachi",
	Short: "simple time tracking app",
	Long:  "simple time tracking app",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`Please run 'help' command.
   e.g.) kachi help`)
	},
}

func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(startCmd)
	RootCmd.AddCommand(stopCmd)
	RootCmd.AddCommand(statsCmd)
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version",
	Long:  "Print the version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("1.0.0") // TODO: gitのタグとかを入れる
	},
}
