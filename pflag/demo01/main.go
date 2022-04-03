package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
)

var name string

// 支持长选项、默认值和使用文本，并将标志的值存储在指针中
var age = flag.Int("age", 18, "Input Your Age")

func init() {
	// 支持长选项、短选项、默认值和使用文本，并将标志的值绑定到变量
	flag.StringVarP(&name, "name", "n", "ljj", "Input Your Name")
}

// 支持多种命令行参数定义方式
func main() {
	flag.Parse()
	// go run main.go --help 查看用法
	fmt.Println("name=", name)
	fmt.Println("age=", *age)
}
