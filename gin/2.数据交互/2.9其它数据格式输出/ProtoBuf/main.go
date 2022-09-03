package main

import (
	"firshdemo/user"
	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	userData := &user.User{
		Name: "zhiliao",
		Age:  18,
	}
	ctx.ProtoBuf(200, userData)
}

// protoc -I . --go_out=plugins=grpc:. user.proto
func main() {
	engine := gin.Default()
	engine.GET("/", Index)
	engine.Run(":9000")
}
