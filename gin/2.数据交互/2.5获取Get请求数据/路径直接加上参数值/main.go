package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}

func main() {
	router := gin.Default()
	// 下面两种都可以访问
	// http://127.0.0.1:9000/user/hallen
	// http://127.0.0.1:9000/user
	router.GET("/user/*name", User)

	// http://127.0.0.1:9000/user/hallen，这里必须指定name这个路径，不然会找不到
	// router.GET("/user/:name", User)
	router.Run(":9000")
}
