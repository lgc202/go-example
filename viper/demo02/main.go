package main

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg  = pflag.StringP("config", "c", "", "Configuration file.")
	help = pflag.BoolP("help", "h", false, "Show this help message.")
)

// 把值存入Viper方式二: 读取配置文件
func main() {
	pflag.Parse()
	if *help {
		pflag.Usage()
		return
	}

	// 从配置文件中读取配置
	if *cfg != "" {
		viper.SetConfigFile(*cfg)   // 指定配置文件名
		viper.SetConfigType("yaml") // 如果配置文件名中没有文件扩展名，则需要指定配置文件的格式，告诉viper以何种格式解析文件
	} else {
		viper.AddConfigPath(".")          // 把当前目录加入到配置文件的搜索路径中
		viper.AddConfigPath("$HOME/.iam") // 配置文件搜索路径，可以设置多个配置文件搜索路径
		viper.SetConfigName("config")     // 配置文件名称（没有文件扩展名）
	}

	if err := viper.ReadInConfig(); err != nil { // 读取配置文件。如果指定了配置文件名，则使用指定的配置文件，否则在注册的搜索路径中搜索
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	fmt.Printf("Used configuration file is: %s\n", viper.ConfigFileUsed())
}

// [root@master demo02]# go run main.go
// Used configuration file is: /root/workplace/go-example/viper/demo02/config.yaml

// [root@master demo02]# go run main.go -c=config/config_test
// Used configuration file is: config/config_test
