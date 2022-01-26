package handle

import (
	"fmt"
	"gorm.io/gorm"
	"todo-list-server/app/model"
)

func GetTasks2(db *gorm.DB) []model.Task {
	fmt.Println("getTask")
	var tasks []model.Task
	db.Find(&tasks)

	for _, t := range tasks {
		fmt.Println(t.ToString())
	}

	return tasks
}

func GetTask2(db *gorm.DB, id int) model.Task {
	fmt.Println("getTask")
	task := model.Task{}
	db.Where("id = ?", id).First(&task)

	fmt.Println(task.ToString())

	return task
}

func InsertTask2(db *gorm.DB, task model.Task) int {
	fmt.Println("insertTask")
	result := db.Create(&task)
	return int(result.RowsAffected)
}

func UpdateTask2(db *gorm.DB, task model.Task) int {
	fmt.Println("updateTask")
	fmt.Println(task.ToString())

	result := db.Updates(&task)
	return int(result.RowsAffected)
}

func DeleteTask2(db *gorm.DB, id int) int {
	fmt.Println("deleteTask")
	result := db.Delete(model.Task{Id: id})
	return int(result.RowsAffected)
}
