package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Name string
	Id   int
}

func Index(ctx *gin.Context) {
	userInfo := UserInfo{
		Name: "user-one",
		Id:   10,
	}

	ctx.HTML(http.StatusOK, "index/index.html", userInfo)
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("template/**/*")
	router.GET("/", Index)

	router.Run(":9000")
}
