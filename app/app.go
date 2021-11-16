package app

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
	"todo-list-server/app/handle"
	"todo-list-server/app/model"
	"todo-list-server/config"
)

type App struct {
	Router  *gin.Engine
	DB		*sql.DB
}

func (app *App) Initialize(config *config.DBConfig) {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
		config.Charset,
	)

	fmt.Println(dbURI)
	fmt.Println(config.Dialect)

	var err error
	app.DB, err = sql.Open(config.Dialect, dbURI)
	if err != nil { log.Fatal("Database Connect Failed.") }

	app.Router = gin.Default()
	app.SetRouters()
}

func (app *App) Run() {
	app.Router.Run(":8080")
}

func (app *App) SetRouters() {
	app.Router.GET("/tasks", func(context *gin.Context) {
		context.JSON(http.StatusOK, handle.GetTasks(app.DB))
	})

	app.Router.GET("/task/:no", func(context *gin.Context) {
		no, _ := strconv.Atoi(context.Param("no"))
		context.JSON(http.StatusOK, handle.GetTask(app.DB, no))
	})

	app.Router.POST("/task", func(context *gin.Context) {
		name := context.PostForm("name")
		fmt.Println(name)

		handle.InsertTask(app.DB, name)
		context.Status(http.StatusOK)
	})

	app.Router.PATCH("/task/:no", func(context *gin.Context) {
		task := model.Task{}
		context.Bind(&task)

		no, _ := strconv.Atoi(context.Param("no"))
		task.No = no

		fmt.Println(task)

		var code int
		if handle.UpdateTask(app.DB, task) > 0 { code = http.StatusOK
		} else { code = http.StatusBadRequest}

		context.Status(code)
	})

	app.Router.DELETE("/task/:no", func(context *gin.Context) {
		no, _ := strconv.Atoi(context.Param("no"))

		var code int
		if handle.DeleteTask(app.DB, no) > 0 { code = http.StatusOK
		} else { code = http.StatusBadRequest}

		context.Status(code)
	})

}

