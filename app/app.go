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
	Router *gin.Engine
	DB     *sql.DB
}

func (app *App) Initialize(config *config.DBConfig) {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
		config.Charset)

	fmt.Println(dbURI)
	fmt.Println(config.Dialect)

	var err error
	app.DB, err = sql.Open(config.Dialect, dbURI)
	if err != nil {
		log.Fatal("Database Connect Failed.")
	}

	app.Router = gin.New()
	app.Router.Use(CORSMiddleware())
	app.SetRouters()

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PATCH, OPTIONS, GET, DELETE, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (app *App) Run() {
	app.Router.Run(":8080")
}

func (app *App) SetRouters() {
	app.Router.GET("/tasks", func(context *gin.Context) {
		context.JSON(http.StatusOK, handle.GetTasks(app.DB))
	})

	app.Router.GET("/task/:id", func(context *gin.Context) {
		id, _ := strconv.Atoi(context.Param("id"))
		context.JSON(http.StatusOK, handle.GetTask(app.DB, id))
	})

	app.Router.POST("/task", func(context *gin.Context) {
		task := model.Task{}
		context.BindJSON(&task)

		handle.InsertTask(app.DB, task.Name)
		context.Status(http.StatusOK)
	})

	app.Router.PATCH("/task/:id", func(context *gin.Context) {
		task := model.Task{}
		context.Bind(&task)

		id, _ := strconv.Atoi(context.Param("id"))
		task.Id = id

		var code int
		if handle.UpdateTask(app.DB, task) > 0 {
			code = http.StatusOK
		} else {
			code = http.StatusBadRequest
		}

		context.Status(code)
	})

	app.Router.DELETE("/task/:id", func(context *gin.Context) {
		id, _ := strconv.Atoi(context.Param("id"))

		var code int
		if handle.DeleteTask(app.DB, id) > 0 {
			code = http.StatusOK
		} else {
			code = http.StatusBadRequest
		}

		context.Status(code)
	})

}
