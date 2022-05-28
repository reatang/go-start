package config

import (
	"flag"
	"github.com/spf13/viper"
)

var configFile string
var configPath string

func init() {
	flag.StringVar(&configFile, "config", "app.yaml", "配置文件名称")
	flag.StringVar(&configPath, "config-path", "./configs/", "配置文件目录")
}

// LoadConfig 加载配置
func LoadConfig(c interface{}) {
	v := viper.New()
	v.SetConfigName(configFile)
	v.AddConfigPath(configPath)
	v.SetConfigType("yaml")
	readErr := v.ReadInConfig()
	if readErr != nil {
		panic("配置文件读取失败，原因：" + readErr.Error())
	}
	unmarshalErr := v.Unmarshal(&c)
	if unmarshalErr != nil {
		panic("配置文件初始化结构体失败，原因：" + unmarshalErr.Error())
	}
}