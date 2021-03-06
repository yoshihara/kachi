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

var scale float64

func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(initCmd)
	RootCmd.AddCommand(startCmd)
	RootCmd.AddCommand(stopCmd)
	RootCmd.AddCommand(currentCmd)

	statsCmd.Flags().Float64VarP(&scale, "scale", "s", 1.0, "scale for stats")
	RootCmd.AddCommand(statsCmd)
	RootCmd.AddCommand(refreshCmd)
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
