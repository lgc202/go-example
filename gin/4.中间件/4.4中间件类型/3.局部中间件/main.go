package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", Middleware, Index)
	router.Run(":9000")
}

func Index(ctx *gin.Context) {
	// 获取在中间件中设置到值
	val, ok := ctx.Get("example")
	if ok {
		fmt.Println(val)
	}
	ctx.String(http.StatusOK, "这是首页")
}

func Middleware(ctx *gin.Context) {
	t := time.Now()
	fmt.Println("我是自定义中间件第1种定义方式---请求之前")
	//在gin上下文中定义一个变量
	ctx.Set("example", "CustomRouterMiddle")
	//请求之前
	ctx.Next()
	fmt.Println("我是自定义中间件第1种定义方式---请求之后")
	//请求之后
	//计算整个请求过程耗时
	t2 := time.Since(t)
	fmt.Println(t2)
}
