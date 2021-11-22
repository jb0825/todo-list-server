package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.String(200, "Hello Gin")
	})
	router.RunTLS(":8080", "인증서 경로.pem", "인증서 경로.key")
}
