package main

import (
	"fmt"
	"github.com/spf13/pflag"
)

// 弃用标志或者标志的简写
func main() {
	var logMode = pflag.StringP("logmode", "l", "debug", "Input log mode")
	// 弃用名为 logmode 的标志，并告知用户应该使用哪个标志代替
	pflag.CommandLine.MarkDeprecated("logmode", "please use --log-mode instead")

	// 保留名为 port 的标志，但是弃用它的简写形式。
	var port int
	pflag.IntVarP(&port, "port", "P", 3306, "MySQL service host port.")
	pflag.CommandLine.MarkShorthandDeprecated("port", "please use --port only")

	pflag.Parse()
	fmt.Println("logMode=", *logMode)
	fmt.Println("port=", port)
}

// [root@master demo04]# go run main.go -l=trace
// Flag --logmode has been deprecated, please use --log-mode instead
// logMode= trace

// [root@master demo04]# go run main.go --logmode=trace
// Flag --logmode has been deprecated, please use --log-mode instead
// logMode= trace

// [root@master demo04]# go run main.go --port=6379
// logMode= debug
// port= 6379

// [root@master demo04]# go run main.go -P=6379
// Flag shorthand -P has been deprecated, please use --port only
// logMode= debug
// port= 6379
