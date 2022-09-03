package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	names := []string{"lena", "austin", "foo"}

	// 将输出：while(1);["lena","austin","foo"]
	ctx.SecureJSON(http.StatusOK, names)
}

func main() {
	engine := gin.Default()
	engine.GET("/", Index)
	engine.Run(":9000")
}
