package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"time"
	// "github.com/spf13/viper"
)

// Task JSONにするときのタスクの持つ情報
type Task struct {
	Name  string
	Start string
	End   string
}

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
	RootCmd.AddCommand(versionCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start task tracking",
	Long:  "start [task name]: start task tracking",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: 関数内部はlib/配下に作ってテストを書きたい
		var task = Task{Name: args[0], Start: time.Now().String()}
		taskJSON, err := json.Marshal(&task)
		if err != nil {
			panic(err)
		}
		f, error := os.Create("current.json")
		if error != nil {
			panic(err)
		}
		defer f.Close()

		// TODO ファイルがあったらエラーメッセージ出して中止する

		writer := bufio.NewWriter(f)
		writer.WriteString(string(taskJSON))
		writer.Flush()

		fmt.Println("Task: " + args[0] + " start")
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version",
	Long:  "Print the version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("1.0.0") // TODO: gitのタグとかを入れる
	},
}
