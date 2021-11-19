package main

import (
	"todo-list-server/app"
	"todo-list-server/config"
)

func main() {
	dsn := config.GetPostgresDSN(config.GetPostgresConfig())

	app := &app.App{}
	app.Initialize(dsn)
	app.Run()
}
