package model

import (
	"fmt"
)

type Task struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

func NewTask(no int, name string, completed bool) *Task {
	t := &Task{
		Id:        no,
		Name:      name,
		Completed: completed,
	}

	return t
}

func (task *Task) ToString() string {
	return fmt.Sprintf("No: %d, Name: %s, Completed: %t", task.Id, task.Name, task.Completed)
}
