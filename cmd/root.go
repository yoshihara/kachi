package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yoshihara/kachi/lib"
	"io/ioutil"
	"os"
	"time"
	// "github.com/spf13/viper"
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
	RootCmd.AddCommand(versionCmd)
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop task tracking",
	Long:  "stop: stop task tracking",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: 関数内部はlib/配下に作ってテストを書きたい

		bytes, err := ioutil.ReadFile("current.json")
		if err != nil {
			return errors.New("[ERROR] Couldn't stop task tracking. Please check if 'current.json' exists")
		}

		var task = lib.Task{}
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
