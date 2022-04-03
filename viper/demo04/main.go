package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// 把值存入Viper方式四: 设置配置值
func main() {
	// 我们可以通过 viper.Set() 函数来显式设置配置
	viper.Set("user.username", "colin")
	fmt.Println("user.username= ", viper.Get("user.username"))
}

// [root@master demo04]# go run main.go
// user.username=  colin
