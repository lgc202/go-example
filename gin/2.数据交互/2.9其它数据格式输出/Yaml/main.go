package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	ctx.YAML(http.StatusOK, gin.H{
		"code": 200,
		"tag":  "<br>",
		"user": gin.H{"name": "zhiliao", "age": 18},
		"html": "<b>Hello, world!</b>",
	})
}

func main() {
	engine := gin.Default()
	engine.GET("/", Index)
	engine.Run(":9000")
}
