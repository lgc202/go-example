package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexAdd(ctx  *gin.Context)  {
	ctx.HTML(http.StatusOK, "index/index.html", nil)
}

func DoAddUser(ctx *gin.Context){
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	fmt.Println(username)
	fmt.Println(password)
	ctx.String(http.StatusOK, "add success")
}

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/**/*")
	engine.GET("/add_user", IndexAdd)
	engine.POST("/add_user", DoAddUser)
	engine.Run(":9000")
}
