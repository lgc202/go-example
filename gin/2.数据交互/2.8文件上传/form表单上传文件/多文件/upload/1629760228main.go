package main

import "github.com/gin-gonic/gin"

func SayHello(ctx *gin.Context) {
	ctx.String(200, "hello gin")
}

func main() {
	// router := gin.New()
	router := gin.Default()

	router.Handle("GET", "/", SayHello)
	// router.GET("/", SayHello)

	router.Run(":9000")
}
