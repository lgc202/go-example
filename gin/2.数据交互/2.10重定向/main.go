package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	ctx.Redirect(http.StatusFound, "/user")
}

func User(ctx *gin.Context) {
	ctx.String(http.StatusOK, "这是User页面")
}

func main() {
	engine := gin.Default()
	engine.GET("/", Index)
	engine.GET("/user", User)
	engine.Run(":9000")
}
