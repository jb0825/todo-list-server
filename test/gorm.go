package test

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"todo-list-server/config"
)

type Task struct {
	Id   int
	Name string
}

func main() {
	config := config.GetConfig()
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&charset=utf8",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
		config.Charset)

	db, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{})

}
