package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var ServerEnv struct {
	LocalIP       string
	MongoAddress  string
	MongoUserName string
	MongoPassword string
	RunMod        int // 0 为生产环境 1 为其它运行模式
}

func LoadServerEnv(envPath string) {
	viper.SetConfigFile(envPath)
	err := viper.ReadInConfig()
	if err != nil {
		errStr := fmt.Errorf("ServerEnv 读取配置文件出错: %+v", err)
		LogErr(errStr)
		panic(errStr)
	}
	viper.Unmarshal(&ServerEnv)
}

var UserEnv struct {
	Port string
}

var AppInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
