package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func main() {
	// 创建一个默认的路由引擎
	engine := gin.Default()
	// 注册路由,并设置一个匿名的handlers，返回JSON格式数据
	engine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "请求成功",
		})
	})
	// 获取url path中的参数
	engine.GET("user/:name/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")
		action = strings.Trim(action, "/")
		// 获取query参数
		id := context.Query("id")
		context.String(http.StatusOK, name+" is "+action+" "+id)
	})

	// 启动服务，并监听端口9090，
	// 不填默认监听 0.0.0.0:8080
	_ = engine.Run(":9090")
}
