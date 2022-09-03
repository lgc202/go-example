package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/login", login)

	// 在这里使用，访问login时不会执行中间件
	// 访问user_list和news_list时会执行中间件
	router.Use(MiddleWare)

	router.GET("/user_list", userList)
	router.GET("/news_list", newsList)
	router.Run(":9000")
}

func MiddleWare(context *gin.Context) {
	fmt.Println("认证成功")
}

func newsList(context *gin.Context) {
	context.String(http.StatusOK, "newsList")
}

func userList(context *gin.Context) {
	context.String(http.StatusOK, "userList")
}

func login(context *gin.Context) {
	context.String(http.StatusOK, "login")
}
