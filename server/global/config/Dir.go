package config

import (
	"os"

	"github.com/EasyGolang/goTools/mPath"
	"github.com/EasyGolang/goTools/mStr"
)

var Dir struct {
	Home string // Home 根目录
	App  string // APP 根目录
	Log  string // 日志文件目录
}

var File struct {
	ServerEnv    string // /root/server_env.yaml
	AppServerEnv string // ./server_env.yaml
	UserConfig   string // ./user_config.yaml
}

func PathInit() {
	Dir.Home = mPath.HomePath()

	Dir.App, _ = os.Getwd()

	Dir.Log = mStr.Join(
		Dir.App,
		mStr.ToStr(os.PathSeparator),
		"logs",
	)

	File.ServerEnv = mStr.Join(
		Dir.Home,
		mStr.ToStr(os.PathSeparator),
		"server_env.yaml",
	)
	File.AppServerEnv = mStr.Join(
		Dir.App,
		mStr.ToStr(os.PathSeparator),
		"server_env.yaml",
	)
	File.UserConfig = mStr.Join(
		Dir.App,
		mStr.ToStr(os.PathSeparator),
		"user_config.yaml",
	)
}
