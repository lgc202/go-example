package main

import (
	"fmt"
	"github.com/spf13/pflag"
)

func main() {
	var secretFlag = pflag.StringP("secretFlag", "s", "secret", "Input secret Flag")

	// 在帮助文档中隐藏参数 secretFlag
	// 其实在把参数标记为废弃时，同时也会设置隐藏参数
	pflag.CommandLine.MarkHidden("secretFlag")

	pflag.Parse()
	fmt.Println("secretFlag=", *secretFlag)
}

// [root@master demo05]# go run main.go --help
// Usage of /tmp/go-build3351075044/b001/exe/main:
// pflag: help requested
// exit status 2

// [root@master demo05]# go run main.go
// secretFlag= secret

// [root@master demo05]# go run main.go -s=public
// secretFlag= public
