package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 全局中间件 > 路由组中间件 > 路由中间件，如果是同一类别，那就取决于append的前后顺序了
func main() {
	router := gin.Default()
	// 绑定在根router上
	router.Use(GlobalMiddleware)
	router.GET("/", Index)
	router.Run(":9000")
}

func GlobalMiddleware(ctx *gin.Context) {
	fmt.Println("这是全局中间件")
}

func Index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "这是首页")
}
