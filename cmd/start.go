package cmd

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/yoshihara/kachi/lib"
	"os"
	"time"
)

// TODO: エラーチェックは予想不可能なものはpanic、それ以外はメッセージを出す
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start task tracking",
	Long:  "start [task name]: start task tracking",
	RunE: func(cmd *cobra.Command, args []string) error {
		// TODO: 関数内部はlib/配下に作ってテストを書きたい
		layout := "Mon Jan 2 15:04:05 MST 2006"

		var task = lib.Task{Name: args[0], Start: time.Now().Format(layout)}
		taskJSON, err := json.Marshal(&task)
		if err != nil {
			panic(err)
		}

		_, err = os.Stat("current.json")
		if !os.IsNotExist(err) {
			return errors.New("[ERROR] 'current.json' exist already. Please stop before start")
		}

		f, error := os.Create("current.json")
		if error != nil {
			return errors.New("[ERROR] Couldn't create 'current.json' file")
		}
		defer f.Close()

		writer := bufio.NewWriter(f)
		writer.WriteString(string(taskJSON))
		writer.Flush()

		fmt.Println("Task: " + args[0] + " start")
		return nil
	},
}
