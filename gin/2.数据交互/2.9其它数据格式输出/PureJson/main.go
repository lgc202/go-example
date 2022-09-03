package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	ctx.PureJSON(http.StatusOK, gin.H{
		"code": 200,
		"tag":  "<br>",
		"msg":  "提交成功",
	})
}

func main() {
	engine := gin.Default()
	engine.GET("/", Index)
	engine.Run(":9000")
}
