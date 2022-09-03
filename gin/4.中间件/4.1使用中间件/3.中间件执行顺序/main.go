package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MiddleWare1(ctx *gin.Context) {
	fmt.Println("这是自定义中间件1--开始")
	ctx.Next()
	fmt.Println("这是自定义中间件1--结束")
}

func MiddleWare2() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("这是自定义中间件2--开始")
		ctx.Next()
		fmt.Println("这是自定义中间件2--结束")
	}
}

func Index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "这是首页")
}

// 输出
// 这是自定义中间件1--开始
// 这是自定义中间件2--开始
// 这是自定义中间件2--结束
// 这是自定义中间件1--结束
func main() {
	router := gin.Default()
	router.Use(MiddleWare1, MiddleWare2())
	router.GET("/", Index)
	router.Run(":9000")
}
