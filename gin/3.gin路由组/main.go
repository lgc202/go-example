package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/login", loginEndpoint)
		v1.GET("/submit", submitEndpoint)
		v1.GET("/read", readEndpoint)
	}

	router.Run(":9000")
}

func readEndpoint(ctx *gin.Context) {
	ctx.String(http.StatusOK, "readEndpoint")
}

func submitEndpoint(ctx *gin.Context) {
	ctx.String(http.StatusOK, "submitEndpoint")
}

func loginEndpoint(ctx *gin.Context) {
	ctx.String(http.StatusOK, "loginEndpoint")
}
