package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"net/http"
)

func main() {
	//创建一个服务
	ginServer := gin.Default()

	// 导入静态文件搜索路径
	ginServer.LoadHTMLGlob("404.html")

	ginServer.Use(favicon.New("./favicon.ico"))

	ginServer.GET("/", func(context *gin.Context) {
		//context.String(http.StatusOK, "Hello World!")
		context.JSON(http.StatusOK, gin.H{"hello": "world"})

	})

	ginServer.POST("/user", func(context *gin.Context) {
		//context.String(http.StatusOK, "Hello World!")
		context.JSON(http.StatusOK, gin.H{"msg": "hello, world!"})

	})

	// 路由分组, 分了shopping和blog两个组
	shop := ginServer.Group("/shopping")
	{
		shop.GET("index", sIndexHandler) // /shopping/index
		shop.GET("home", sHomeHandler)   // /shopping/home
	}
	blog := ginServer.Group("blog")
	{
		blog.GET("index", liveIndexHandler) // /blog/index
		blog.GET("home", liveHomeHandler)   // /blog/home
	}

	// 路由组嵌套
	app01 := ginServer.Group("app01")
	{
		app01.GET("index", sIndexHandler) // /app01/index
		ashop := app01.Group("shopping")
		{
			ashop.GET("bed", ashopBedHandler)   // /app01/shopping/bed
			ashop.GET("food", ashopFoodHandler) // /app01/shopping/food
		}
	}

	// 获取基本请求数据和请求参数
	ginServer.GET("test", func(context *gin.Context) {
		fmt.Println(context.Request.Method)
		fmt.Println(context.Request.URL)
		fmt.Println(context.Request.RemoteAddr)
		fmt.Println(context.Request.Header)

		fmt.Println(context.GetHeader("User-Agent"))
		fmt.Println(context.ClientIP())

		name := context.Query("name")
		fmt.Println(name)
		name = context.DefaultQuery("name", "zhangSan")
		fmt.Println(name)

		id, ok := context.GetQuery("id")
		if !ok {
			fmt.Println("参数不存在")
		} else {
			fmt.Println(id)
		}
		context.JSON(http.StatusOK, gin.H{"getQuery": name})
	})

	// 获取post请求的参数
	ginServer.POST("post", func(context *gin.Context) {
		name := context.PostForm("name")
		fmt.Println(name)
		name = context.DefaultPostForm("name", "xiaoMing")
		fmt.Println(name)

		id, ok := context.GetPostForm("id")
		if !ok {
			fmt.Printf("no id var in post form: %s\n", id)
		} else {
			fmt.Printf("id in post form is: %s\n", id)
		}
		context.JSON(http.StatusOK, gin.H{"getPostForm": name})
	})

	// 获取url path中的路径参数
	ginServer.GET("/user/:id", func(context *gin.Context) {
		id := context.Param("id")
		fmt.Println(id)
		context.JSON(http.StatusOK, gin.H{"getPathParams": id})
	})

	// url路径不存在，返回指定的html页面
	ginServer.NoRoute(func(context *gin.Context) {
		//context.JSON(http.StatusOK, gin.H{"msg": "Bye!"})
		context.HTML(http.StatusNotFound, "404.html", gin.H{})
	})

	err := ginServer.Run("127.0.0.1:8082")
	if err != nil {
		return
	}
}

func ashopFoodHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"Food": "duck"})
}

func ashopBedHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"Bed": "xiMengSi"})
}

func liveHomeHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"blogHome": "moHuiShou"})
}

func liveIndexHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"blogIndex": "docker"})
}

func sHomeHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"shopHome": "all"})
}

func sIndexHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"shopIndex": "apple"})
}
