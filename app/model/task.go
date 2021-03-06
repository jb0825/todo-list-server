package model

import (
	"fmt"
)

type Tabler interface {
	TableName() string
}

func (Task) TableName() string {
	return "task"
}

type Task struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
	UserId    string `json:"-" gorm:"-"`
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
