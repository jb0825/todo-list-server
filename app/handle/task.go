package handle

import (
	"database/sql"
	"fmt"
	"log"
	"todo-list-server/app/model"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetTasks(db *sql.DB) []model.Task {
	var tasks []model.Task

	rows, err := db.Query("SELECT * FROM task")
	checkError(err)

	task := model.Task{}
	for rows.Next() {
		err := rows.Scan(&task.No, &task.Name, &task.Completed)
		checkError(err)

		tasks = append(tasks, task)
	}

	for _, task := range tasks {
		fmt.Println(task.ToString())
	}

	return tasks
}

func GetTask(db *sql.DB, no int) model.Task {
	task := model.Task{}
	err := db.QueryRow("SELECT * FROM task WHERE no=?", no).Scan(&task.No, &task.Name, &task.Completed)
	checkError(err)

	fmt.Println(task.ToString())

	return task
}

func InsertTask(db *sql.DB, name string) {
	_, err := db.Exec("INSERT INTO task VALUES (nextval('task'), ?, false)", name)
	checkError(err)
}

func UpdateTask(db *sql.DB, task model.Task) int {
	result, err := db.Exec("UPDATE task SET name=?, completed=? WHERE no=?", task.Name, task.Completed, task.No)
	checkError(err)

	cnt, _ := result.RowsAffected()
	return int(cnt)
}

func DeleteTask(db *sql.DB, no int) int {
	result, err := db.Exec("DELETE FROM task WHERE no=?", no)
	checkError(err)

	cnt, _ := result.RowsAffected()
	return int(cnt)
}