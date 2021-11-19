package app

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"todo-list-server/app/handle"
	"todo-list-server/app/model"
)

type App struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func (app *App) Initialize(dsn string) {

	/*
		var err error
		app.DB, err = sql.Open(config.Dialect, dbURI)
		if err != nil {
			log.Fatal("Database Connect Failed.")
		}
	*/

	// gorm 사용
	var err error
	app.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
		context.JSON(http.StatusOK, handle.GetTasks2(app.DB))
	})

	app.Router.GET("/task/:id", func(context *gin.Context) {
		id, _ := strconv.Atoi(context.Param("id"))
		context.JSON(http.StatusOK, handle.GetTask2(app.DB, id))
	})

	app.Router.POST("/task", func(context *gin.Context) {
		task := model.Task{}
		context.BindJSON(&task)

		result := handle.InsertTask2(app.DB, task)
		context.Status(resultChecker(result))
	})

	app.Router.PATCH("/task/:id", func(context *gin.Context) {
		task := model.Task{}
		context.Bind(&task)

		id, _ := strconv.Atoi(context.Param("id"))
		task.Id = id

		result := handle.UpdateTask2(app.DB, task)
		context.Status(resultChecker(result))
	})

	app.Router.DELETE("/task/:id", func(context *gin.Context) {
		id, _ := strconv.Atoi(context.Param("id"))
		result := handle.DeleteTask2(app.DB, id)

		context.Status(resultChecker(result))
	})
}

func resultChecker(result int) int {
	if result > 0 {
		return http.StatusOK
	} else {
		return http.StatusBadRequest
	}
}
