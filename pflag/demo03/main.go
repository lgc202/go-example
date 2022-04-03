package main

import (
	"fmt"
	"github.com/spf13/pflag"
)

// 指定了选项但是没有指定选项值时的默认值
func main() {
	var ip = pflag.IntP("flagname", "f", 1234, "help message")
	pflag.Lookup("flagname").NoOptDefVal = "4321"

	pflag.Parse()
	fmt.Println("ip=", *ip)
}

// [root@master demo03]# go run main.go
// ip= 1234
// [root@master demo03]# go run main.go -f
// ip= 4321
// [root@master demo03]# go run main.go -f=5678
// ip= 5678
