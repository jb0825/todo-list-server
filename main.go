package main

import (
	"todo-list-server/app"
	"todo-list-server/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run()
}