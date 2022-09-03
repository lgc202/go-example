package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexAdd(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index/index.html", nil)
}

func DoAddUser(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	fmt.Println(username)
	fmt.Println(password)
	messgaeMap := map[string]interface{}{
		"code": 200,
		"msg":  "提交成功",
	}
	ctx.JSON(http.StatusOK, messgaeMap)
}

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("template/**/*")
	engine.Static("/static", "static")
	engine.GET("/add_user", IndexAdd)
	engine.POST("/do_add", DoAddUser)
	engine.Run(":9000")
}
