package main

import (
	"fmt"
	"github.com/spf13/pflag"
)

// 获取非选项参数
func main() {
	pflag.Parse()
	fmt.Printf("argument number is: %v\n", pflag.NArg())
	fmt.Printf("argument list is: %v\n", pflag.Args())
	fmt.Printf("the first argument is: %v\n", pflag.Arg(0))
}

// [root@master demo02]# go run main.go arg1 arg2
// argument number is: 2
// argument list is: [arg1 arg2]
// the first argument is: arg1
