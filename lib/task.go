package lib

// Task JSONにするときのタスクの持つ情報
type Task struct {
	Name  string
	Start string
	End   string
}

// DateTimeLayout using datetime format
const DateTimeLayout string = "Mon Jan 2 15:04:05 MST 2006"
