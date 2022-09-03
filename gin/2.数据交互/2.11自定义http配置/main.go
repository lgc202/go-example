package main

import (
	"net/http"
	"time"

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

	// http.ListenAndServe(":9000", engine)
	// 或者
	server := http.Server{
		Addr:           ":9000",
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
