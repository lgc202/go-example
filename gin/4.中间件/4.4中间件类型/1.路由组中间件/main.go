package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("v1")
	// 访问 http://127.0.0.1:9000/v1/get_name和http://127.0.0.1:9000/v1/get_id 都会使用到中间件
	v1.Use(Middleware)
	{
		v1.GET("/get_name", GetName)
		v1.GET("/get_id", GetId)
	}

	router.Run(":9000")
}

func GetId(ctx *gin.Context) {
	ctx.String(http.StatusOK, "获取Id")
}

func GetName(ctx *gin.Context) {
	ctx.String(http.StatusOK, "获取姓名")
}

func Middleware(ctx *gin.Context) {
	fmt.Println("这是路由组中间件")
}
