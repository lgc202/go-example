package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {
	name := c.Query("name")
	c.String(http.StatusOK, "Hello %s", name)

	// 有默认值：
	// http://127.0.0.1:8080/user/?name=hallen
	// c.DefaultQuery("name","hallen")
	// 数组：
	// http://127.0.0.1:8080/user?name=1,2,3,4,5
	// names := c.QueryArray("name")
	// map:
	// http://127.0.0.1:8080/user?name[1]=hallen1&name[2]=hallen2
	// name_map := c.QueryMap("name")
}

func main() {
	router := gin.Default()
	// http://127.0.0.1:9000/user/?name=jack
	router.GET("/user", User)
	router.Run(":9000")
}
