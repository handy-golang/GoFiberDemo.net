package global

import (
	"fmt"

	"GoFiberDemo.net/server/global/config"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mPath"
)

func SysEnvInt() {
	isAppEnvFile := mPath.Exists(config.File.LocalSysEnv)
	isHomeEnvFile := mPath.Exists(config.File.SysEnv)

	if isHomeEnvFile || isAppEnvFile {
		//
	} else {
		errStr := fmt.Errorf("没找到 sys_env.yaml 配置文件")
		LogErr(errStr)
		panic(errStr)
	}

	if isAppEnvFile {
		config.LoadSysEnv(config.File.LocalSysEnv)
	} else {
		config.LoadSysEnv(config.File.SysEnv)
	}

	Log.Println("加载 SysEnv : ", mJson.JsonFormat(mJson.ToJson(config.SysEnv)))
}

func AppEnvInt() {
	// 检查配置文件在不在
	isUserEnvPath := mPath.Exists(config.File.AppEnv)
	if isUserEnvPath {
		config.LoadAppEnv()
	}

	if config.AppEnv.Port < 80 {
		config.AppEnv.Port = 9876
	}

	Log.Println("加载 AppEnv : ", mJson.JsonFormat(mJson.ToJson(config.AppEnv)))
}
