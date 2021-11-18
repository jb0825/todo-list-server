package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"todo-list-server/config"
)

type User struct {
	Id       string `gorm:"primaryKey"`
	password string
}

type Task struct {
	Id        int `gorm:"primaryKey"`
	Name      string
	Completed bool
	UserId    string
}

func (task *Task) ToString() string {
	return fmt.Sprintf("No: %d, Name: %s, Completed: %t", task.Id, task.Name, task.Completed)
}

// Table Name 정의
type Tabler interface {
	TableName() string
}

func (Task) TableName() string {
	return "task"
}

func main() {
	config := config.GetPostgresCongif()

	dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Seoul",
		config.Host,
		config.Username,
		config.Password,
		config.Name,
		config.Port,
	)

	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// SELECT * FROM task ORDER BY id LIMIT 1
	println("First : ")
	task := Task{}

	db.First(&task)
	fmt.Println(task)

	// SELECT * FROM task
	println("Find : ")
	var tasks []Task
	db.Find(&tasks)

	for idx, t := range tasks {
		fmt.Println(idx, t)
	}

	// ORDER
	// SELECT * FROM task ORDER BY name desc
	// Multiple order 도 가능함
	println("Order : ")
	db.Order("name desc").Find(&tasks)
	for idx, t := range tasks {
		fmt.Println(idx, t)
	}

	// WHERE
	// SELECT * FROM task WHERE name = 'test' LIMIT 1
	// ?, IN, LIKE, AND, Time, BETWEEN 등 사용 가능
	println("Where : ")
	db.Where("name = ?", "test").First(&task)
	fmt.Println(task)

	// Map & Slice Conditions
	// struct
	db.Where(&Task{Name: "test", Completed: true}, "name", "completed").Find(&task)
	// Map
	db.Where(map[string]interface{}{"name": "test", "completed": true}).Find(&task)
	// Slice of primary keys
	// SELECT * FROM task WHERE id IN (1, 2)
	db.Where([]int{1, 2}).Find(&tasks)

	// Inline Condition
	db.Find(&tasks, Task{Completed: true})
	db.Find(&tasks, map[string]interface{}{"completed": true})

	// NOT
	db.Not("name = ?", "test").Find(&tasks)

	// OR
	db.Where("name = 'test'").Or("name = ?", "test2").Find(&tasks)

	// SELECT Specific Fields
	println("Selecting Specific Fields : ")
	db.Select("completed", "name").Find(&tasks)
	for idx, t := range tasks {
		fmt.Println(idx, t)
	}

	// GROUP BY & HAVING
	type group struct {
		Name  string
		Total int
	}
	var groups []group

	println("Group By & Having : ")
	db.Model(&Task{}).
		Select("name, sum(id) as total").
		Group("name").
		Having("name = ?", "test").
		Find(&groups)

	for idx, g := range groups {
		fmt.Println(idx, g)
	}

	// DISTINCT
	db.Distinct("name", "completed").Find(&tasks)

	// JOIN
	println("Join : ")

	db.Table("task").
		Select("task.id, task.name, task.completed").
		Joins("JOIN users ON users.id = task.user_id").
		Where("task.user_id = ?", "id").
		Scan(&tasks)

	for idx, t := range tasks {
		fmt.Println(idx, t)
	}

	db.Table("users").
		Select("*").
		Joins("JOIN task ON users.id = task.user_id").
		Where("task.user_id = ?", "id").
		Scan(&tasks)

}
