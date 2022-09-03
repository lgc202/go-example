package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// 加密的盐
	store := cookie.NewStore([]byte("hallen"))
	// 使用 session 中间件
	router.Use(sessions.Sessions("gin_session", store))
	// 设置路由
	router.GET("session_test", SessionTest)

	router.Run(":9000")
}

func SessionTest(c *gin.Context) {
	// 初始化session对象
	session := sessions.Default(c)
	// 设置session
	session.Set("name", "halllen")
	session.Set("age", 18)
	session.Set("addr", "xxx")

	// 获取session
	name := session.Get("name")
	fmt.Println("===================")
	fmt.Println("name = ", name)

	// 删除指定key的session
	session.Delete("name")
	name = session.Get("name")
	fmt.Println("===================")
	fmt.Println("name = ", name)

	// 删除所以的session
	session.Clear()
	name = session.Get("age")
	fmt.Println("===================")
	fmt.Println("age = ", name) // 获取不到
}
