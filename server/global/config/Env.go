package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var SysEnv struct {
	LocalIP       string `json:"LocalIP"`
	MongoAddress  string `json:"MongoAddress"`
	MongoUserName string `json:"MongoUserName"`
	MongoPassword string `json:"MongoPassword"`
	RunMod        int    `json:"RunMod"` // 0 为生产环境 1 为其它运行模式
}
var AppEnv struct {
	Port int `json:"Port"`
}
var AppInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func LoadSysEnv(envPath string) {
	viper.SetConfigFile(envPath)
	err := viper.ReadInConfig()
	if err != nil {
		errStr := fmt.Errorf("SysEnv 读取配置文件出错: %+v", err)
		LogErr(errStr)
		panic(errStr)
	}
	viper.Unmarshal(&SysEnv)
}

func LoadAppEnv() {
	viper.SetConfigFile(File.AppEnv)

	err := viper.ReadInConfig()
	if err != nil {
		errStr := fmt.Errorf("AppEnv 读取配置文件出错: %+v", err)
		LogErr(errStr)
		panic(errStr)
	}
	viper.Unmarshal(&AppEnv)
}
