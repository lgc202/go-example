package main

import "github.com/gin-gonic/gin"

func SayHello(ctx *gin.Context) {
	ctx.String(200, "hello gin")
}

func main() {
	// gin.Default 会使用 Logger 和 Recovery 这两个中间件，而 gin.New 不会
	// router := gin.New()
	router := gin.Default()

	router.Handle("GET", "/", SayHello)
	// router.GET("/", SayHello)

	router.Run(":9000")
}
