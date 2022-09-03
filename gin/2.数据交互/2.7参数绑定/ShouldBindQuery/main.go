package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id   int    `form:"id"`
	Name string `form:"name"`
}

func Index(ctx *gin.Context) {
	var u User
	// 只绑定 url 查询参数(get)而忽略 post 数据
	err := ctx.ShouldBindQuery(&u)
	fmt.Println(err)
	fmt.Println(u)
	ctx.String(http.StatusOK, "Hello %s", u.Name)

}

// http://127.0.0.1:9000/?id=1&name=lgc
func main() {
	engine := gin.Default()
	engine.GET("/", Index)
	engine.Run(":9000")
}
