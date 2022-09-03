package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	name := "hello world"
	ctx.HTML(http.StatusOK, "index/index.html", name)
}

func User(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "user/user.html", nil)
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("template/**/*")
	router.Static("/static", "static")
	// router.StaticFS("/static", http.Dir("static"))
	router.GET("/", Index)
	router.GET("/user", User)

	router.Run(":9000")
}
