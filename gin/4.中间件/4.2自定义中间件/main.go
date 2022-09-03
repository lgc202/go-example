package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MiddleWare1(ctx *gin.Context) {
	fmt.Println("这是自定义中间件方式1")
}

func MiddleWare2() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("这是自定义中间件方式2")
	}
}

func Index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "这是首页")
}

func main() {
	router := gin.Default()
	// 注意使用方式也不太一样
	router.Use(MiddleWare1, MiddleWare2())
	router.GET("/", Index)
	router.Run(":9000")
}
