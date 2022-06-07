package global

import (
	"fmt"

	"GoFiberDemo.net/server/global/config"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mPath"
)

/* ==================================================================== */
/* ================= ServerEnv ================= */
/* ==================================================================== */

func ServerEnvInt() {
	isAppEnvFile := mPath.Exists(config.File.AppServerEnv)
	isHomeEnvFile := mPath.Exists(config.File.ServerEnv)

	if isHomeEnvFile || isAppEnvFile {
		//
	} else {
		errStr := fmt.Errorf("没找到 server_env.yaml 配置文件")
		LogErr(errStr)
		panic(errStr)
	}

	if isAppEnvFile {
		config.LoadServerEnv(config.File.AppServerEnv)
	} else {
		config.LoadServerEnv(config.File.ServerEnv)
	}

	Log.Println("加载 ServerEnv : ", mJson.JsonFormat(mJson.ToJson(config.ServerEnv)))
}
