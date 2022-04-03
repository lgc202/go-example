package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 把值存入Viper方式三: 监听和重新读取配置文件
func main() {
	viper.AddConfigPath(".")      // 把当前目录加入到配置文件的搜索路径中
	viper.SetConfigName("config") // 配置文件名称（没有文件扩展名）

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})

	select {}
}

// [root@master demo03]# go run main.go
// 修改配置并保存后会输出
// Config file changed: /root/workplace/go-example/viper/demo03/config.yaml
