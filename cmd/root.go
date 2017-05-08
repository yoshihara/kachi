package cmd

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
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
	RootCmd.AddCommand(stopCmd)
	RootCmd.AddCommand(versionCmd)
}

// TODO: エラーチェックは予想不可能なものはpanic、それ以外はメッセージを出す
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start task tracking",
	Long:  "start [task name]: start task tracking",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: 関数内部はlib/配下に作ってテストを書きたい
		var task = Task{Name: args[0], Start: time.Now().String()}
		taskJSON, err := json.Marshal(&task)
		if err != nil {
			panic(err)
		}

		_, err = os.Stat("current.json")
		if !os.IsNotExist(err) {
			return errors.New("[ERROR] 'current.json' exist already. Please stop before start.")
		}

		f, error := os.Create("current.json")
		if error != nil {
			return errors.New("[ERROR] Couldn't create 'current.json' file.")
		}
		defer f.Close()

		writer := bufio.NewWriter(f)
		writer.WriteString(string(taskJSON))
		writer.Flush()

		fmt.Println("Task: " + args[0] + " start")
		return nil
	},
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop task tracking",
	Long:  "stop: stop task tracking",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: 関数内部はlib/配下に作ってテストを書きたい

		bytes, err := ioutil.ReadFile("current.json")
		if err != nil {
			return errors.New("[ERROR] Couldn't stop task tracking. Please check if 'current.json' exists.")
		}

		var task = Task{}
		err = json.Unmarshal(bytes, &task)
		if err != nil {
			panic(err)
		}

		task.End = time.Now().String()
		taskJSON, err := json.Marshal(&task)
		if err != nil {
			panic(err)
		}

		f, err := os.OpenFile("log.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		_, err = f.WriteString(string(taskJSON) + "\n")
		if err != nil {
			panic(err)
		}

		os.Remove("current.json")
		fmt.Println("Task: " + task.Name + " stop")
		return nil
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
