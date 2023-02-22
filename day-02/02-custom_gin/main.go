package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"github.com/thinkerou/favicon"
	"log"
	"net/http"
	"time"
)

type User struct {
	Name  string `json:"name" form:"name" binding:"required"`
	Email string `json:"email" form:"email"`
}

func main() {
	//创建一个服务
	ginServer := gin.Default()
	// 默认使用了2个中间件Logger(), Recovery()

	// 导入静态文件搜索路径
	ginServer.LoadHTMLGlob("404.html")
	// 静态文件配置
	ginServer.Static("/static", "./static")

	ginServer.Use(favicon.New("./favicon.ico"))

	ginServer.GET("/", func(context *gin.Context) {
		//context.String(http.StatusOK, "Hello World!")  // 响应简单字符串
		context.JSON(http.StatusOK, gin.H{"hello": "world"}) // 响应json

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

	// ShouldBind强大的功能，可以基于请求的Content-Type识别请求数据类型并利用反射机制自动提取请求中querystring、form表单、JSON、XML等参数到结构体中。
	//能够基于请求自动提取JSON、form表单和querystring类型的数据，并把值绑定到指定的结构体对象。
	ginServer.POST("/users/:id", func(context *gin.Context) {
		u := User{}
		if context.ShouldBind(&u) == nil {
			log.Println(":::", u.Name)
			log.Println(":::", u.Email)
		}
		context.String(http.StatusOK, "success")
	})

	ginServer.GET("/xml", func(context *gin.Context) { // 响应xml格式
		context.XML(http.StatusOK, gin.H{"message": "xml消息", "data": "hello world"})
	})

	// 重定向
	ginServer.GET("redirect", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "/xml")

		// 方式二
		//context.Request.URL.Path = "/xml"
		//ginServer.HandleContext(context)
	})

	// 异步响应
	ginServer.GET("async", func(context *gin.Context) {
		// 需要搞一个副本
		copyContext := context.Copy()
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步执行： " + copyContext.Request.URL.Path)
		}()
		context.JSON(http.StatusOK, gin.H{"msg": "异步执行中..."})
	})

	ginServer.GET("/protobuffer", ProtoBufferHandler)
	// url路径不存在，返回指定的html页面
	ginServer.NoRoute(func(context *gin.Context) {
		//context.JSON(http.StatusOK, gin.H{"msg": "Bye!"})
		context.HTML(http.StatusNotFound, "404.html", gin.H{}) // 响应html
	})

	err := ginServer.Run("127.0.0.1:8082")
	if err != nil {
		return
	}
}

func ProtoBufferHandler(c *gin.Context) {
	resp := []int64{int64(9), int64(55)}
	label := "你好啊，今天天气挺好的"
	data := &protoexample.Test{
		Label: &label,
		Reps:  resp,
	}
	fmt.Println(data)
	c.ProtoBuf(http.StatusOK, data)
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
