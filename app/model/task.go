package model

import (
	"fmt"
)

type Task struct {
	No 	 	  int		`json:"-"`
	Name 	  string	`json:"name"`
	Completed bool		`json:"completed"`
}

func NewTask(no int, name string, completed bool) *Task {
	t := &Task {
		No: no,
		Name: name,
		Completed: completed,
	}

	return t
}

func (task *Task) ToString() string {
	return fmt.Sprintf("No: %d, Name: %s, Completed: %t", task.No, task.Name, task.Completed)
}