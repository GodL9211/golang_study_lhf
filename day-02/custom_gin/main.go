package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"net/http"
)

func main() {
	//创建一个服务
	ginServer := gin.Default()
	ginServer.Use(favicon.New("./favicon.ico"))

	ginServer.GET("/", func(context *gin.Context) {
		//context.String(http.StatusOK, "Hello World!")
		context.JSON(http.StatusOK, gin.H{"hello": "world"})

	})

	ginServer.POST("/user", func(context *gin.Context) {
		//context.String(http.StatusOK, "Hello World!")
		context.JSON(http.StatusOK, gin.H{"msg": "hello, world!"})

	})

	err := ginServer.Run(":8082")
	if err != nil {
		return
	}
}
