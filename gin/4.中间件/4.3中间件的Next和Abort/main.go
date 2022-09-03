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

		// 满足条件， 后面的中间件不会被执行
		if 3 < 4 {
			ctx.Abort()
		}
		ctx.Next()
		fmt.Println("这是自定义中间件2--结束")
	}
}

func MiddleWare3(ctx *gin.Context) {
	fmt.Println("这是自定义中间件3--开始")
	ctx.Next()
	fmt.Println("这是自定义中间件3--结束")
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
	router.Use(MiddleWare1, MiddleWare2(), MiddleWare3)
	router.GET("/", Index)
	router.Run(":9000")
}
