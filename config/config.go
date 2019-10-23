package config

import (
	"github.com/spf13/viper"
	"log"
)
//配置文件
var  Config  *viper.Viper

func init() {
	Config = viper.New()
	//设置配置文件的名字
	Config.SetConfigName("config")
	//添加配置文件所在的路径,注意在Linux环境下%GOPATH要替换为$GOPATH
	Config.AddConfigPath("./")
	//设置配置文件类型
	Config.SetConfigType("yaml")

	if err := Config.ReadInConfig(); err != nil{
		log.Fatal("config.ReadInConfig error",err)
	}
}