package main

import (
	_ "embed"

	"GoFiberDemo.net/server/global"
	"GoFiberDemo.net/server/global/config"
	"GoFiberDemo.net/server/ready"
	"GoFiberDemo.net/server/router"
	jsoniter "github.com/json-iterator/go"
)

//go:embed package.json
var AppPackage []byte

func main() {
	jsoniter.Unmarshal(AppPackage, &config.AppInfo)
	// 初始化系统参数
	global.Start()

	// 启动数据准备
	ready.Start()

	// 启动 http 监听服务
	router.Start()
}
