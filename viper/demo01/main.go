package main

import (
	"fmt"
	"github.com/spf13/viper"
)

// 把值存入Viper方式一: 建立默认值
func main() {
	// 一个好的配置系统应该支持默认值。Viper 支持对 key 设置默认值，
	// 当没有通过配置文件、环境变量、远程配置或命令行标志设置 key 时，
	// 设置默认值通常是很有用的，可以让程序在没有明确指定配置时也能够正常运行
	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("LayoutDir", "layouts")
	viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})

	fmt.Printf("%+v\n", viper.Get("ContentDir"))
	fmt.Printf("%+v\n", viper.Get("LayoutDir"))
	fmt.Printf("%+v\n", viper.Get("Taxonomies"))
}

// [root@master demo01]# go run main.go
// content
// layouts
// map[category:categories tag:tags]
